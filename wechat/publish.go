package wechat

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"news/base"
	"news/db"
	"news/message"
	"strings"
	"time"
)

func buildMessage(msg db.Message) string {
	str := fmt.Sprintf("## %s", msg.DailyTitle)
	str += fmt.Sprintln()
	for _, v := range msg.TextUrls {
		if len(v.MsgUrl) == 0 || len(v.MsgText) == 0 {
			continue
		}

		if strings.Contains(v.MsgText, "GoCN归档") || strings.Contains(v.MsgText, "订阅新闻") {
			continue
		}

		textValue := strings.Replace(v.MsgText, "\n", "", -1)
		realText := strings.Replace(textValue, " ", "", -1)

		str += fmt.Sprintf("- [%s](%s)", realText, v.MsgUrl)
		str += fmt.Sprintln()
	}

	index := strings.Index(msg.Author, "订阅新闻")
	author := msg.Author
	if index > 0 {
		author = msg.Author[:index]
	}

	str += fmt.Sprintln()
	str += fmt.Sprintf("编辑：%s", author)
	str += fmt.Sprintln()
	str += fmt.Sprintln()
	str += fmt.Sprintf("原文地址: %s", msg.PostUrl)
	return str
}

func SendNews(msg db.Message) error {
	var userList []db.WeChatUserInfo
	err := db.DB.Model(&db.WeChatUserInfo{}).Find(&userList).Error
	if err != nil {
		base.Log("query wechat user info error:%v", err)
		return err
	}
	//var news model.WeChatNewsMessage
	return nil
}

func CheckSend(dailyTitle string) bool {
	var msg db.Message
	err := db.DB.Model(&db.Message{}).Where("daily_title = ?", dailyTitle).First(&msg).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false
	}
	if msg.ID != "" || err == gorm.ErrRecordNotFound {
		return true
	}
	return false
}

func Insert(msg db.Message) error {
	err := db.DB.Model(&db.Message{}).Save(&msg).Error
	if err != nil {
		base.Log("save to db error:%v", err)
		return err
	}
	return nil
}

func Publish() {
	msg, err := message.Pop()
	if err != nil {
		return
	}

	if !CheckSend(msg.DailyTitle) {
		//content := buildMessage(msg)
		//_ = content
		Insert(msg)
		SendNews(msg)
		time.Sleep(time.Second * 1)
	}

}
