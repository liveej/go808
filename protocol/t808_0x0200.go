package protocol

import (
	"time"
)

// 汇报位置
type T808_0x0200 struct {
	Alarm     uint32            // 警告
	State     T808_0x0200Status // 状态
	Latitude  uint32            // 纬度
	Longitude uint32            // 经度
	Altitude  uint16            // 海拔高度
	Speed     uint16            // 速度
	Direction uint16            // 方向
	Time      time.Time         // 时间
}

// 纬度类型
type LatitudeType int

const (
	_ LatitudeType = iota
	// 北纬
	SouthLatitudeType = 0
	// 南纬
	NorthLatitudeType = 1
)

// 经度类型
type LongitudeType int

const (
	_ LongitudeType = iota
	// 东经
	EastLongitudeType = 0
	// 西经
	WestLongitudeType = 1
)

// 位置状态
type T808_0x0200Status uint32

// 获取Acc状态
func (state T808_0x0200Status) GetAccState() bool {
	value := (state >> 0) & 1
	return value == 1
}

// 是否正在定位
func (state T808_0x0200Status) Positioning() bool {
	value := (state >> 1) & 1
	return value == 1
}

// 获取纬度类型
func (state T808_0x0200Status) GetLatitudeType() LatitudeType {
	value := (state >> 2) & 1
	if value == 1 {
		return NorthLatitudeType
	}
	return SouthLatitudeType
}

// 获取经度类型
func (state T808_0x0200Status) GetLongitudeType() LongitudeType {
	value := (state >> 3) & 1
	if value == 1 {
		return WestLongitudeType
	}
	return EastLongitudeType
}

// 获取类型
func (entity *T808_0x0200) Type() Type {
	return TypeT808_0x0200
}

// 消息编码
func (entity *T808_0x0200) Encode() ([]byte, error) {
	return nil, nil
}

// 消息解码
func (entity *T808_0x0200) Decode(data []byte) (int, error) {
	return 0, nil
	//// 读取警告标志
	//buffer := make([]byte, 6)
	//reader := bytes.NewReader(data)
	//count, err := reader.Read(buffer[:4])
	//if err != nil || count != 4 {
	//	return 0, ErrT808_0x0200
	//}
	//alarm := binary.BigEndian.Uint32(buffer[:4])
	//
	//// 读取状态信息
	//count, err = reader.Read(buffer[:4])
	//if err != nil || count != 4 {
	//	return 0, ErrT808_0x0200
	//}
	//state := binary.BigEndian.Uint32(buffer[:4])
	//
	//// 读取纬度信息
	//count, err = reader.Read(buffer[:4])
	//if err != nil || count != 4 {
	//	return 0, ErrT808_0x0200
	//}
	//latitude := binary.BigEndian.Uint32(buffer[:4])
	//
	//// 读取经度信息
	//count, err = reader.Read(buffer[:4])
	//if err != nil || count != 4 {
	//	return 0, ErrT808_0x0200
	//}
	//longitude := binary.BigEndian.Uint32(buffer[:4])
	//
	//// 读取海拔高度
	//count, err = reader.Read(buffer[:2])
	//if err != nil || count != 2 {
	//	return 0, ErrT808_0x0200
	//}
	//altitude := binary.BigEndian.Uint16(buffer[:2])
	//
	//// 读取行驶速度
	//count, err = reader.Read(buffer[:2])
	//if err != nil || count != 2 {
	//	return 0, ErrT808_0x0200
	//}
	//speed := binary.BigEndian.Uint16(buffer[:2])
	//
	//// 读取行驶方向
	//count, err = reader.Read(buffer[:2])
	//if err != nil || count != 2 {
	//	return 0, ErrT808_0x0200
	//}
	//direction := binary.BigEndian.Uint16(buffer[:2])
	//
	//// 读取上报时间
	//count, err = reader.Read(buffer[:6])
	//if err != nil || count != 6 {
	//	return 0, ErrT808_0x0200
	//}
	//time, err := FromBCDTime(buffer[:6])
	//if err != nil {
	//	return 0, ErrT808_0x0200
	//}
	//
	//// 解码附加信息
	//extras := make([]extra.Entity, 0)
	//buffer = data[len(data)-reader.Len():]
	//for {
	//	if len(buffer) < 2 {
	//		break
	//	}
	//	id, length := buffer[0], int(buffer[1])
	//	buffer = buffer[2:]
	//	if len(buffer) < length {
	//		return 0, ErrInvalidExtraLength
	//	}
	//
	//	extraEntity := extra.MakeEntity(extra.Type(id))
	//	if extraEntity != nil {
	//		count, err := extraEntity.Decode(buffer[:length])
	//		if err != nil {
	//			return 0, err
	//		}
	//		if count > length {
	//			return 0, ErrInvalidExtraLength
	//		}
	//		extras = append(extras, extraEntity)
	//	} else {
	//		log.WithFields(log.Fields{
	//			"id": fmt.Sprintf("0x%x", id),
	//		}).Trace("[JT/T808] unknown T808_0x0200 extra type")
	//	}
	//	buffer = buffer[length:]
	//}
	//
	//// 更新Entity信息
	//entity.Alarm = alarm
	//entity.State = T808_0x0200State(state)
	//entity.Latitude = latitude
	//entity.Longitude = longitude
	//entity.Altitude = altitude
	//entity.Speed = speed
	//entity.Direction = direction
	//entity.Time = time
	//entity.Extras = extras
	//return len(data) - reader.Len(), nil
}
