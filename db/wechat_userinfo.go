package db

type WeChatUserInfo struct {
	Model
	OpenId string `json:"open_id"`
}
