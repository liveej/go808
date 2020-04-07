package test

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"go808/protocol"
	"reflect"
	"testing"
	"time"
)

func TestT808_0x0802_EncodeDecode(t *testing.T) {
	message := protocol.T808_0x0802{
		AnswerMessageSerialNo: 1234,
		Items: []protocol.T808_0x0802_Item{
			{
				MediaID:   345,
				Type:      protocol.T808_0x0801_MediaTypeImage,
				ChannelID: 24,
				Event:     12,
				Location: protocol.T808_0x0200{
					Alarm:     2342,
					Status:    8,
					Lat:       decimal.NewFromFloat(23.562345),
					Lon:       decimal.NewFromFloat(-128.323123),
					Altitude:  2345,
					Speed:     160,
					Direction: 72,
					Time:      time.Unix(time.Now().Unix(), 0),
				},
			},
		},
	}
	data, err := message.Encode()
	if err != nil {
		assert.Error(t, err, "encode error")
	}

	var message2 protocol.T808_0x0802
	_, err = message2.Decode(data)
	if err != nil {
		assert.Error(t, err, "decode error")
	}
	assert.True(t, reflect.DeepEqual(message, message2))
}
