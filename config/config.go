package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	//viper库来读取项目配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")
	//读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("读取配置文件出错：%s", err.Error()))
	}
	fmt.Printf("服务端口:%s\n", viper.GetString("server.port"))
}
