package protocol

import (
	"go808/errors"
)

// 临时位置跟踪控制
type T808_0x8202 struct {
	Interval uint16
	Expire   uint32
}

func (entity *T808_0x8202) MsgID() MsgID {
	return MsgT808_0x8202
}

func (entity *T808_0x8202) Encode() ([]byte, error) {
	writer := NewWriter()

	// 写入时间间隔
	writer.WriteUint16(entity.Interval)

	// 写入跟踪有效期
	writer.WriteUint32(entity.Expire)
	return writer.Bytes(), nil
}

func (entity *T808_0x8202) Decode(data []byte) (int, error) {
	if len(data) < 4 {
		return 0, errors.ErrEntityDecodeFail
	}
	reader := NewReader(data)

	// 读取时间间隔
	var err error
	entity.Interval, err = reader.ReadUint16()
	if err != nil {
		return 0, errors.ErrEntityDecodeFail
	}

	// 读取跟踪有效期
	entity.Expire, err = reader.ReadUint32()
	if err != nil {
		return 0, errors.ErrEntityDecodeFail
	}
	return len(data) - reader.Len(), nil
}
