package chatbot

const (
	PotatoChartTypePerson  = 1
	PotatoChartTypeGroup   = 2
	PotatoChartTypeChannel = 3
)

type PotatoClient interface {
	SetWebhook(opts PotatoWebhookOpts) (err error)
	RespUpdateMessage(opts SendTextMessage) (err error)
	BroadcastToGroup(msg string) (err error)
}

// PotatoWebhookOpts 設定 Webhook 資訊
type PotatoWebhookOpts struct {
	URL string `json:"url"`
	// Certificate string `json:"certificate"` // or type file
}

// UpdateMessage 接收 potato chatbot 訊息的格式 ( ps: 用 Update 是因為 potato api 文件用該詞於定義接收訊息)
type UpdateMessage struct {
	Ok     bool `json:"ok"`
	Result []struct {
		UpdateID int `json:"update_id"`
		Message  struct {
			MessageID int `json:"message_id"`
			Chat      struct {
				ID    int    `json:"id"`
				Type  int    `json:"type"`
				Title string `json:"title"`
			} `json:"chat"`
			From struct {
				ID        int    `json:"id"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				Username  string `json:"username"`
			} `json:"from"`
			Text string `json:"text"`
			Date int    `json:"date"`
		} `json:"message"`
		InlineQuery interface{} `json:"inline_query"`
		Lang        string      `json:"lang"`
	} `json:"result"`
}

// SendTextMessage 傳送訊息
type SendTextMessage struct {
	ChatType         int    `json:"chat_type"`
	ChatID           int64  `json:"chat_id"`
	Text             string `json:"text"`
	ReplyToMessageID int    `json:"reply_to_message_id"`
	Markdown         bool   `json:"markdown"`
	// ReplyMarkup      *ReplyMarkup `json:"reply_markup"`
}

// ReplyMarkup 按鈕選項
type ReplyMarkup struct {
	Type           int `json:"type"`
	InlineKeyboard []struct {
		Buttons []struct {
			Text         string `json:"text"`
			CallbackData string `json:"callback_data"`
		} `json:"buttons"`
	} `json:"inline_keyboard"`
}

// Groups 取得群組的訊息
type Groups struct {
	Items struct {
		Channels []struct {
			ID   int64  `json:"PeerID"`
			Name string `json:"PeerName"`
		} `json:"Channels"`
		Groups []struct {
			ID   int64  `json:"PeerID"`
			Name string `json:"PeerName"`
		} `json:"Groups"`
	} `json:"result"`
}
