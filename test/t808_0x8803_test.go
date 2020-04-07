package test

import (
	"github.com/stretchr/testify/assert"
	"go808/protocol"
	"reflect"
	"testing"
	"time"
)

func TestT808_0x8803_EncodeDecode(t *testing.T) {
	message := protocol.T808_0x8803{
		Type:       protocol.T808_0x0801_MediaTypeAudio,
		ChannelID:  56,
		Event:      87,
		StartTime:  time.Unix(time.Now().Unix(), 0),
		EndTime:    time.Unix(time.Now().Unix(), 0),
		RemoveFlag: 1,
	}
	data, err := message.Encode()
	if err != nil {
		assert.Error(t, err, "encode error")
	}

	var message2 protocol.T808_0x8803
	_, err = message2.Decode(data)
	if err != nil {
		assert.Error(t, err, "decode error")
	}
	assert.True(t, reflect.DeepEqual(message, message2))
}
