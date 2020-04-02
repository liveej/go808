package go808

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/funny/link"
	log "github.com/sirupsen/logrus"
	"go808/errors"
	"go808/protocol"
	"io"
)

type Protocol struct{}

// 创建编解码器
func (Protocol) NewCodec(rw io.ReadWriter) (link.Codec, error) {
	codec := &ProtocolCodec{
		w:               rw,
		r:               rw,
		bufferReceiving: bytes.NewBuffer(nil),
	}
	codec.closer, _ = rw.(io.Closer)
	return codec, nil
}

// 编解码器
type ProtocolCodec struct {
	w               io.Writer
	r               io.Reader
	closer          io.Closer
	bufferReceiving *bytes.Buffer
}

// 关闭读写
func (codec *ProtocolCodec) Close() error {
	if codec.closer != nil {
		return codec.closer.Close()
	}
	return nil
}

// 发送消息
func (codec *ProtocolCodec) Send(msg interface{}) error {
	message, ok := msg.(protocol.Message)
	if !ok {
		log.WithFields(log.Fields{
			"reason": errors.ErrInvalidMessage,
		}).Error("[JT/T 808] failed to write message")
		return errors.ErrInvalidMessage
	}

	data, err := message.Encode()
	if err != nil {
		log.WithFields(log.Fields{
			"id":     fmt.Sprintf("0x%x", message.Header.MsgID),
			"reason": err,
		}).Error("[JT/T 808] failed to write message")
		return err
	}

	count, err := codec.w.Write(data)
	if err != nil {
		log.WithFields(log.Fields{
			"id":     fmt.Sprintf("0x%x", message.Header.MsgID),
			"reason": err,
		}).Error("[JT/T 808] failed to write message")
		return err
	}

	log.WithFields(log.Fields{
		"id":    fmt.Sprintf("0x%x", message.Header.MsgID),
		"bytes": count,
	}).Debug("[JT/T 808] write message success")
	return nil
}

// 接收消息
func (codec *ProtocolCodec) Receive() (interface{}, error) {
	var buffer [128]byte
	for {
		count, err := io.ReadAtLeast(codec.r, buffer[:], 1)
		if err != nil {
			return nil, err
		}
		codec.bufferReceiving.Write(buffer[:count])

		if codec.bufferReceiving.Len() == 0 {
			continue
		}
		if codec.bufferReceiving.Len() > 0xffff {
			return nil, errors.ErrBodyTooLong
		}

		data := codec.bufferReceiving.Bytes()
		if data[0] != protocol.PrefixID {
			i := 0
			for ; i < len(data); i++ {
				if data[i] == protocol.PrefixID {
					break
				}
			}
			codec.bufferReceiving.Next(i)
			log.WithFields(log.Fields{
				"data":   hex.EncodeToString(data),
				"reason": errors.ErrNotFoundPrefixID,
			}).Error("[JT/T 808] failed to receive message")
			return nil, errors.ErrNotFoundPrefixID
		}

		if codec.bufferReceiving.Len() == 0 {
			continue
		}
		i := 1
		data = codec.bufferReceiving.Bytes()
		for ; i < len(data); i++ {
			if data[i] == protocol.PrefixID {
				break
			}
		}
		if i == len(data) {
			continue
		}

		var message protocol.Message
		if err = message.Decode(data[:i+1]); err != nil {
			codec.bufferReceiving.Next(i + 1)
			log.WithFields(log.Fields{
				"data":   fmt.Sprintf("0x%x", hex.EncodeToString(data[:i+1])),
				"reason": err,
			}).Error("[JT/T 808] failed to receive message")
			return nil, err
		}
		codec.bufferReceiving.Next(i + 1)

		log.WithFields(log.Fields{
			"id": fmt.Sprintf("0x%x", message.Header.MsgID),
		}).Debug("[JT/T 808] new message received,")
		log.WithFields(log.Fields{
			"data": fmt.Sprintf("0x%x", hex.EncodeToString(data[:i+1])),
		}).Debug("[JT/T 808] show message hex string")
		return message, nil
	}
}
