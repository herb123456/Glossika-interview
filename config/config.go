package config

import (
	"github.com/spf13/viper"
	"log/slog"
)

// 初始化Viper庫
func init() {
	viper.SetConfigName("config")   // 配置文件名稱 (無擴展名)
	viper.SetConfigType("yaml")     // 配置文件類型
	viper.AddConfigPath("./config") // 配置文件路徑，"." 代表當前目錄

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		slog.Error("Error reading config file, " + err.Error())
		panic("Error reading config file, " + err.Error())
	}
}

// Get 提供一個獲取配置的公共接口，返回一個 interface{} 類型的配置值
func Get(key string) interface{} {
	return viper.Get(key)
}

// GetString 專門用於獲取字符串類型的配置值
func GetString(key string) string {
	return viper.GetString(key)
}
