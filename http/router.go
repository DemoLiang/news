package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"news/base"
	"news/config"
)

var GRouter *gin.Engine

func Start() {

	addr := config.Cfg().Http.Listen
	if addr == "" {
		panic(errors.New("http.addr is empty"))
		return
	}
	err := GRouter.Run(addr)
	if err != nil {
		base.Log("")
		panic("http run failed:" + err.Error())
	}
	return
}

func init() {
	r := gin.Default()
	r.GET("/version", versionHandler)
	r.GET("/reload", reloadHandler)
	r.GET("/health", healthHandler)
	httpGroup := r.Group("/api/v1/wechat")
	{
		httpGroup.GET("/auth", wechatAuthHandler)
		httpGroup.POST("/auth", wechatHandler)
		httpGroup.POST("/news", newsHandler)
	}
	r.GET("/server", defaultHandler)
	GRouter = r
	return
}
