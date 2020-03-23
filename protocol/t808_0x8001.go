package protocol

import (
	"bytes"
	"encoding/binary"
)

// 平台通用应答
type T808_0x8001 struct {
	MessageSerialNo uint16         // 对应的终端消息的流水号
	MsgID           Type           // 对应的终端消息的ID
	Result          ResponseResult // 处理结果
}

// 获取类型
func (entity *T808_0x8001) Type() Type {
	return TypeT808_0x8001
}

// 消息编码
func (entity *T808_0x8001) Encode() ([]byte, error) {
	var temp [2]byte
	buffer := bytes.NewBuffer(nil)

	// 写入流水号
	binary.BigEndian.PutUint16(temp[:2], entity.MessageSerialNo)
	buffer.Write(temp[:2])

	// 写入消息ID
	binary.BigEndian.PutUint16(temp[:2], uint16(entity.MsgID))
	buffer.Write(temp[:2])

	// 写入处理结果
	buffer.WriteByte(byte(entity.Result))
	return buffer.Bytes(), nil
}

// 消息解码
func (entity *T808_0x8001) Decode(data []byte) (int, error) {
	if len(data) < 5 {
		return 0, ErrEntityDecode
	}

	var temp [2]byte
	reader := bytes.NewReader(data)

	// 读取流水号
	count, err := reader.Read(temp[:2])
	if err != nil || count != 2 {
		return 0, ErrEntityDecode
	}
	messageSerialNo := binary.BigEndian.Uint16(temp[:2])

	// 读取消息ID
	count, err = reader.Read(temp[:2])
	if err != nil || count != 2 {
		return 0, ErrEntityDecode
	}
	msgID := binary.BigEndian.Uint16(temp[:2])

	// 读取处理结果
	b, err := reader.ReadByte()
	if err != nil {
		return 0, ErrEntityDecode
	}

	entity.MsgID = Type(msgID)
	entity.Result = ResponseResult(b)
	entity.MessageSerialNo = messageSerialNo
	return len(data) - reader.Len(), nil
}
