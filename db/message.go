package db

type MESSAGE_STATUS_ENUM int

const (
	MESSAGE_STATUS_NOT_SEND MESSAGE_STATUS_ENUM = iota
	MESSAGE_STATUS_SENDED
)

// TextUrl 每日新闻内容
type MsgTextUrl struct {
	Model
	MsgText string `json:"msg_text"`
	MsgUrl  string `json:"msg_url"`
}

type Message struct {
	Model
	TextUrls   []MsgTextUrl        `json:"text_urls"`
	DailyTitle string              `json:"daily_title"`
	Author     string              `json:"author"`   // 编辑
	PostUrl    string              `json:"post_url"` // 原文地址
	Status     MESSAGE_STATUS_ENUM `json:"status"`
}
