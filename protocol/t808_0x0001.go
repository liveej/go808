package protocol

import (
	"bytes"
	"encoding/binary"
)

// 终端应答
type T808_0x0001 struct {
	ResponseMessageSerialNo uint16
	ResponseMessageID       uint16
	ResponseResult          ResponseResult
}

// 获取类型
func (entity *T808_0x0001) Type() Type {
	return TypeT808_0x0001
}

// 消息编码
func (entity *T808_0x0001) Encode() ([]byte, error) {
	var temp [2]byte
	buffer := bytes.NewBuffer(nil)

	// 写入消息序列号
	binary.BigEndian.PutUint16(temp[:], entity.ResponseMessageSerialNo)
	buffer.Write(temp[:])

	// 写入响应消息ID
	binary.BigEndian.PutUint16(temp[:], entity.ResponseMessageID)
	buffer.Write(temp[:])

	// 写入响应结果
	buffer.WriteByte(byte(entity.ResponseResult))
	return buffer.Bytes(), nil
}

// 消息解码
func (entity *T808_0x0001) Decode(data []byte) (int, error) {
	if len(data) < 5 {
		return 0, ErrEntityDecode
	}

	// 读取消息序列号
	var temp [2]byte
	reader := bytes.NewReader(data)
	count, err := reader.Read(temp[:2])
	if err != nil || count != 2 {
		return 0, ErrEntityDecode
	}
	responseMessageSerialNo := binary.BigEndian.Uint16(temp[:2])

	// 读取响应消息ID
	count, err = reader.Read(temp[:2])
	if err != nil || count != 2 {
		return 0, ErrEntityDecode
	}
	responseMessageID := binary.BigEndian.Uint16(temp[:2])

	// 读取响应结果
	count, err = reader.Read(temp[:1])
	if err != nil || count != 1 {
		return 0, ErrEntityDecode
	}
	result := temp[0]

	// 更新Entity信息
	entity.ResponseMessageSerialNo = responseMessageSerialNo
	entity.ResponseMessageID = responseMessageID
	entity.ResponseResult = ResponseResult(result)
	return len(data) - reader.Len(), nil
}
