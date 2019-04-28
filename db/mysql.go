package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"news/base"
	"news/config"
)

var (
	DB *gorm.DB
)

func Init() {
	if db, err := gorm.Open("mysql", config.Cfg().Mysql.Addr); err == nil {
		db.Callback().Create().Before("gorm:create").Register("auto_gen_string_primary_key", CreateCallback)
		db.DB().SetMaxIdleConns(config.Cfg().Mysql.MaxIdleConns)
		db.DB().SetMaxOpenConns(config.Cfg().Mysql.MaxOpenConns)
		DB = db
		return
	} else {
		base.Log("初始化数据库失败:%v", err.Error())
	}
}

func TableRegister(tableStruRef interface{}) {
	if !DB.HasTable(tableStruRef) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(tableStruRef)
	}
}
