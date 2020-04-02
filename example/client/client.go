package main

import (
	"bytes"
	"fmt"
	"github.com/shopspring/decimal"
	"go808"
	"go808/protocol"
	"net"
	"sync"
	"time"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr(
		"tcp", "127.0.0.1:8808")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go onMessageReceived(conn, &waitGroup)

	// 终端鉴权
	message := protocol.Message{
		Header: protocol.Header{
			IccID:           19901234567,
			MessageSerialNo: 1,
		},
		Body: &protocol.T808_0x0102{
			AuthKey: "12345678",
		},
	}
	data, err := message.Encode()
	if err != nil {
		panic(data)
	}
	if _, err = conn.Write(data); err != nil {
		panic(err)
	}

	// 上报位置
	message = protocol.Message{
		Header: protocol.Header{
			IccID:           19901234567,
			MessageSerialNo: 2,
		},
		Body: &protocol.T808_0x0200{
			Alarm:     2342,
			Status:    0,
			Latitude:  decimal.NewFromFloat(23.562345),
			Longitude: decimal.NewFromFloat(-128.323123),
			Altitude:  2345,
			Speed:     160,
			Direction: 72,
			Time:      time.Unix(time.Now().Unix(), 0),
		},
	}
	data, err = message.Encode()
	if err != nil {
		panic(data)
	}
	if _, err = conn.Write(data); err != nil {
		panic(err)
	}

	// 上传媒体文件
	offset := 0
	limit := 512
	mediaData := make([]byte, 1024*2)
	for {
		if offset >= len(mediaData) {
			break
		}

		if offset+limit > len(mediaData) {
			limit = (offset + limit) - len(mediaData)
		}

		sum := len(mediaData) / limit
		if len(mediaData)%limit > 0 {
			sum += 1
		}
		seq := (offset + limit) / limit
		if (offset+limit)%limit > 0 {
			seq += 1
		}

		header := protocol.Header{
			IccID:           19901234567,
			MessageSerialNo: 3,
		}
		header.Packet = &protocol.Packet{
			Sum: uint16(sum),
			Seq: uint16(seq),
		}
		message = protocol.Message{
			Header: header,
			Body: &protocol.T808_0x0801{
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
				Packet: bytes.NewReader(mediaData[offset : offset+limit]),
			},
		}
		data, err = message.Encode()
		if err != nil {
			panic(data)
		}
		if _, err = conn.Write(data); err != nil {
			panic(err)
		}
		offset += limit
	}

	waitGroup.Wait()
}

func onMessageReceived(conn *net.TCPConn, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	var p go808.Protocol
	codec, err := p.NewCodec(conn)
	if err != nil {
		panic(err)
	}

	for {
		msg, err := codec.Receive()
		if err != nil {
			conn.Close()
			break
		}

		message := msg.(protocol.Message)
		if message.Header.MsgID == protocol.MsgT808_0x8104 {
			message = protocol.Message{
				Header: protocol.Header{
					IccID:           19901234567,
					MessageSerialNo: 1000,
				},
				Body: &protocol.T808_0x0104{
					ResponseMessageSerialNo: message.Header.MessageSerialNo,
					Params: []*protocol.Param{
						new(protocol.Param).SetByte(0x0084, 24),
						new(protocol.Param).SetBytes(0x0110, []byte{1, 2, 3, 4, 5, 6, 7, 8}),
						new(protocol.Param).SetUint16(0x0031, 100),
						new(protocol.Param).SetUint32(0x0046, 64000),
						new(protocol.Param).SetString(0x0083, "车牌号码"),
					},
				},
			}
			data, err := message.Encode()
			if err != nil {
				panic(data)
			}
			if _, err = conn.Write(data); err != nil {
				panic(err)
			}
		} else if message.Header.MsgID == protocol.MsgT808_0x8800 {
			fmt.Println("===========================> 媒体上传成功 <===========================")
		}
	}
}
