package main

import (
	"github.com/BurntSushi/toml"
	"github.com/switfs/oms/config"
	"github.com/switfs/oms/server"
	"github.com/switfs/oms/utils"
	"github.com/urfave/cli/v2"
)

var Run = &cli.Command{
	Name:  "run",
	Usage: "运行服务",
	Action: func(cctx *cli.Context) error {

		if _, err := toml.DecodeFile("./config.toml", &config.SR); err != nil {
			panic("配置文件初始失败")
		}

		utils.Init()
		if utils.GVA_DB != nil {
			utils.RegisterTables()
			db, _ := utils.GVA_DB.DB()
			defer db.Close()
		}
		log.Info("ok")

		if err := server.RunNewServer(); err != nil {
			log.Error("server run error on", err.Error())
			panic(err)
		}
		log.Info("server run success ")
		return nil
	},
}
