package protocol

import (
	"bytes"
	"encoding/binary"
)

// 删除多边形区域
type T808_0x8605 struct {
	IDs []uint32
}

// 获取类型
func (entity *T808_0x8605) Type() Type {
	return TypeT808_0x8605
}

// 消息编码
func (entity *T808_0x8605) Encode() ([]byte, error) {
	buffer := bytes.NewBuffer(nil)

	// 写入ID总数
	buffer.WriteByte(byte(len(entity.IDs)))

	// 写入ID列表
	var tmp [4]byte
	for _, id := range entity.IDs {
		binary.BigEndian.PutUint32(tmp[:4], id)
		buffer.Write(tmp[:4])
	}

	return buffer.Bytes(), nil
}

// 消息解码
func (entity *T808_0x8605) Decode(data []byte) (int, error) {
	return 0, nil
}
