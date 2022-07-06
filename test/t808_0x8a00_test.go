package test

import (
	"github.com/liveej/go808/protocol"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestT808_0x8A00_EncodeDecode(t *testing.T) {
	privateKey, err := GetTestPrivateKey()
	if err != nil {
		t.Error(err)
	}

	message := protocol.T808_0x8A00{
		PublicKey: &privateKey.PublicKey,
	}
	data, err := message.Encode()
	if err != nil {
		assert.Error(t, err, "encode error")
	}

	var message2 protocol.T808_0x8A00
	_, err = message2.Decode(data)
	if err != nil {
		assert.Error(t, err, "decode error")
	}
	assert.True(t, reflect.DeepEqual(message, message2))
}
