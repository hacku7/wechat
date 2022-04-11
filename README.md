# WeChat SDK for Go
>注意：当前版本为v2版本

## 文档 && 例子
[API列表](https://github.com/hacku7/wechat/tree/doc/api)

[Wechat SDK 2.0 文档](https://hacku7.com/wechat)

[Wechat SDK 2.0 例子](https://github.com/hacku7/example)


## 快速开始
```
import "github.com/hacku7/wechat"
```

以下是一个微信公众号处理消息接收以及回复的例子：

```go
//使用memcache保存access_token，也可选择redis或自定义cache
wc := wechat.NewWechat()
memory := cache.NewMemory()
cfg := &offConfig.Config{
    AppID:     "xxx",
    AppSecret: "xxx",
    Token:     "xxx",
    //EncodingAESKey: "xxxx",
    Cache: memory,
}
officialAccount := wc.GetOfficialAccount(cfg)

// 传入request和responseWriter
server := officialAccount.GetServer(req, rw)
//设置接收消息的处理方法
server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {

    //回复消息：演示回复用户发送的消息
    text := message.NewText(msg.Content)
    return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
})

//处理消息接收以及回复
err := server.Serve()
if err != nil {
    fmt.Println(err)
    return
}
//发送回复的消息
server.Send()

```

## 目录说明
- officialaccount: 微信公众号API
- miniprogram: 小程序API
- minigame:小游戏API
- pay:微信支付API
- openplatform:开放平台API
- work:企业微信
- aispeech:智能对话
- doc: api文档

## License

Apache License, Version 2.0
