package telebot

import "encoding/json"

type APIResponse struct {
	OK          bool            `json:"ok"`
	Result      json.RawMessage `json:"result"`
	ErrorCode   int             `json:"error_code"`
	Description string          `json:"description"`
}

type Update struct {
	UpdateID           int32       `json:"update_id"`
	Message            interface{} `json:"message"`
	EditedMessage      interface{} `json:"edited_message"`
	ChannelPost        interface{} `json:"channel_post"`
	EditedChannelPost  interface{} `json:"edited_channel_post"`
	InlineQuery        interface{} `json:"inline_query"`
	ChosenInlineResult interface{} `json:"chosen_inline_result"`
	CallbackQuery      interface{} `json:"callback_query"`
	ShippingQuery      interface{} `json:"shipping_query"`
	PreCheckoutQuery   interface{} `json:"pre_checkout_query"`
	Poll               interface{} `json:"poll"`
	PollAnswer         interface{} `json:"poll_answer"`
}

type User struct {
	Id                      int32  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Username                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}
