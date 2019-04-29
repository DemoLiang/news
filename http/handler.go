package http

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"news/base"
	"news/config"
	"news/g"
	"news/model"
	"news/wechat"
	"sort"
	"strings"
	"time"
)

type Ret struct {
	Code   base.ErrorCode `json:"code"`
	ErrMsg string         `json:"err_msg"`
	Status string         `json:"status"`
}

type RespData struct {
	Ret
	Data interface{} `json:"data"`
}

func WriteResp(ctx *gin.Context, code base.ErrorCode, data interface{}) {
	resp := RespData{}
	resp.Code = code
	resp.ErrMsg = code.Error()
	resp.Data = data

	ctx.JSON(http.StatusOK, &resp)
}

func WechatWriteResp(ctx *gin.Context, data interface{}) {
	ctx.XML(http.StatusOK, data.([]byte))
}

func versionHandler(ctx *gin.Context) {
	WriteResp(ctx, base.Success, gin.Version)
	return
}

func reloadHandler(ctx *gin.Context) {
	//TODO reload cfg
	WriteResp(ctx, base.Success, nil)
	return
}

func healthHandler(ctx *gin.Context) {
	WriteResp(ctx, base.Success, "health")
	return
}

func wechatAuthHandler(ctx *gin.Context) {
	signature := ctx.GetString("signature")
	timestamp := ctx.GetString("timestamp")
	nonce := ctx.GetString("nonce")
	echostr := ctx.GetString("echostr")

	var list = []string{config.Cfg().Wechat.Token, timestamp, nonce}
	sort.Strings(list)
	sigStr := strings.Join(list, "")

	sha1 := sha1.New()
	sha1.Write([]byte(sigStr))
	hash := sha1.Sum(nil)
	base.Log("signature:%v \n timestamp:%v \n nonce:%v\n echostr:%v \n"+
		"sigStr:%v \n list:%v\n hash:%v\n", signature, timestamp, nonce, echostr, sigStr, list, fmt.Sprintf("%x", hash))
	if signature == fmt.Sprintf("%x", hash) {
		base.Log("hash eq echostr ec")
		ctx.Writer.Write([]byte(echostr))
	}
	return
}

func wechatHandler(ctx *gin.Context) {
	var respMessage string
	var baseMsg model.WeChatBaseMessage
	var err error

	buf := make([]byte, 1024)
	n, err := ctx.Request.Body.Read(buf)
	if err != nil {
		base.Log("read body error:%v", err)
		return
	}
	base.Log("read body %v bytes", n)

	err = xml.Unmarshal(buf, &baseMsg)
	if err != nil {
		base.Log("unmarshal request error:%v", err)
		return
	}
	base.Log("%v", baseMsg)

	switch baseMsg.MsgType {
	case g.REQ_MESSAGE_TYPE__TEXT:
		var textMsg model.WeChatTextMessage
		err = xml.Unmarshal(buf, &textMsg)
		if err != nil {
			base.Log("unmarshal request error:%v", err)
		}
		base.Log("%v->%v", textMsg.FromUserName, textMsg.Content)
		respMessage = "success"
	case g.REQ_MESSAGE_TYPE__IMAGE:
		base.Log("%v : sending a image message", baseMsg.FromUserName)
		respMessage = "success"
	case g.REQ_MESSAGE_TYPE__LOCATION:
		base.Log("%v : sending a location message", baseMsg.FromUserName)
		respMessage = "success"
	case g.REQ_MESSAGE_TYPE__LINK:
		base.Log("%v : sending a link message", baseMsg.FromUserName)
		respMessage = "success"
	case g.REQ_MESSAGE_TYPE__VOICE:
		base.Log("%v : sending a voice message", baseMsg.FromUserName)
		respMessage = "success"
	case g.REQ_MESSAGE_TYPE__EVENT:
		switch baseMsg.Event {
		case g.REQ__EVENT_TYPE__SUBSCRIBE:
			base.Log("%v : sending a Subscribe Event", baseMsg.FromUserName)
			respMessage = "success"
			respMessage, _ = wechat.SendTextMessage(&baseMsg, g.DEFAULT_SUBSCRIBE__MSG)
			wechat.Subscribe(&baseMsg)
		case g.REQ_MESSAGE_TYPE_EVENT__UNSUBSCRIBE:
			base.Log("%v : sending a unsubscribe Event", baseMsg.FromUserName)
			wechat.UnSubscribe(&baseMsg)
			respMessage = "success"
		case g.REQ__EVENT_TYPE__CLICK:
			var menuClickMsg model.WeChatMenuClick
			err = xml.Unmarshal(buf, &menuClickMsg)
			if err != nil {
				base.Log("unmarshal request error:%v", err)
			}
			base.Log("%v : sending a menuClickMsg Event-->%v", baseMsg.FromUserName, menuClickMsg.EventKey)
			switch menuClickMsg.EventKey {
			case g.WECHAT__MENU_BUTTON__BT0:
				menuClickMsg.ToUserName = baseMsg.FromUserName
				menuClickMsg.FromUserName = baseMsg.ToUserName
				menuClickMsg.CreateTime = time.Now().Unix()
				menuClickMsg.MsgType = baseMsg.MsgType
				respMessage, _ = wechat.SendTextMessage(&baseMsg, g.DEFAULT_LATEST_NEWS_NULL)
				respMsgCnt, newsType, _ := wechat.GetLatestNews(&baseMsg)
				respMessage, err = wechat.SendLatestNew(&baseMsg, respMsgCnt, newsType)
			case g.WECHAT__MENU_BUTTON__BT1:
				respMessage = "success"
				base.Log("%v : sending a menuClickMsg Event-->%v", baseMsg.FromUserName, menuClickMsg.EventKey)
			case g.WECHAT__MENU_BUTTON__BT2:
				respMessage = "success"
				base.Log("%v : sending a menuClickMsg Event-->%v", baseMsg.FromUserName, menuClickMsg.EventKey)
			default:
				respMessage = "success"
				base.Log("%v : sending a menuClickMsg Event-->%v", baseMsg.FromUserName, menuClickMsg.EventKey)
			}
		case g.REQ__EVENT_TYPE__VIEW:
			var menuClickMsg model.WeChatMenuClick
			err = xml.Unmarshal(buf, &menuClickMsg)
			if err != nil {
				base.Log("unmarshal request error:%v", err)
			}
			base.Log("%v : sending a menuClickMsg Event View-->%v", baseMsg.FromUserName, menuClickMsg.EventKey)
			respMessage = "success"
		default:
			base.Log("%v", baseMsg.Event)
			respMessage = "success"
		}
	default:
		base.Log("msgType is :%v", baseMsg.MsgType)
		respMessage = "success"
	}

	WechatWriteResp(ctx, respMessage)
	return
}

func newsHandler(ctx *gin.Context) {
	return
}

func defaultHandler(ctx *gin.Context) {
	WriteResp(ctx, base.Success, "hiahiahia")
	return
}
