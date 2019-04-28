package config

import (
	"encoding/json"
	"github.com/toolkits/file"
	"news/base"
	"sync"
)

var (
	config  *GlobalConfig
	cfglock sync.RWMutex
)

func Cfg() *GlobalConfig {
	cfglock.Lock()
	defer cfglock.Unlock()
	return config
}

func ParserConfig(cfg string) {
	if cfg == "" {
		panic("use -c to specify configuration file")
		return
	}
	if !file.IsExist(cfg) {
		panic("config file is not existent")
		return
	}
	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		panic("read config file error:" + err.Error())
		return
	}
	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		panic("parser config file error " + err.Error())
		return
	}
	cfglock.Lock()
	defer cfglock.Unlock()
	config = &c
	base.Log("read config file:%v successfully", cfg)
	return
}

type GlobalConfig struct {
	Splider struct {
		Parallelism int  `json:"parallelism"`
		Delay       int  `json:"delay"`
		All         bool `json:"all"`
	} `json:"splider"`
	Wechat struct {
		Appid   string `json:"appid"`
		Seceret string `json:"seceret"`
		Token   string `json:"token"`
	} `json:"wechat"`
	Mysql struct {
		Addr         string `json:"addr"`
		MaxIdleConns int    `json:"max_idle_conns"`
		MaxOpenConns int    `json:"max_open_conns"`
	} `json:"mysql"`
	Http struct {
		Listen string `json:"listen"`
	} `json:"http"`
}
