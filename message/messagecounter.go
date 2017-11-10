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
func (this *MessageCounter)Genarate() uint32 {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()

	this.C += 1
	return this.C
}