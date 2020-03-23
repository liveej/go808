package protocol

import (
	"bytes"
	"encoding/binary"
)

// 矩形区域
type RectArea struct {
	ID             uint32
	Attribute      AreaAttribute
	LeftTopLat     uint32
	LeftTopLon     uint32
	RightBottomLat uint32
	RightBottomLon uint32
}

// 设置矩形区域
type T808_0x8602 struct {
	Action AreaAction
	Items  []RectArea
}

// 获取类型
func (entity *T808_0x8602) Type() Type {
	return TypeT808_0x8602
}

// 消息编码
func (entity *T808_0x8602) Encode() ([]byte, error) {
	buffer := bytes.NewBuffer(nil)

	// 写入设置属性
	buffer.WriteByte(byte(entity.Action))

	// 写入区域总数
	buffer.WriteByte(byte(len(entity.Items)))

	// 写入区域信息
	var tmp [4]byte
	for _, item := range entity.Items {
		// 写入区域ID
		binary.BigEndian.PutUint32(tmp[:4], item.ID)
		buffer.Write(tmp[:4])

		// 写入区域属性
		binary.BigEndian.PutUint16(tmp[:2], uint16(item.Attribute))
		buffer.Write(tmp[:2])

		// 写入左上角纬度
		binary.BigEndian.PutUint32(tmp[:4], item.LeftTopLat)
		buffer.Write(tmp[:4])

		// 写入左上角经度
		binary.BigEndian.PutUint32(tmp[:4], item.LeftTopLon)
		buffer.Write(tmp[:4])

		// 写入右下角纬度
		binary.BigEndian.PutUint32(tmp[:4], item.RightBottomLat)
		buffer.Write(tmp[:4])

		// 写入右下角经度
		binary.BigEndian.PutUint32(tmp[:4], item.RightBottomLon)
		buffer.Write(tmp[:4])
	}
	return buffer.Bytes(), nil
}

// 消息解码
func (entity *T808_0x8602) Decode(data []byte) (int, error) {
	return 0, nil
}
