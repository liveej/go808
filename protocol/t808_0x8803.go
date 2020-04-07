package protocol

import "time"

// 存储多媒体数据上传命令
type T808_0x8803 struct {
	Type       T808_0x0801_MediaType
	ChannelID  byte
	Event      byte
	StartTime  time.Time
	EndTime    time.Time
	RemoveFlag byte
}

func (entity *T808_0x8803) MsgID() MsgID {
	return MsgT808_0x8803
}

func (entity *T808_0x8803) Encode() ([]byte, error) {
	writer := NewWriter()

	// 写入媒体类型
	writer.WriteByte(byte(entity.Type))

	// 写入通道ID
	writer.WriteByte(entity.ChannelID)

	// 写入事件项编码
	writer.WriteByte(entity.Event)

	// 写入开始时间
	writer.WriteBcdTime(entity.StartTime)

	// 写入结束时间
	writer.WriteBcdTime(entity.EndTime)

	// 写入删除标志
	writer.WriteByte(entity.RemoveFlag)
	return writer.Bytes(), nil
}

func (entity *T808_0x8803) Decode(data []byte) (int, error) {
	if len(data) < 16 {
		return 0, ErrInvalidBody
	}
	reader := NewReader(data)

	// 读取媒体类型
	mediaType, err := reader.ReadByte()
	if err != nil {
		return 0, err
	}
	entity.Type = T808_0x0801_MediaType(mediaType)

	// 读取通道ID
	entity.ChannelID, err = reader.ReadByte()
	if err != nil {
		return 0, err
	}

	// 读取事件项编码
	entity.Event, err = reader.ReadByte()
	if err != nil {
		return 0, err
	}

	// 读取开始时间
	entity.StartTime, err = reader.ReadBcdTime()
	if err != nil {
		return 0, err
	}

	// 读取结束时间
	entity.EndTime, err = reader.ReadBcdTime()
	if err != nil {
		return 0, err
	}

	// 读取删除标志
	entity.RemoveFlag, err = reader.ReadByte()
	if err != nil {
		return 0, err
	}
	return len(data) - reader.Len(), nil
}
