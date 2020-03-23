package protocol

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestT808_0x0107_EncodeDecode(t *testing.T) {
	message := T808_0x0107{
		TerminalType:    123,
		ManufactureID:   "wer",
		Model:           "1231231",
		TerminalID:      "t800",
		HardwareVersion: "v0.1.1",
		SoftwareVersion: "v0.1.2",
		GNSSProperty:    23,
		COMMProperty:    34,
	}
	data, err := message.Encode()
	if err != nil {
		assert.Error(t, err, "encode error")
	}

	var message2 T808_0x0107
	_, err = message2.Decode(data)
	if err != nil {
		assert.Error(t, err, "decode error")
	}
	assert.True(t, reflect.DeepEqual(message, message2))
}
