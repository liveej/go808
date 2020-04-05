package test

import (
	"github.com/stretchr/testify/assert"
	"go808/protocol"
	"reflect"
	"testing"
)

func TestT808_0x8701_EncodeDecode(t *testing.T) {
	message := protocol.T808_0x8701{
		Cmd:  26,
		Data: []byte{1, 3, 5, 7, 9},
	}
	data, err := message.Encode()
	if err != nil {
		assert.Error(t, err, "encode error")
	}

	var message2 protocol.T808_0x8701
	_, err = message2.Decode(data)
	if err != nil {
		assert.Error(t, err, "decode error")
	}
	assert.True(t, reflect.DeepEqual(message, message2))
}
