package main

import (
	"github.com/BurntSushi/toml"
	logging "github.com/ipfs/go-log/v2"
	"github.com/switfs/oms/config"
	"github.com/switfs/oms/utils"
)

var log = logging.Logger("main")

func init() {
	if _, err := toml.DecodeFile("./config.toml", &config.SR); err != nil {
		log.Errorf("配置文件初始失败 %s", err.Error())
		return
	}
	utils.GVA_DB = utils.Gorm()
	if utils.GVA_DB != nil {
		utils.RegisterTables()
		db, _ := utils.GVA_DB.DB()
		defer db.Close()
	}
}
