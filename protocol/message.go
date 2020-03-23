package protocol

import (
	"bytes"
	"reflect"
	"log"
)

// 消息包
type Message struct {
	Header Header
	Body   Entity
}

// 协议编码
func (message *Message) Encode() ([]byte, error) {
	// 编码消息体
	count := 0
	var err error
	var body []byte
	checkSum := byte(0x00)
	if message.Body != nil && !reflect.ValueOf(message.Body).IsNil() {
		body, err = message.Body.Encode()
		if err != nil {
			return nil, err
		}
	}
	checkSum, count = message.computeChecksum(body, checkSum, count)

	// 编码消息头
	err = message.Header.Property.SetBodySize(uint16(len(body)))
	if err != nil {
		return nil, err
	}
	header, err := message.Header.Encode()
	if err != nil {
		return nil, err
	}
	checkSum, count = message.computeChecksum(header, checkSum, count)

	// 二进制转义
	buffer := bytes.NewBuffer(nil)
	buffer.Grow(count + 2)
	buffer.WriteByte(PrefixID)
	message.write(buffer, header).write(buffer, body).write(buffer, []byte{checkSum})
	buffer.WriteByte(PrefixID)
	return buffer.Bytes(), nil
}

// 协议解码
func (message *Message) Decode(data []byte) error {
	// 检验标志位
	if len(data) < 2 || data[0] != PrefixID || data[len(data)-1] !=PrefixID {
		return ErrInvalidMessage
	}
	data = data[1 : len(data)-1]
	if len(data) == 0 {
		return ErrInvalidMessage
	}

	// 获取校验和
	sum := data[len(data)-1]
	if data[len(data)-2] != EscapeByte {
		data = data[:len(data)-1]
	} else {
		if (data[len(data)-1]) == EscapeByteSufix1 {
			sum = EscapeByte
		} else if data[len(data)-1] == EscapeByteSufix2 {
			sum = PrefixID
		} else {
			return ErrInvalidMessage
		}
		data = data[:len(data)-2]
	}

	// 二进制转义
	checkSum := byte(0x00)
	buffer := make([]byte, 0, len(data))
	for i := 0; i < len(data); {
		b := data[i]
		if b != EscapeByte {
			checkSum = checkSum ^ b
			buffer = append(buffer, b)
			i++
			continue
		}

		if i+1 >= len(data) {
			return ErrInvalidMessage
		}

		b = data[i+1]
		if b == EscapeByteSufix1 {
			checkSum = checkSum ^ EscapeByte
			buffer = append(buffer, EscapeByte)
		} else if b == EscapeByteSufix2 {
			checkSum = checkSum ^ PrefixID
			buffer = append(buffer, PrefixID)
		} else {
			return ErrInvalidMessage
		}
		i += 2
	}

	// 检查校验和
	if len(buffer) == 0 || checkSum != sum {
		return ErrInvalidCheckSum
	}

	// 解码消息头
	if len(buffer) < MessageHeaderSize {
		return ErrInvalidHeader
	}
	var header Header
	err := header.Decode(buffer)
	if err != nil {
		return err
	}
	buffer = buffer[MessageHeaderSize:]

	// 解码消息体
	if uint16(len(buffer)) != header.Property.GetBodySize() {
		log.Printf("[JT/T808] body length mismatch, id: 0x%x, header: %+v",
			header.MsgID, header)
	} else {
		entity, _, err := Decode(uint16(header.MsgID), buffer)
		if err == nil {
			message.Body = entity
		} else {
			log.Printf("[JT/T808] failed to decode message, id: 0x%x, reason: %s",
				header.MsgID, err)
		}
	}
	message.Header = header
	return nil
}

// 写入二进制数据
func (message *Message) write(buffer *bytes.Buffer, data []byte) *Message {
	for _, b := range data {
		if b == PrefixID {
			buffer.WriteByte(EscapeByte)
			buffer.WriteByte(EscapeByteSufix2)
		} else if b == EscapeByte {
			buffer.WriteByte(EscapeByte)
			buffer.WriteByte(EscapeByteSufix1)
		} else {
			buffer.WriteByte(b)
		}
	}
	return message
}

// 校验和累加计算
func (message *Message) computeChecksum(data []byte, checkSum byte, count int) (byte, int) {
	for _, b := range data {
		checkSum = checkSum ^ b
		if b != PrefixID && b != EscapeByte {
			count++
		} else {
			count += 2
		}
	}
	return checkSum, count
}
