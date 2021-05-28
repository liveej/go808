package test

import (
	"gitee.com/coco/go808/protocol"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestT808_0x8301_EncodeDecode(t *testing.T) {
	message := protocol.T808_0x8301{
		Type: 2,
		Events: []protocol.T808_0x8301_Event{
			{
				ID:      1,
				Content: "事件1",
			},
			{
				ID:      2,
				Content: "事件2",
			},
			{
				ID:      3,
				Content: "事件3",
			},
		},
	}
	data, err := message.Encode()
	if err != nil {
		assert.Error(t, err, "encode error")
	}

	var message2 protocol.T808_0x8301
	_, err = message2.Decode(data)
	if err != nil {
		assert.Error(t, err, "decode error")
	}
	assert.True(t, reflect.DeepEqual(message, message2))
}
