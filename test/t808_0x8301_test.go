package test

import (
	"github.com/stretchr/testify/assert"
	"go808/protocol"
	"reflect"
	"testing"
)

func TestT808_0x8301_EncodeDecode(t *testing.T) {
	message := protocol.T808_0x8301{
		Type: 2,
		Events: []protocol.Event{
			{
				EventID: 1,
				Content: "事件1",
			},
			{
				EventID: 2,
				Content: "事件2",
			},
			{
				EventID: 3,
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
