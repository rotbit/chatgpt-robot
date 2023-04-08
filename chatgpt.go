package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/golang/groupcache"
	"github.com/sashabaranov/go-openai"
)

var cache *groupcache.Group

func init() {
	// 初始化缓存，最多缓存128M数据
	cache = groupcache.NewGroup("chatgpt", 64<<20, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			// 从后端数据源获取数据
			elems := strings.Split(key, "______")
			value, err := chatCompletion(elems[0], elems[1])
			if err != nil {
				log.Printf("chatCompletion Err=%v\n", err)
				return err
			}
			dest.SetString(value)
			return nil
		}))
}

// 构建缓存key
func getKey(key, content string) string {
	return fmt.Sprintf("%v______%v", key, content)
}

// GetChatData 获取OPenAI聊天数据
func GetChatData(appKey, content string) string {
	key := getKey(appKey, content)
	var value string
	err := cache.Get(nil, key, groupcache.StringSink(&value))
	if err != nil {
		return err.Error()
	}
	return value
}

// 从openAI接口取数据
func chatCompletion(key, content string) (string, error) {
	client := openai.NewClient(key)
	rsp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 2048, // 限制最大返回token，提速
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	if len(rsp.Choices) == 0 {
		return "", fmt.Errorf("Get empty response")
	}
	return rsp.Choices[0].Message.Content, nil
}
