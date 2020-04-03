package protocol

import (
	"fmt"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"go808/errors"
	"go808/protocol/extra"
	"math"
	"reflect"
	"time"
)

// 纬度类型
type LatitudeType int

const (
	_ LatitudeType = iota
	// 北纬
	NorthLatitudeType = 0
	// 南纬
	SouthLatitudeType = 1
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
type T808_0x0200_Status uint32

// 获取Acc状态
func (status T808_0x0200_Status) GetAccState() bool {
	return GetBitUint32(uint32(status), 0)
}

// 是否正在定位
func (status T808_0x0200_Status) Positioning() bool {
	return GetBitUint32(uint32(status), 1)
}

// 设置南纬
func (status *T808_0x0200_Status) SetSouthLatitude(b bool) {
	SetBitUint32((*uint32)(status), 2, b)
}

// 设置西经
func (status *T808_0x0200_Status) SetWestLongitude(b bool) {
	SetBitUint32((*uint32)(status), 3, b)
}

// 获取纬度类型
func (status T808_0x0200_Status) GetLatitudeType() LatitudeType {
	if GetBitUint32(uint32(status), 2) {
		return SouthLatitudeType
	}
	return NorthLatitudeType
}

// 获取经度类型
func (status T808_0x0200_Status) GetLongitudeType() LongitudeType {
	if GetBitUint32(uint32(status), 3) {
		return WestLongitudeType
	}
	return EastLongitudeType
}

// 汇报位置
type T808_0x0200 struct {
	Alarm     uint32             // 警告
	Status    T808_0x0200_Status // 状态
	Lat       decimal.Decimal    // 纬度
	Lon       decimal.Decimal    // 经度
	Altitude  uint16             // 海拔高度
	Speed     uint16             // 速度
	Direction uint16             // 方向
	Time      time.Time          // 时间
	Extras    []extra.Entity     // 附加信息
}

func (entity *T808_0x0200) MsgID() MsgID {
	return MsgT808_0x0200
}

func (entity *T808_0x0200) Encode() ([]byte, error) {
	writer := NewWriter()

	// 写入警告标志
	writer.WriteUint32(entity.Alarm)

	// 计算经纬度
	mul := decimal.NewFromFloat(1000000)
	lat := entity.Lat.Mul(mul).IntPart()
	if lat < 0 {
		entity.Status.SetSouthLatitude(true)
	}
	lon := entity.Lon.Mul(mul).IntPart()
	if lon < 0 {
		entity.Status.SetWestLongitude(true)
	}

	// 写入状态信息
	writer.WriteUint32(uint32(entity.Status))

	// 写入纬度信息
	writer.WriteUint32(uint32(math.Abs(float64(lat))))

	// 写入经度信息
	writer.WriteUint32(uint32(math.Abs(float64(lon))))

	// 写入海拔高度
	writer.WriteUint16(entity.Altitude)

	// 写入速度信息
	writer.WriteUint16(entity.Speed)

	// 写入方向信息
	writer.WriteUint16(entity.Direction)

	// 写入时间信息
	writer.WriteBcdTime(entity.Time)

	// 写入附加信息
	for i := 0; i < len(entity.Extras); i++ {
		ext := entity.Extras[i]
		if ext == nil || reflect.ValueOf(ext).IsNil() {
			continue
		}
		data := ext.Data()
		full := make([]byte, len(data)+2)
		full[0], full[1] = ext.ID(), byte(len(data))
		copy(full[2:], data)
		writer.Write(full)
	}
	return writer.Bytes(), nil
}

func (entity *T808_0x0200) Decode(data []byte) (int, error) {
	if len(data) < 28 {
		return 0, ErrInvalidBody
	}
	reader := NewReader(data)

	// 读取警告标志
	alarm, err := reader.ReadUint32()
	if err != nil {
		return 0, err
	}

	// 读取状态信息
	status, err := reader.ReadUint32()
	if err != nil {
		return 0, err
	}

	// 读取纬度信息
	latitude, err := reader.ReadUint32()
	if err != nil {
		return 0, err
	}

	// 读取经度信息
	longitude, err := reader.ReadUint32()
	if err != nil {
		return 0, err
	}

	// 读取海拔高度
	altitude, err := reader.ReadUint16()
	if err != nil {
		return 0, err
	}

	// 读取行驶速度
	speed, err := reader.ReadUint16()
	if err != nil {
		return 0, err
	}

	// 读取行驶方向
	direction, err := reader.ReadUint16()
	if err != nil {
		return 0, err
	}

	// 读取上报时间
	time, err := reader.ReadBcdTime()
	if err != nil {
		return 0, err
	}

	// 解码附加信息
	extras := make([]extra.Entity, 0)
	buffer := data[len(data)-reader.Len():]
	for {
		if len(buffer) < 2 {
			break
		}
		id, length := buffer[0], int(buffer[1])
		buffer = buffer[2:]
		if len(buffer) < length {
			return 0, errors.ErrInvalidExtraLength
		}

		extraEntity, count, err := extra.Decode(id, buffer[:length])
		if err != nil {
			if err == errors.ErrTypeNotRegistered {
				buffer = buffer[length:]
				log.WithFields(log.Fields{
					"id": fmt.Sprintf("0x%x", id),
				}).Warn("[JT/T808] unknown T808_0x0200 extra type")
				continue
			}
			return 0, err
		}
		if count != length {
			return 0, errors.ErrInvalidExtraLength
		}
		extras = append(extras, extraEntity)
		buffer = buffer[length:]
	}

	entity.Alarm = alarm
	entity.Status = T808_0x0200_Status(status)
	entity.Altitude = altitude
	entity.Speed = speed
	entity.Direction = direction
	entity.Time = time
	entity.Lat, entity.Lon = getGeoPoint(
		latitude, entity.Status.GetLatitudeType() == SouthLatitudeType,
		longitude, entity.Status.GetLongitudeType() == WestLongitudeType)
	entity.Extras = extras
	return len(data) - reader.Len(), nil
}
