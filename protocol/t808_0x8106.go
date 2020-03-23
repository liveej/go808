package protocol

import (
	"bytes"
	"encoding/binary"
)

// 查询终端参数
type T808_0x8106 struct {
	Params []uint32
}

// 获取类型
func (entity *T808_0x8106) Type() Type {
	return TypeT808_0x8106
}

// 消息编码
func (entity *T808_0x8106) Encode() ([]byte, error) {
	var temp [4]byte
	buffer := bytes.NewBuffer(nil)
	buffer.WriteByte(byte(len(entity.Params)))
	for _, param := range entity.Params {
		binary.BigEndian.PutUint32(temp[:], param)
		buffer.Write(temp[:])
	}
	return buffer.Bytes(), nil
}

// 消息解码
func (entity *T808_0x8106) Decode(data []byte) (int, error) {
	return 0, nil
}
