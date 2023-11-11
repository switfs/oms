package main

import (
	"github.com/BurntSushi/toml"
	"github.com/switfs/oms/config"
	"github.com/switfs/oms/utils"
	"github.com/urfave/cli/v2"
)

var Run = &cli.Command{
	Name:  "run",
	Usage: "Proxy  运行服务",
	Action: func(cctx *cli.Context) error {
		log.Info("ok")

		if _, err := toml.DecodeFile("./config.toml", &config.PR); err != nil {
			log.Errorf("配置文件初始失败 %s", err.Error())
			panic("配置文件初始失败")
		}
		utils.Init()
		if utils.GVA_DB != nil {
			utils.RegisterTables()
			db, _ := utils.GVA_DB.DB()
			defer db.Close()
		}
		return nil
	},
}
