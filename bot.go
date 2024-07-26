package telebot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Bot struct {
	Token  string
	client *http.Client
	Me     *User
}

func New(token string) (*Bot, error) {
	b := &Bot{
		Token:  token,
		client: &http.Client{},
	}
	me, err := b.GetMe()
	if err != nil {
		return nil, err
	}
	b.Me = me
	return b, nil
}

func (b *Bot) GetMe() (*User, error) {
	apiResp, err := b.makeRequest("getMe")
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(apiResp.Result, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (b *Bot) GetUpdatesChan() <-chan *Update {
	ch := make(chan *Update, 100)

	go func() {
		for {
			updates, err := b.GetUpdates()
			if err != nil {
				close(ch)
				return
			}

			for _, upd := range updates {
				ch <- upd
			}

		}
	}()

	return ch
}

func (b *Bot) GetUpdates() ([]*Update, error) {
	apiResp, err := b.makeRequest("getUpdates")
	if err != nil {
		return nil, err
	}

	var updates []*Update
	if err := json.Unmarshal(apiResp.Result, &updates); err != nil {
		return nil, err
	}

	return updates, err
}

func (b *Bot) makeRequest(method string) (*APIResponse, error) {
	endpoint = fmt.Sprintf(endpoint, b.Token, method)
	req, err := http.NewRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := b.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp APIResponse
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &apiResp)
	if err != nil {
		return nil, err
	}

	if !apiResp.OK {
		return nil, fmt.Errorf("error: %d %s", apiResp.ErrorCode, apiResp.Description)
	}

	return &apiResp, err
}
