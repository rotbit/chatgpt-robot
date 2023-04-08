package main

import (
	"fmt"
	"log"
	"sync/atomic"

	"github.com/spf13/viper"
)

// Config 配置内容
type Config struct {
	Key   string
	Port  int32
	Token string
}

var gConfig atomic.Value

// LoadConfig 加载配置
func LoadConfig() {
	// 读取配置文件
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s", err))
	}

	// 解析配置文件
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("解析配置文件失败: %s", err))
	}
	gConfig.Store(config)
}

// GetConfig 获取配置
func GetConfig() Config {
	return gConfig.Load().(Config)
}

// ShowConfig 展示配置信息
func ShowConfig() {
	config := GetConfig()
	log.Println("====================================================")
	log.Printf("OpenAIKey=%v\n", config.Key)
	log.Printf("Token =%v\n", config.Token)
	log.Printf("port=%v\n", config.Port)
	log.Println("====================================================")
}
