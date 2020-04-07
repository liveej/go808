package protocol

// 摄像头立即拍摄命令
type T808_0x8801 struct {
	ChannelID    byte
	Cmd          uint16
	Duration     uint16
	Save         T808_0x8801_SaveFlag
	Resolution   T808_0x8801_Resolution
	Quality      byte
	Lighting     byte
	Contrast     byte
	Saturability byte
	Chroma       byte
}

// 保存标志
type T808_0x8801_SaveFlag byte

var (
	// 实时上传
	T808_0x8801_SaveFlagRemote T808_0x8801_SaveFlag = 0
	// 本地保存
	T808_0x8801_SaveFlagLocal T808_0x8801_SaveFlag = 1
)

// 分辨率
type T808_0x8801_Resolution byte

var (
	T808_0x8801_Resolution320x240  T808_0x8801_Resolution = 0x01
	T808_0x8801_Resolution640x480  T808_0x8801_Resolution = 0x02
	T808_0x8801_Resolution800x600  T808_0x8801_Resolution = 0x03
	T808_0x8801_Resolution1024x768 T808_0x8801_Resolution = 0x04
	T808_0x8801_Resolution176x144  T808_0x8801_Resolution = 0x05
	T808_0x8801_Resolution352x288  T808_0x8801_Resolution = 0x06
	T808_0x8801_Resolution704x288  T808_0x8801_Resolution = 0x07
	T808_0x8801_Resolution704x576  T808_0x8801_Resolution = 0x08
)

func (entity *T808_0x8801) MsgID() MsgID {
	return MsgT808_0x8801
}

func (entity *T808_0x8801) Encode() ([]byte, error) {
	writer := NewWriter()

	// 写入通道ID
	writer.WriteByte(entity.ChannelID)

	// 写入拍摄命令
	writer.WriteUint16(entity.Cmd)

	// 写入拍摄时间
	writer.WriteUint16(entity.Duration)

	// 写入保存标志
	writer.WriteByte(byte(entity.Save))

	// 写入分辨率
	writer.WriteByte(byte(entity.Resolution))

	// 写入图像质量
	writer.WriteByte(entity.Quality)

	// 写入亮度
	writer.WriteByte(entity.Lighting)

	// 写入对比度
	writer.WriteByte(entity.Contrast)

	// 写入饱和度
	writer.WriteByte(entity.Saturability)

	// 写入色度
	writer.WriteByte(entity.Chroma)
	return writer.Bytes(), nil
}

func (entity *T808_0x8801) Decode(data []byte) (int, error) {
	if len(data) < 12 {
		return 0, ErrInvalidBody
	}
	reader := NewReader(data)

	// 读取通道ID
	var err error
	entity.ChannelID, err = reader.ReadByte()
	if err != nil {
		return 0, err
	}

	// 读取拍摄命令
	entity.Cmd, err = reader.ReadUint16()
	if err != nil {
		return 0, err
	}

	// 读取拍摄时间
	entity.Duration, err = reader.ReadUint16()
	if err != nil {
		return 0, err
	}

	// 读取保存标志
	flag, err := reader.ReadByte()
	if err != nil {
		return 0, err
	}
	entity.Save = T808_0x8801_SaveFlag(flag)

	// 读取分辨率
	resolution, err := reader.ReadByte()
	if err != nil {
		return 0, err
	}
	entity.Resolution = T808_0x8801_Resolution(resolution)

	// 读取图像质量
	entity.Quality, err = reader.ReadByte()
	if err != nil {
		return 0, err
	}

	// 读取亮度
	entity.Lighting, err = reader.ReadByte()
	if err != nil {
		return 0, err
	}

	// 读取对比度
	entity.Contrast, err = reader.ReadByte()
	if err != nil {
		return 0, err
	}

	// 读取饱和度
	entity.Saturability, err = reader.ReadByte()
	if err != nil {
		return 0, err
	}

	// 读取色度
	entity.Chroma, err = reader.ReadByte()
	if err != nil {
		return 0, err
	}
	return len(data) - reader.Len(), nil
}
