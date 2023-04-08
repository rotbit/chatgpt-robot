# chatgpt-robot
chatgpt-robot支持接入微信公众号、支持iphone siri、ipa、mac等苹果系统通过「捷径」接入chatGPT

## 项目信息
开发语言: golang
golang版本要求: go 1.8+以上版本
系统要求: windows/linux/mac

## 安装执行

从仓库获取源码
```
git clone https://github.com/rotbit/chatgpt-robot.git
```

## 源码编译
```
1、设置goproxy
go env -w GOSUMDB=off // Windows  
export GOPROXY=https://goproxy.cn         // macOS 或 Linux

2、进入到chatgpt-robot目录，执行go build
生成的chatgpt-robot为目标程序

```

## 配置解析
```
# config.yaml文件
# 这是一个示例配置文件
key: sk-xxxx  # openAI的key, 在https://platform.openai.com/account/api-keys获取, 仅用于微信公众号
port: 8080    # 服务运行的端口
token: 1234567890  # 微信公众号后台自定义的token
```

## 运行
注意: chatgpt-robot调用openai的接口，部署chatgpt-robot需要要一台海外服务器。
```
到项目目录下，执行nohup ./chatgpt-robot &
```

若本项目对你有用，欢迎star~


