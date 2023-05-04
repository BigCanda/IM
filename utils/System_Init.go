// Package utils 工具类
package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB *gorm.DB
)

// InitConfig 初始化应用程序的一些设置
func InitConfig() {
	// 设置设置名为"config"
	viper.SetConfigName("config")
	// 设置读取配置路径为"config"
	viper.AddConfigPath("config")
	// 读取配置文件"
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("app init!")
}

// InitMysql 从viper中获取mysql设置中的dsn即数据源url的值,特别注意这里要把DB传出去, 不然链接无效
func InitMysql() {
	// 打印Mysql日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢SQL阈值
			LogLevel:      logger.Info, // 级别
			Colorful:      true,        // 彩色
		},
	)
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{Logger: newLogger})
	fmt.Println("Mysql init!")
}
