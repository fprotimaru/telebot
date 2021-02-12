package telebot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Bot struct {
	Token string

	Client *http.Client
}

func New(token string) *Bot {
	return &Bot{
		Token:  token,
		Client: &http.Client{},
	}
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

func (b *Bot) makeRequest(method string) (*APIResponse, error) {
	endpoint = fmt.Sprintf(endpoint, b.Token, method)
	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return nil, err
	}
	// TODO should I close req.Body?
	defer req.Body.Close()

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := b.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp APIResponse
	data, err := ioutil.ReadAll(resp.Body)
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
