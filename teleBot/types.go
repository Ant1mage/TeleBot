package teleBot

import (
	"encoding/json"
)

type APIResponse struct {
	Ok          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
	Parameters  *ResponseParameters `json:"parameters"`
}

type Error struct {
	Message string
	ResponseParameters
}

func (e Error) Error() string {
	return e.Message
}

type ResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id"` // optional
	RetryAfter      int   `json:"retry_after"`        // optional
}

type User struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`     // optional
	UserName     string `json:"username"`      // optional
	LanguageCode string `json:"language_code"` // optional
	IsBot        bool   `json:"is_bot"`        // optional
}

type Message struct {
	MessageID            int              `json:"message_id"`
	From                 *User            `json:"from"` // optional
	Date                 int              `json:"date"`
	Chat                 *Chat            `json:"chat"`
	ForwardFrom          *User            `json:"forward_from"`            // optional
	ForwardFromChat      *Chat            `json:"forward_from_chat"`       // optional
	ForwardFromMessageID int              `json:"forward_from_message_id"` // optional
	ForwardDate          int              `json:"forward_date"`            // optional
	ReplyToMessage       *Message         `json:"reply_to_message"`        // optional
	EditDate             int              `json:"edit_date"`               // optional
	Text                 string           `json:"text"`                    // optional
	Entities             *[]MessageEntity `json:"entities"`                // optional
	// Voice                 *Voice             `json:"voice"`                   // optional
	Caption string `json:"caption"` // optional
	// Contact               *Contact           `json:"contact"`                 // optional
	// Location              *Location          `json:"location"`                // optional
	// Venue                 *Venue             `json:"venue"`                   // optional
	NewChatMembers        *[]User  `json:"new_chat_members"`        // optional
	LeftChatMember        *User    `json:"left_chat_member"`        // optional
	NewChatTitle          string   `json:"new_chat_title"`          // optional
	DeleteChatPhoto       bool     `json:"delete_chat_photo"`       // optional
	GroupChatCreated      bool     `json:"group_chat_created"`      // optional
	SuperGroupChatCreated bool     `json:"supergroup_chat_created"` // optional
	ChannelChatCreated    bool     `json:"channel_chat_created"`    // optional
	MigrateToChatID       int64    `json:"migrate_to_chat_id"`      // optional
	MigrateFromChatID     int64    `json:"migrate_from_chat_id"`    // optional
	PinnedMessage         *Message `json:"pinned_message"`          // optional
	// Invoice               *Invoice           `json:"invoice"`                 // optional
	// SuccessfulPayment     *SuccessfulPayment `json:"successful_payment"`      // optional
}

type Chat struct {
	ID                  int64  `json:"id"`
	Type                string `json:"type"`
	Title               string `json:"title"`                          // optional
	UserName            string `json:"username"`                       // optional
	FirstName           string `json:"first_name"`                     // optional
	LastName            string `json:"last_name"`                      // optional
	AllMembersAreAdmins bool   `json:"all_members_are_administrators"` // optional
	// Photo               *ChatPhoto `json:"photo"`
	Description string `json:"description,omitempty"` // optional
	InviteLink  string `json:"invite_link,omitempty"` // optional
}

type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url"`  // optional
	User   *User  `json:"user"` // optional
}

// 收取信息的channel
type UpdatesChannel <-chan Update

type UpdateConfig struct {
	Offset  int
	Limit   int
	Timeout int
}

type Update struct {
	UpdateID          int      `json:"update_id"`
	Message           *Message `json:"message"`
	EditedMessage     *Message `json:"edited_message"`
	ChannelPost       *Message `json:"channel_post"`
	EditedChannelPost *Message `json:"edited_channel_post"`
	// InlineQuery        *InlineQuery        `json:"inline_query"`
	// ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
	// CallbackQuery      *CallbackQuery      `json:"callback_query"`
	// ShippingQuery      *ShippingQuery      `json:"shipping_query"`
	// PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query"`
}

// BaseChat is base type for all chat config types.
type BaseChat struct {
	ChatID              int64 // required
	ChannelUsername     string
	ReplyToMessageID    int
	ReplyMarkup         interface{}
	DisableNotification bool
}

// MessageConfig contains information about a SendMessage request.
type MessageConfig struct {
	BaseChat
	Text                  string
	Photo                 string
	ParseMode             string
	DisableWebPagePreview bool
	MessageType           int
}

type TuringParams struct {
	ReqType    int              `json:"reqType"`
	Perception TuringPerception `json:"perception"`
	UserInfo   TuringUser       `json:"userInfo"`
}

type TuringPerception struct {
	InputText  TuringInputText  `json:"inputText"`
	InputImage TuringInputImage `json:"inputImage"`
}

type TuringInputText struct {
	Text string `json:"text"`
}

type TuringInputImage struct {
	Url string `json:"url"`
}

type TuringUser struct {
	ApiKey string `json:"apiKey"`
	UserId string `json:"userId"`
}

type TuringResponse struct {
	Intent  TuringIntent    `json:"intent"`
	Results []TuringResults `json:"results"`
}

type TuringIntent struct {
	Code       int    `json:"code"`
	IntentName string `json:"intentName"`
	ActionName string `json:"actionName"`
}

type TuringResults struct {
	GroupType  int          `json:"groupType"`
	ResultType string       `json:"resultType"`
	Values     TuringValues `json:"values"`
}

type TuringValues struct {
	Text  string `json:"text"`
	Url   string `json:"url"`
	Voice string `json:"voice"`
	Image string `json:"image"`
}
