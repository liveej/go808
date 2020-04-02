package protocol

import (
	"go808/errors"
)

// 数据下行透传
type T808_0x8900 struct {
	MessageType byte
	Data        []byte
}

func (entity *T808_0x8900) MsgID() MsgID {
	return MsgT808_0x8900
}

func (entity *T808_0x8900) Encode() ([]byte, error) {
	writer := NewWriter()
	writer.WriteByte(entity.MessageType)
	writer.WriteBytes(entity.Data)
	return writer.Bytes(), nil
}

func (entity *T808_0x8900) Decode(data []byte) (int, error) {
	if len(data) < 1 {
		return 0, errors.ErrEntityDecodeFail
	}
	entity.MessageType, entity.Data = data[0], data[1:]
	return len(data), nil
}
