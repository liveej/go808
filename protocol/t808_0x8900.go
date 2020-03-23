package protocol

import (
	"bytes"
)

// 数据下行透传
type T808_0x8900 struct {
	MessageType byte
	Data        []byte
}

// 获取类型
func (entity *T808_0x8900) Type() Type {
	return TypeT808_0x8900
}

// 消息编码
func (entity *T808_0x8900) Encode() ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteByte(entity.MessageType)
	buffer.Write(entity.Data)
	return buffer.Bytes(), nil
}

// 消息解码
func (entity *T808_0x8900) Decode(data []byte) (int, error) {
	if len(data) < 1 {
		return 0, ErrEntityDecode
	}
	return 0, nil
}
