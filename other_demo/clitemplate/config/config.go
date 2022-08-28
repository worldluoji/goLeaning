package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("./config")
    viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil { // 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}