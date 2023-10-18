package main

import (
	"github.com/BurntSushi/toml"
	"github.com/switfs/oms/config"
)

func init() {
	if _, err := toml.DecodeFile("./config.toml", &config.CR); err != nil {
		log.Errorf("配置文件初始失败 %s", err.Error())
		return
	}

}
