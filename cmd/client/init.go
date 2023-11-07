package main

import (
	"github.com/BurntSushi/toml"
	logging "github.com/ipfs/go-log/v2"
	"github.com/switfs/oms/config"
)

var log = logging.Logger("main")

func init() {
	if _, err := toml.DecodeFile("./config.toml", &config.CR); err != nil {
		log.Errorf("配置文件初始失败 %s", err.Error())
		return
	}

}
