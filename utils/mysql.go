package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/switfs/oms/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func Init() {
	GVA_DB = GormMysql()
}

func GormMysql() *gorm.DB {
	m := config.SR.Mysql
	if m.DbName == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       Dsn(m), // DSN data source name
		DefaultStringSize:         191,    // string 类型字段的默认长度
		SkipInitializeWithVersion: false,  // 根据版本自动配置
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxOpenConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}



func Dsn(m *config.Mysql) string {
	return m.Users + ":" + m.Passwd + "@tcp(" + m.Hosts + ":" + m.Ports + ")/" + m.DbName + "?parseTime=true&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci&readTimeout=10s&writeTimeout=10s"
}
