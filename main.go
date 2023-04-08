package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	LoadConfig()
	ShowConfig()

	r := gin.New()
	// 微信公众号接口
	r.GET("/wechat", wxCheckSign)
	r.POST("/wechat", wxChatMessage)

	// iPhone 捷径接口
	r.POST("/chatgpt/api/completions", completions)

	config := GetConfig()
	if err := r.Run(fmt.Sprintf(":%v", config.Port)); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
