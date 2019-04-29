package model

import (
	"encoding/xml"
	"sync"
)

//获取accesstoken的输出结构
type WeChatAuthOutput struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

//微信的用户信息
type WeChatUserInfo struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	UserId  string `json:"userid"`
}

//加密的请求内容
type EncryptRequestBody struct {
	XMLName    xml.Name `xml:"xml"`
	ToUserName string   `xml:"ToUserName"`
	Encrypt    string   `xml:"Encrypt"`
	AgentId    string   `xml:"AgentID"`
}

//加密的返回内容
type EncryptResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      string   `xml:"Encrypt"`
	MsgSignature string   `xml:"MsgSignature"`
	TimeStamp    string   `xml:"TimeStamp"`
	Nonce        string   `xml:"Nonce"`
}

//XMl CDATA 数据编码使用
type CDATAText struct {
	Text string `xml:",innerxml"`
}

//基础请求信息
type BasicRequestInfo struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Event        string   `xml:"Event"`
	MsgId        int64    `xml:"MsgId"`
	AgentID      string   `xml:"AgentID"`
}

//基础响应信息，返回格式
type BasicResponseInfo struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
}

//文本请求body content 内容
type TextRequestMessage struct {
	BasicRequestInfo
	Content string `xml:"Content"`
	MsgId   string `xml:"MsgId"`
}

//文本请求响应body content 内容
type TextResponseMessage struct {
	BasicResponseInfo
	Content string `xml:"Content"`
}

//图片消息请求body content 内容
type ImageRequestMessage struct {
	BasicRequestInfo
	PicUrl   string `xml:"PicUrl"`
	MedialId string `xml:"MedialId"`
}

//图片消息回复消息格式
type ImageResponseMessage struct {
	BasicResponseInfo
	Image struct {
		MediaId string `xml:"MediaId"`
	} `xml:"Image"`
}

//图片消息响应媒体结构
//type ImageResponseImage struct {
//	MediaId string `xml:"MediaId"`
//}

//语音消息请求body content 内容
type VoiceRequestMessage struct {
	BasicRequestInfo
	MediaId string `xml:"MediaId"`
	Format  string `xml:"Format"`
}

//语音消息响应消息格式
type VoiceResponseMessage struct {
	BasicResponseInfo
	Voice struct {
		MediaId string `xml:"MediaId"`
	} `xml:"Voice"`
}

//视频消息请求body
type VideoRequestMessage struct {
	BasicRequestInfo
	MediaId      string `xml:"MediaId"`
	ThumbMediaId string `xml:"ThumbMediaId"`
}

//视频消息响应格式
type VideoResponseMessage struct {
	BasicResponseInfo
	Video struct {
		MediaId     string `xml:"MediaId"`
		Title       string `xml:"Title"`
		Description string `xml:"Description"`
	} `xml:"Video"`
}

//图文消息响应格式
type NewsResponseMessage struct {
	BasicResponseInfo
	ArticleCount int `xml:"ArticleCount"`
	Articles     struct {
		Item []NewsItem `xml:"Item"`
	} `xml:"Articles"`
}

//图文列表结构
type NewsItem struct {
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	PicUrl      string `xml:"PicUrl"`
	Url         string `xml:"Url"`
}

//位置消息请求body
type LocationRequestMessage struct {
	BasicRequestInfo
	LocationX float64 `xml:"Location_X"`
	LocationY float64 `xml:"Location_Y"`
	Scale     int     `xml:"Scale"`
	Label     string  `xml:"Label"`
}

//链接消息请求body
type LinkRequestMessage struct {
	BasicRequestInfo
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	PicUrl      string `xml:"PicUrl"`
}

//关注消息事件请求body
type SubscribeRequestEventMessage struct {
	BasicRequestInfo
}

//菜单点击请求事件
type MenuClickRequestEventMessage struct {
	BasicRequestInfo
	EventKey string `xml:"EventKey"`
}

//菜单点击跳转到链接事件
type MenuClickRedirectRequesstEventMessage struct {
	BasicRequestInfo
	EventKey string `xml:"EventKey"`
}

//发送图片的信息
type PicSysphotoSendPicsInfo struct {
	Count   int                  `xml:"Count"`
	PicList []PicSysphotoPicList `xml:"PicList"`
}

//图片列表
type PicSysphotoPicList struct {
	Item PicSysphotoPicItem `xml:"Item"`
}

//图片MD5信息
type PicSysphotoPicItem struct {
	PicMd5Sum string `xml:"PicMd5Sum"`
}

//弹出地理位置选择器的事件推送
type LocationSelectRequestEventMessage struct {
	BasicRequestInfo
	EventKey         string                         `xml:"EventKey"`
	SendLocationInfo LocationSelectSendLocationInfo `xml:"SendLocationInfo"`
}

//发送的位置信息
type LocationSelectSendLocationInfo struct {
	LocationX float64 `xml:"Location_X"`
	LocationY float64 `xml:"Location_Y"`
	Scale     int     `xml:"Scale"`
	Label     string  `xml:"Label"`
	Poiname   string  `xml:"Poiname"`
}

//wechat news
type CDATA string

type WeChatToken struct {
	TokenLock sync.RWMutex `json:"token_lock"`
	Token     string       `json:"token"`
}

type WeChatBaseMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Event        string   `xml:"Event"`
}

type WeChatArticle struct {
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	PicUrl      string `xml:"PicUrl"`
	Url         string `xml:"Url"`
}

type WeChatArticleItem struct {
	Item []WeChatArticle `xml:"item"`
}

type WeChatTextMessage struct {
	WeChatBaseMessage
	Content string `xml:"Content"`
}

type WeChatNewsMessage struct {
	WeChatBaseMessage
	ArticleCount int               `xml:"ArticleCount"`
	Articles     WeChatArticleItem `xml:"Articles"`
}

type WeChatEventMessage struct {
	WeChatBaseMessage
}

type WeChatMenuClick struct {
	WeChatBaseMessage
	Event    string `xml:"Event"`
	EventKey string `xml:"EventKey"`
}
