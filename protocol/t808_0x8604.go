package protocol

import (
	"github.com/shopspring/decimal"
	"math"
	"time"
)

// 顶点
type Vertex struct {
	Lat decimal.Decimal
	Lon decimal.Decimal
}

// 多边形区域
type PolygonArea struct {
	ID        uint32
	Attribute AreaAttribute
	StartTime time.Time
	EndTime   time.Time
	MaxSpeed  uint16
	Duration  byte
	Vertexes  []Vertex
}

// 设置多边形区域
type T808_0x8604 struct {
	Action AreaAction
	Items  []PolygonArea
}

func (entity *T808_0x8604) MsgID() MsgID {
	return MsgT808_0x8604
}

func (entity *T808_0x8604) Encode() ([]byte, error) {
	writer := NewWriter()

	// 写入设置属性
	writer.WriteByte(byte(entity.Action))

	// 写入区域总数
	writer.WriteByte(byte(len(entity.Items)))

	// 写入区域信息
	for i := range entity.Items {
		item := &entity.Items[i]

		// 写入区域ID
		writer.WriteUint32(item.ID)

		// 计算经纬度
		mul := decimal.NewFromFloat(1000000)
		vertexes := make([]Vertex, 0, len(item.Vertexes))
		for j := range item.Vertexes {
			vertex := &item.Vertexes[j]
			lat := vertex.Lat.Mul(mul)
			if lat.Cmp(decimal.Zero) < 0 {
				item.Attribute.SetSouthLatitude(true)
			}
			lon := vertex.Lon.Mul(mul)
			if lon.Cmp(decimal.Zero) < 0 {
				item.Attribute.SetWestLongitude(true)
			}
			vertexes = append(vertexes, Vertex{
				Lat: lat,
				Lon: lon,
			})
		}

		// 写入区域属性
		writer.WriteUint16(uint16(item.Attribute))

		// 写入时间参数
		if item.Attribute&1 == 1 {
			// 写入开始时间
			writer.WriteBcdTime(item.StartTime)

			// 写入结束时间
			writer.WriteBcdTime(item.EndTime)

			// 写入最高速度
			writer.WriteUint16(item.MaxSpeed)

			// 写入持续时间
			writer.WriteByte(item.Duration)
		}

		// 写入顶点总数
		writer.WriteUint16(uint16(len(item.Vertexes)))

		// 写入顶点信息
		for _, vertex := range vertexes {
			// 写入纬度
			writer.WriteUint32(uint32(math.Abs(float64(vertex.Lat.IntPart()))))

			// 写入经度
			writer.WriteUint32(uint32(math.Abs(float64(vertex.Lon.IntPart()))))
		}
	}
	return writer.Bytes(), nil
}

func (entity *T808_0x8604) Decode(data []byte) (int, error) {
	if len(data) < 16 {
		return 0, ErrInvalidBody
	}
	reader := NewReader(data)

	// 读取设置属性
	action, err := reader.ReadByte()
	if err != nil {
		return 0, err
	}
	entity.Action = AreaAction(action)

	// 读取区域总数
	count, err := reader.ReadByte()
	if err != nil {
		return 0, err
	}

	// 读取区域信息
	entity.Items = make([]PolygonArea, 0, count)
	for i := 0; i < int(count); i++ {
		var area PolygonArea

		// 读取区域ID
		area.ID, err = reader.ReadUint32()
		if err != nil {
			return 0, err
		}

		// 读取区域属性
		attribute, err := reader.ReadUint16()
		if err != nil {
			return 0, err
		}
		area.Attribute = AreaAttribute(attribute)

		// 读取时间参数
		if area.Attribute&1 == 1 {
			// 读取开始时间
			area.StartTime, err = reader.ReadBcdTime()
			if err != nil {
				return 0, err
			}

			// 读取结束时间
			area.EndTime, err = reader.ReadBcdTime()
			if err != nil {
				return 0, err
			}

			// 读取最高速度
			area.MaxSpeed, err = reader.ReadUint16()
			if err != nil {
				return 0, err
			}

			// 读取持续时间
			area.Duration, err = reader.ReadByte()
			if err != nil {
				return 0, err
			}
		}

		// 读取顶点总数
		vertexes, err := reader.ReadUint16()
		if err != nil {
			return 0, err
		}
		area.Vertexes = make([]Vertex, 0, int(vertexes))

		// 读取顶点列表
		for j := 0; j < int(vertexes); j++ {
			var vertex Vertex

			// 读取纬度
			lat, err := reader.ReadUint32()
			if err != nil {
				return 0, err
			}

			// 读取经度
			lon, err := reader.ReadUint32()
			if err != nil {
				return 0, err
			}

			vertex.Lat, vertex.Lon = getGeoPoint(
				lat, area.Attribute.GetLatitudeType() == SouthLatitudeType,
				lon, area.Attribute.GetLongitudeType() == WestLongitudeType)
			area.Vertexes = append(area.Vertexes, vertex)
		}
		entity.Items = append(entity.Items, area)
	}
	return len(data) - reader.Len(), nil
}
