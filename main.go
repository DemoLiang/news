package main

import (
	"flag"
	"news/config"
	"news/http"
	"news/wechat"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := flag.String("c", "./cfg.json", "cfg file")
	flag.Parse()
	config.ParserConfig(*cfg)

	wechat.InitWeChatToken()
	go wechat.UpdateWeChatToken()

	//go cron.AutoFetch()
	//go cron.AutoPublish()

	go http.Start()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		os.Exit(0)
	}()
	select {}
}
