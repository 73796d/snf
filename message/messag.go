package message

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
)

// 消息位置
const (
	MessageIdBit    = 0  //消息来源或者目的id
	MessageUnlenBit = 4  //未压缩前消息体长度，可为0，和len相等表示没有压缩
	MessageLenBit   = 6  //消息体长度，可为0
	MessageCmdBit   = 8  //消息命令字
	MessageSeqBit   = 10 //消息序号
	MessageRetBit   = 12 //消息返回值，消息返回时使用
	MessageMaskBit  = 14 //一些标志

	MessageHeaderLen = 16 //消息长度

)

type Message struct {
	Header []byte // 消息头
	Data   []byte // 消息体
}

func NewMessage()*Message  {
	msg := new(Message)
	msg.Header = make([]byte, MessageHeaderLen)
	return msg
	/*
	return &Message{
		Header:make([]byte, MessageHeaderLen),
	}
	*/
}


