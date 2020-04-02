package protocol

import (
	"go808/errors"
)

// 提问答案
type T808_0x0302 struct {
	AnswerID byte
}

func (entity *T808_0x0302) MsgID() MsgID {
	return MsgT808_0x0302
}

func (entity *T808_0x0302) Encode() ([]byte, error) {
	return []byte{entity.AnswerID}, nil
}

func (entity *T808_0x0302) Decode(data []byte) (int, error) {
	if len(data) < 1 {
		return 0, errors.ErrEntityDecodeFail
	}
	entity.AnswerID = data[0]
	return 1, nil
}
