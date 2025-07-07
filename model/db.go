package model

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var db *gorm.DB

func InitDb() {

	// 设置日志级别
	logMode := logger.Info
	if !viper.GetBool("mode.develop") {
		logMode = logger.Error
	}
	//连接数据库
	var err error
	db, err = gorm.Open(mysql.Open(viper.GetString("db.dsn")), &gorm.Config{Logger: logger.Default.LogMode(logMode)})
	if err != nil {
		fmt.Println("数据库连接失败，请检查参数：", err)
		return
	}
	fmt.Println("数据库连接成功")
	//获取底层连接池
	sqlDB, _ := db.DB()
	err = db.AutoMigrate(&User{}, &Article{})
	if err != nil {
		fmt.Println("自动迁移失败:", err)
		return
	}
	fmt.Println("自动迁移成功")
	//配置连接池
	sqlDB.SetMaxIdleConns(viper.GetInt("db.SetMaxIdleConns"))
	sqlDB.SetMaxOpenConns(viper.GetInt("db.SetMaxOpenConns"))
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
