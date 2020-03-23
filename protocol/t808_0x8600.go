package protocol

import (
	"bytes"
	"encoding/binary"
)

// 区域属性
type AreaAttribute uint16

// 设置东经
func (attr *AreaAttribute) SetEastLon(b bool) {
	if !b {
		*attr |= 1 << 7
	} else {
		*attr &= ^AreaAttribute(1 << 7)
	}
}

// 设置北纬
func (attr *AreaAttribute) SetNorthLat(b bool) {
	if !b {
		*attr |= 1 << 6
	} else {
		*attr &= ^AreaAttribute(1 << 6)
	}
}

// 设置离开报警平台
func (attr *AreaAttribute) SetExitAlarm(b bool) {
	if b {
		*attr |= 1 << 5
	} else {
		*attr &= ^AreaAttribute(1 << 5)
	}
}

// 设置进入报警平台
func (attr *AreaAttribute) SetEnterAlarm(b bool) {
	if b {
		*attr |= 1 << 3
	} else {
		*attr &= ^AreaAttribute(1 << 3)
	}
}

// 区域动作
type AreaAction byte

var (
	AreaActionUpdate AreaAction = 0
	AreaActionAdd    AreaAction = 1
	AreaActionEdit   AreaAction = 2
)

// 圆形区域
type CircleArea struct {
	ID        uint32
	Attribute AreaAttribute
	Lat       uint32
	Lon       uint32
	Radius    uint32
}

// 设置圆形区域
type T808_0x8600 struct {
	Action AreaAction
	Items  []CircleArea
}

// 获取类型
func (entity *T808_0x8600) Type() Type {
	return TypeT808_0x8600
}

// 消息编码
func (entity *T808_0x8600) Encode() ([]byte, error) {
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

		// 写入中心点纬度
		binary.BigEndian.PutUint32(tmp[:4], item.Lat)
		buffer.Write(tmp[:4])

		// 写入中心点经度
		binary.BigEndian.PutUint32(tmp[:4], item.Lon)
		buffer.Write(tmp[:4])

		// 写入半径
		binary.BigEndian.PutUint32(tmp[:4], item.Radius)
		buffer.Write(tmp[:4])
	}
	return buffer.Bytes(), nil
}

// 消息解码
func (entity *T808_0x8600) Decode(data []byte) (int, error) {
	return 0, nil
}
