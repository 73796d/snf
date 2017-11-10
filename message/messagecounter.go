package message

import (
	"sync"
)

// 消息序列号计数器，用于判断数据包的连续性
type MessageCounter struct {
	C     uint32
	Mutex *sync.Mutex
}

// 创建一个新的计数器
func NewMessageCounter() *MessageCounter {
	messageCounter := new(MessageCounter)
	messageCounter.C = 0
	messageCounter.Mutex = new(sync.Mutex)
	return messageCounter
}

// 返回计数器自增后的序列号
func (mc *MessageCounter) Genarate() uint32 {
	mc.Mutex.Lock()
	defer mc.Mutex.Unlock()

	mc.C += 1
	return mc.C
}
