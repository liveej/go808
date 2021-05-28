package test

import (
	"gitee.com/coco/go808/protocol"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestT808_0x0805_EncodeDecode(t *testing.T) {
	message := protocol.T808_0x0805{
		ReplyMsgSerialNo: 2342,
		Result:           protocol.T808_0x0805_ResultNotSupported,
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
