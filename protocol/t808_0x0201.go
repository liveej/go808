package protocol

import (
	"bytes"
	"encoding/binary"
)

// 位置信息查询应答
type T808_0x0201 struct {
	ResponseMessageSerialNo uint16
	GeoPointReport          *T808_0x0200
}

// 获取类型
func (entity *T808_0x0201) Type() Type {
	return TypeT808_0x0201
}

// 消息编码
func (entity *T808_0x0201) Encode() ([]byte, error) {
	return nil, nil
}

// 消息解码
func (entity *T808_0x0201) Decode(data []byte) (int, error) {
	if len(data) <= 3 {
		return 0, ErrEntityDecode
	}

	// 读取消息序列号
	buffer := make([]byte, 255)
	reader := bytes.NewReader(data)
	count, err := reader.Read(buffer[:2])
	if err != nil || count != 2 {
		return 0, ErrEntityDecode
	}
	responseMessageSerialNo := binary.BigEndian.Uint16(buffer[:2])

	// 读取位置信息
	var geoPoint T808_0x0200
	size, err := geoPoint.Decode(data[len(data)-reader.Len():])
	if err != nil {
		return 0, err
	}

	// 更新Entity信息
	entity.GeoPointReport = &geoPoint
	entity.ResponseMessageSerialNo = responseMessageSerialNo
	return len(data) - reader.Len() + size, nil
}
