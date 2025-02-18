package utils

import (
	"context"
)

const (
	RedisPre = "lm:websocket"
)

// Publish 发布消息到指定的频道。
// 该函数接收一个上下文、一个频道名称和一个消息字符串作为参数。
// 它将消息发布到Redis中以指定的频道名称为键的频道上。
// 返回值是错误信息，如果没有错误发生，则返回nil。
func Publish(ctx context.Context, channel string, msg string) (err error) {
	// 使用Redis的PUBLISH命令发布消息到频道。
	// RedisPre是预设的前缀，与频道名称组合使用以遵循命名约定。
	err = CACHE.Publish(ctx, channel, msg).Err()
	return
}

// Subscribe 订阅指定的频道并返回接收到的消息。
// 该函数接收一个上下文和一个频道名称作为参数。
// 它订阅Redis中以指定频道名称为键的频道，并等待消息。
// 返回值是一个字符串类型的接收到的消息和一个错误信息。
// 如果没有错误发生，错误信息将为nil。
func Subscribe(ctx context.Context, channel string) (string, error) {
	// 使用Redis的SUBSCRIBE命令订阅频道。
	sub := CACHE.Subscribe(ctx, channel)
	// 使用ReceiveMessage方法阻塞等待，直到接收到一条消息。
	msg, err := sub.ReceiveMessage(ctx)
	// 返回接收到的消息内容和可能的错误信息。
	return msg.Payload, err
}
