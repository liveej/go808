package test

import (
	"gitee.com/coco/go808/protocol"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestT808_0x0705_EncodeDecode(t *testing.T) {
	message := protocol.T808_0x0705{
		ReceiveTime: time.Second*13457 + time.Millisecond*999,
		Items: []protocol.T808_0x0705_CAN{
			{},
			{},
		},
	}
	data, err := message.Encode()
	if err != nil {
		assert.Error(t, err, "encode error")
	}

	var message2 protocol.T808_0x0705
	_, err = message2.Decode(data)
	if err != nil {
		assert.Error(t, err, "decode error")
	}
	assert.True(t, reflect.DeepEqual(message, message2))
}
