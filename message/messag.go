package message

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
)

// 消息标志位
const (
	MessageMaskDisconn = 0 //是否断开连接
	MessageMaskNotify  = 1 //是否通知
)

// 消息头位位置
const (
	MessageIdBit    = 0  //消息来源或者目的id
	MessageUnlenBit = 4  //未压缩前消息体长度，可为0，和len相等表示没有压缩
	MessageLenBit   = 6  //消息体长度，可为0
	MessageCmdBit   = 8  //消息命令字
	MessageSeqBit   = 10 //消息序号
	MessageRetBit   = 12 //消息返回值，消息返回时使用
	MessageMaskBit  = 14 //一些标志

	MessageHeaderLen = 16 //消息头长度

)

/*
 * 消息（消息头 + 消息体）
 */
type Message struct {
	Header []byte // 消息头
	Data   []byte // 消息体
}

// 新建消息
func NewMessage() *Message {
	msg := new(Message)
	msg.Header = make([]byte, MessageHeaderLen)
	return msg
	/*
		return &Message{
			Header:make([]byte, MessageHeaderLen),
		}
	*/
}

// 打包原生字符串
func (m *Message) Package(buf []byte) error {
	l := len(buf)
	if l == 0 {
		return nil
	}
	var b bytes.Buffer
	c := false

	// 小于指定长度不用检查是否需要压缩
	if l > 10 {
		w := zlib.NewWriter(&b)
		w.Write(buf)
		w.Close()
		c = true
	}

	// 压缩后长度比原来小， 就保存压缩数据
	if c && b.Len() < l {
		m.SetUnlen(uint16(l))
		m.SetLen(uint16(b.Len()))
		m.Data = make([]byte, b.Len())
		copy(m.Data[:], b.Bytes())
	} else {
		m.SetUnlen(uint16(l))
		m.SetLen(uint16(l))
		m.Data = make([]byte, l)
		copy(m.Data[:], buf)
	}

	return nil

}

// 解包原生字符串
func (m *Message) Unpackage() ([]byte, error) {
	if m.GetLen() == 0 {
		return []byte(""), nil
	}
	if m.GetLen() == m.GetUnlen() {
		data := make([]byte, m.GetLen())
		copy(data[:], m.Data)
		return data, nil
	} else if m.GetLen() < m.GetUnlen() {
		var b bytes.Buffer
		b.Write(m.Data)
		r, err := zlib.NewReader(&b)
		if err != nil {
			return []byte(""), err
		}
		defer r.Close()

		data := make([]byte, m.GetUnlen())
		l, _ := r.Read(data)
		if l != int(m.GetUnlen()) {
			return []byte(""), errors.New("uncompress erro")
		}
		return data, nil
	} else {
		return []byte(""), errors.New("message unpackage erro")
	}
}

// 打包protubuf消息
func (m *Message) PackagePbmsg(pb proto.Message) error {
	buf, err := proto.Marshal(pb)
	if err != nil {
		return err
	}
	return m.Package(buf)
}

// 解包protobuf消息
func (m *Message) UnpackagePbmsg(pb proto.Message) error {
	data, err := m.Unpackage()
	if err != nil {
		return err
	}
	return proto.Unmarshal(data, pb)
}

// 根据消息长度初始化消息体内存
func (m *Message) InitData() {
	if m.GetLen() == 0 {
		return
	}
	m.Data = make([]byte, m.GetLen())
}

// 获取t位置开始的两个字节的数值
func (m *Message) get16(t uint32) uint16 {
	buf := bytes.NewBuffer(make([]byte, 0, MessageHeaderLen))
	buf.Write(m.Header[t : t+2])

	i16 := make([]byte, 2)
	buf.Read(i16)
	return binary.BigEndian.Uint16(i16)
}

// 将id设置到t位置开始的两个字节内
func (m *Message) set16(t uint32, id uint16) {
	binary.BigEndian.PutUint16(m.Header[t:t+2], id)
}

// 获取t位置开始的四个字节的数值
func (m *Message) get32(t uint32) uint32 {
	buf := bytes.NewBuffer(make([]byte, 0, MessageHeaderLen))
	buf.Write(m.Header[t : t+4])

	i32 := make([]byte, 4)
	buf.Read(i32)
	return binary.BigEndian.Uint32(i32)
}

// 将id设置到t位置开始的四个字节内
func (m *Message) set32(t uint32, id uint32) {
	binary.BigEndian.PutUint32(m.Header[t:t+4], id)
}

// 获取消息id(消息头前四位0 - 3)
func (m *Message) GetId() uint32 {
	return m.get32(MessageIdBit)
}

// 设置消息id
func (m *Message) SetId(id uint32) {
	m.set32(MessageIdBit, id)
}

// 获取命令（消息头8 - 9）
func (m *Message) GetCmd() uint16 {
	return m.get16(MessageCmdBit)
}
func (m *Message) SetCmd(cmd uint16) {
	m.set16(MessageCmdBit, cmd)
}

// 获取消息序号（消息头10 - 11）
func (m *Message) GetSeq() uint16 {
	return m.get16(MessageSeqBit)
}

func (m *Message) SetSeq(seq uint16) {
	m.set16(MessageSeqBit, seq)
}

// 获取消息未压缩长度（消息头4 - 5）
func (m *Message) GetUnlen() uint16 {
	return m.get16(MessageUnlenBit)
}

func (m *Message) SetUnlen(len uint16) {
	m.set16(MessageUnlenBit, len)
}

// 获取消息体长度（消息头 6 - 7）
func (m *Message) GetLen() uint16 {
	return m.get16(MessageLenBit)
}

func (m *Message) SetLen(len uint16) {
	m.set16(MessageLenBit, len)
}

// 获取消息返回值（消息头 12 - 13）
func (m *Message) GetRet() uint16 {
	return m.get16(MessageRetBit)
}

func (m *Message) SetRet(ret uint16) {
	m.set16(MessageRetBit, ret)
}

// 获取消息标志（消息头 14 - 16）
func (m *Message) GetMask(mask uint16) bool {
	i := m.get16(MessageMaskBit)
	return (i & (1 << mask)) != 0
}

func (m *Message) SetMask(mask uint16, b bool) {
	i := m.get16(MessageMaskBit)

	if b {
		i |= 1 << mask
	} else {
		i &= ^(1 << mask)
	}

	m.set16(MessageMaskBit, i)
}

// 打印的借口实现
func (m *Message) String() string {
	return fmt.Sprintf("id: %d, unlen: %d, len: %d, cmd: %d, seq: %d, ret: %d", m.GetId(), m.GetUnlen(), m.GetLen(), m.GetCmd(), m.GetSeq(), m.GetRet())
}

// 消息的拷贝
func (m *Message) Copy() *Message {
	newMsg := NewMessage()
	newMsg.Data = make([]byte, m.GetLen())
	copy(newMsg.Header[:], m.Header)
	copy(newMsg.Data[:], m.Data)
	return newMsg
}
