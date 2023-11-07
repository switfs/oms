package utils

import "os"

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables() {
	db := GVA_DB
	err := db.AutoMigrate(
	// 系统模块表

	)
	if err != nil {
		log.Error("register table failed", err.Error())
		os.Exit(0)
	}
	log.Info("register table success")
}
