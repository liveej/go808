package protocol

import (
	"go808/errors"
)

// 终端升级结果通知
type T808_0x0108 struct {
	Type   byte
	Result byte
}

func (entity *T808_0x0108) MsgID() MsgID {
	return MsgT808_0x0108
}

func (entity *T808_0x0108) Encode() ([]byte, error) {
	return []byte{entity.Type, entity.Result}, nil
}

func (entity *T808_0x0108) Decode(data []byte) (int, error) {
	if len(data) < 2 {
		return 0, errors.ErrEntityDecodeFail
	}

	entity.Type, entity.Result = data[0], data[1]
	return 2, nil
}
