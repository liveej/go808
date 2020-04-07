package test

import (
	"github.com/stretchr/testify/assert"
	"go808/protocol"
	"testing"
)

func TestT808_0x0805_EncodeDecode(t *testing.T) {
	message := protocol.T808_0x0805{
		AnswerMessageSerialNo: 2342,
		Result:                protocol.T808_0x0805_ResultNotSupported,
		MediaIDs: []uint32{
			1232313,
			213214,
			213123123,
		},
	}
	data, err := message.Encode()
	if err != nil {
		assert.Error(t, err, "encode error")
	}

	var message2 protocol.T808_0x0805
	_, err = message2.Decode(data)
	if err != nil {
		assert.Error(t, err, "decode error")
	}
}
