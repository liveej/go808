package test

import (
	"bytes"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"go808/protocol"
	"testing"
	"time"
)

func TestT808_0x0801_EncodeDecode(t *testing.T) {
	message := protocol.T808_0x0801{
		MediaID:   1024,
		Type:      protocol.MediaTypeAudio,
		Coding:    protocol.MediaCodingJPEG,
		Event:     13,
		ChannelID: 28,
		Location: protocol.T808_0x0200{
			Alarm:     2342,
			Status:    0,
			Latitude:  decimal.NewFromFloat(23.562345),
			Longitude: decimal.NewFromFloat(-128.323123),
			Altitude:  2345,
			Speed:     160,
			Direction: 72,
			Time:      time.Unix(time.Now().Unix(), 0),
		},
		Packet: bytes.NewReader(make([]byte, 512)),
	}
	data, err := message.Encode()
	if err != nil {
		assert.Error(t, err, "encode error")
	}

	var message2 protocol.T808_0x0801
	err = message2.DecodePacket(data)
	if err != nil {
		assert.Error(t, err, "decode error")
	}
}