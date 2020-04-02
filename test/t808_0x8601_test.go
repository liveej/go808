package test

import (
	"github.com/stretchr/testify/assert"
	"go808/protocol"
	"reflect"
	"testing"
)

func TestT808_0x8601_EncodeDecode(t *testing.T) {
	message := protocol.T808_0x8601{
		IDs: []uint32{123, 456, 789},
	}
	data, err := message.Encode()
	if err != nil {
		assert.Error(t, err, "encode error")
	}

	var message2 protocol.T808_0x8601
	_, err = message2.Decode(data)
	if err != nil {
		assert.Error(t, err, "decode error")
	}
	assert.True(t, reflect.DeepEqual(message, message2))
}
