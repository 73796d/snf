package message

import (
	//"github.com/golang/protobuf/proto"
	"testing"
	"github.com/golang/protobuf/proto"
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

func TestMessage_Package(t *testing.T) {
	msg := NewMessage()

	buf1 := "TestPackage"

	err1 := msg.Package([]byte(buf1))
	if err1 != nil {
		t.Error("package message error")
	}

	buf2, err2 := msg.Unpackage()
	if err2 != nil {
		t.Error("unpackage message error")
	}

	if string(buf2)!= buf1 {
		t.Error("package and unpackage message error", string(buf2), buf1)
	}
	buf1 = "Test_Packageaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	err1 = msg.Package([]byte(buf1))
	if err1 != nil {
		t.Error("package message error")
	}

	buf2, err2 = msg.Unpackage()
	if err2 != nil {
		t.Error("unpackage message error")
	}

	if string(buf2) != buf1 {
		t.Error("package and unpackage message error", string(buf2), buf1, msg.GetUnlen(), msg.GetLen())
	}
}



func TestMessage_GetCmd(t *testing.T) {
 	msg := NewMessage()

	var ni MessageNot
	ni.Id = proto.Uint32(123)
	ni.Text = proto.String("TestPbPackage")
	err := msg.PackagePbmsg(&ni)
	if err != nil {
		t.Error("package protobuf message error")
	}

	var n2 MessageNot
	err2 := msg.UnpackagePbmsg(&n2)
	if err2 != nil {
		t.Error("unpackage protobuf message error")
	}

	if ni.GetId() != n2.GetId() || ni.GetText() != n2.GetText() {
		t.Error("package and unpackage protobuf messgae error", ni.String(), n2.String())
	}

}

func TestMessage_SetGet(t *testing.T)  {
	msg := NewMessage()

	msg.SetLen(111)
	if msg.GetLen() != 111 {
		t.Error("set message data length error, ", msg.GetLen())
	}

	msg.SetId(222)
	if msg.GetId() != 222 {
		t.Error("set message id error, ", msg.GetId())
	}

	msg.SetSeq(333)
	if msg.GetSeq() != 333 {
		t.Error("set message seq error, ", msg.GetSeq())
	}

	msg.SetCmd(444)
	if msg.GetCmd() != 444 {
		t.Error("set message cmd error, ", msg.GetCmd())
	}

	msg.SetMask(3, true)
	if !msg.GetMask(3) {
		t.Error("set message mask 1 error, ", msg.GetMask(3))
	}
	if msg.GetMask(0) {
		t.Error("set message mask 2 error, ", msg.GetMask(0))
	}
	if msg.GetMask(1) {
		t.Error("set message mask 3 error, ", msg.GetMask(1))
	}
}