package protocol

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"strconv"
)

// 消息头
type Header struct {
	MsgID           Type
	Property        Property
	ICCID           uint64
	MessageSerialNo uint16
}

// 协议编码
func (header *Header) Encode() ([]byte, error) {
	// 写入消息ID
	var tmp [8]byte
	buffer := bytes.NewBuffer(nil)
	binary.BigEndian.PutUint16(tmp[:2], uint16(header.MsgID))
	buffer.Write(tmp[:2])

	// 写入消息体属性
	binary.BigEndian.PutUint16(tmp[:2], uint16(header.Property))
	buffer.Write(tmp[:2])

	// 写入终端号码
	simID, err := hex.DecodeString(strconv.FormatUint(header.ICCID, 0))
	if err != nil {
		return nil, ErrInvalidHeader
	}
	buffer.Write(simID[2:])

	// 写入消息流水号
	binary.BigEndian.PutUint16(tmp[:2], uint16(header.MessageSerialNo))
	buffer.Write(tmp[:2])
	return buffer.Bytes(), nil
}

// 协议解码
func (header *Header) Decode(data []byte) error {
	if len(data) < MessageHeaderSize {
		return ErrInvalidHeader
	}

	// 读取消息ID
	var buffer [6]byte
	reader := bytes.NewReader(data)
	count, err := reader.Read(buffer[:2])
	if err != nil || count != 2 {
		return ErrInvalidHeader
	}
	msgID := binary.BigEndian.Uint16(buffer[:2])

	// 读取消息体属性
	count, err = reader.Read(buffer[:2])
	if err != nil || count != 2 {
		return ErrInvalidHeader
	}
	property := binary.BigEndian.Uint16(buffer[:2])

	// 读取终端号码
	count, err = reader.Read(buffer[:6])
	if err != nil || count != 6 {
		return ErrInvalidHeader
	}
	simID, _ := strconv.ParseUint(hex.EncodeToString(buffer[:6]), 10, 64)

	// 读取消息流水号
	count, err = reader.Read(buffer[:2])
	if err != nil || count != 2 {
		return ErrInvalidHeader
	}
	serialNo := binary.BigEndian.Uint16(buffer[:2])

	// 更新消息头信息
	header.MsgID = Type(msgID)
	header.ICCID = simID
	header.Property = Property(property)
	header.MessageSerialNo = serialNo
	return nil
}
