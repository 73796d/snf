package message

import (
	//"github.com/golang/protobuf/proto"
	"testing"
)

func TestNewMessage(t *testing.T) {
	msg := NewMessage()
	if msg == nil {
		t.Error("new nil message")
	}
	if len(msg.Header) != MessageHeaderLen {
		t.Error("new message header length error, length: ", len(msg.Header))
	}
}

func TestMessage_GetCmd(t *testing.T) {
	
}