package protocol

import (
	"bytes"
	"encoding/binary"
)

// 顶点
type Vertex struct {
	Lat uint32
	Lon uint32
}

// 多边形区域
type PolygonArea struct {
	ID        uint32
	Attribute AreaAttribute
	Vertexes  []Vertex
}

// 设置多边形区域
type T808_0x8604 struct {
	Action AreaAction
	Items  []PolygonArea
}

// 获取类型
func (entity *T808_0x8604) Type() Type {
	return TypeT808_0x8604
}

// 消息编码
func (entity *T808_0x8604) Encode() ([]byte, error) {
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

		// 写入顶点总数
		binary.BigEndian.PutUint16(tmp[:2], uint16(len(item.Vertexes)))
		buffer.Write(tmp[:2])

		// 写入顶点信息
		for _, vertex := range item.Vertexes {
			// 写入纬度
			binary.BigEndian.PutUint32(tmp[:4], vertex.Lat)
			buffer.Write(tmp[:4])

			// 写入经度
			binary.BigEndian.PutUint32(tmp[:4], vertex.Lon)
			buffer.Write(tmp[:4])
		}
	}
	return buffer.Bytes(), nil
}

// 消息解码
func (entity *T808_0x8604) Decode(data []byte) (int, error) {
	return 0, nil
}
