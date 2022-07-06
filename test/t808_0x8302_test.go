package test

import (
	"github.com/liveej/go808/protocol"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestT808_0x8302_EncodeDecode(t *testing.T) {
	message := protocol.T808_0x8302{
		Flag:     2,
		Question: "今天星期几？",
		CandidateAnswers: []protocol.T808_0x8302_Answer{
			{
				ID:      1,
				Content: "星期一",
			},
			{
				ID:      2,
				Content: "星期二",
			},
			{
				ID:      3,
				Content: "星期三",
			},
		},
	}
	data, err := message.Encode()
	if err != nil {
		assert.Error(t, err, "encode error")
	}

	var message2 protocol.T808_0x8302
	_, err = message2.Decode(data)
	if err != nil {
		assert.Error(t, err, "decode error")
	}
	assert.True(t, reflect.DeepEqual(message, message2))
}
