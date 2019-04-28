package cron

import (
	"news/splider"
	"news/wechat"
	"time"
)

func AutoFetch() {
	var timeTomorrow time.Time
	tomorrow := time.Now().Add(24 * time.Hour)
	interToday := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(),
		8, 30, 00, 0, time.Local).Unix() - time.Now().Unix()
	time.Sleep(time.Second * time.Duration(interToday))
	go func() {
		for {
			go splider.Run()
			timeTomorrow = time.Now().Add(24 * time.Hour)
			time.Sleep(time.Second * time.Duration(time.Date(tomorrow.Year(), tomorrow.Month(),
				tomorrow.Day(), 8, 30, 0, 0, time.Local).Unix()))
		}
	}()
}

func AutoPublish() {
	var timeTomorrow time.Time
	tomorrow := time.Now().Add(24 * time.Hour)
	interToday := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(),
		8, 30, 00, 0, time.Local).Unix() - time.Now().Unix()
	time.Sleep(time.Second * time.Duration(interToday))
	go func() {
		for {
			go wechat.Publish()
			timeTomorrow = time.Now().Add(24 * time.Hour)
			time.Sleep(time.Second * time.Duration(time.Date(tomorrow.Year(), tomorrow.Month(),
				tomorrow.Day(), 8, 30, 0, 0, time.Local).Unix()))
		}
	}()
}
