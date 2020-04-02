package protocol

import (
	"go808/errors"
)

// 终端鉴权
type T808_0x0102 struct {
	AuthKey string
}

func (entity *T808_0x0102) MsgID() MsgID {
	return MsgT808_0x0102
}

func (entity *T808_0x0102) Encode() ([]byte, error) {
	writer := NewWriter()
	if len(entity.AuthKey) > 0 {
		if err := writer.WritString(entity.AuthKey); err != nil {
			return nil, err
		}
	}
	return writer.Bytes(), nil
}

func (entity *T808_0x0102) Decode(data []byte) (int, error) {
	if len(data) == 0 {
		return 0, errors.ErrEntityDecodeFail
	}

	reader := NewReader(data)
	authKey, err := reader.ReadString()
	if err != nil {
		return 0, errors.ErrEntityDecodeFail
	}
	entity.AuthKey = authKey
	return len(data) - reader.Len(), nil
}
