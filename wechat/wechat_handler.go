package wechat

import (
	"fmt"
	"news/base"
	"news/db"
	"news/g"
	"news/model"
	"time"
)

func Subscribe(baseMsg *model.WeChatBaseMessage) error {
	var userInfo db.WeChatUserInfo
	userInfo.OpenId = baseMsg.FromUserName
	err := db.DB.Model(&db.WeChatUserInfo{}).Save(&userInfo).Error
	if err != nil {
		base.Log("save user info error:%v", err)
		return err
	}
	return nil
}

func UnSubscribe(baseMsg *model.WeChatBaseMessage) {
	return
}

func SendTextMessage(baseMsg *model.WeChatBaseMessage, content string) (string, error) {
	var textMessage model.WeChatTextMessage
	textMessage.ToUserName = baseMsg.FromUserName
	textMessage.FromUserName = baseMsg.ToUserName
	textMessage.MsgType = g.RESP_MESSAGE_TYPE__TEXT
	textMessage.CreateTime = time.Now().Unix()
	textMessage.Content = content

	return TextMessageToXml(textMessage)
}

func GetLatestNews(baseMsg *model.WeChatBaseMessage) (respMsg string, newsType string, errs error) {
	var newsMsg model.WeChatNewsMessage
	var msg db.Message
	err := db.DB.Model(&db.Message{}).Where("created_at > ?", time.Date(time.Now().Year(),
		time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)).First(&msg).Error
	if err != nil {
		base.Log("db query lastest news error:%v", err)
		return
	}
	for idx, _ := range msg.TextUrls {

		urlDetail := buildItemUrlDetail()

		var article model.WeChatArticle
		article.Title = GetWechatWarnNewsTitle(dpDataMap[key].VideoTime, dpDataMap[key].Cid, dpDataMap[key].AliasName)
		article.Description = GetWechatWarnNewsDescription()
		article.PicUrl = timeUrlList[0]
		article.Url = urlDetail
		newsMsg.Articles.Item = append(newsMsg.Articles.Item, article)
	}

	newsMsg.ToUserName = baseMsg.FromUserName
	newsMsg.FromUserName = baseMsg.ToUserName
	newsMsg.CreateTime = time.Now().Unix()
	newsMsg.MsgType = g.RESP_MESSAGE_TYPE__NEWS

	newsMsg.ArticleCount = len(newsMsg.Articles.Item)

	respMsg, _ = NewsMessageToXml(newsMsg)
	newsType = g.RESP_MESSAGE_TYPE__NEWS
	errs = nil

	return
}

func SendLatestNew(baseMsg *model.WeChatBaseMessage, cnt string, newsType string) (string, error) {

	return "", nil
}
