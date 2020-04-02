package protocol

// 查询终端属性应答
type T808_0x0107 struct {
	TerminalType    uint16 // 终端类型
	ManufactureID   string // 制造商
	Model           string // 终端型号
	TerminalID      string // 终端ID
	Sim             string // SIM卡号
	HardwareVersion string // 终端硬件版本
	SoftwareVersion string // 终端固件版本号
	GNSSProperty    byte   // GNSS模块属性
	COMMProperty    byte   // 通信模块属性
}

func (entity *T808_0x0107) MsgID() MsgID {
	return MsgT808_0x0107
}

func (entity *T808_0x0107) Encode() ([]byte, error) {
	writer := NewWriter()

	// 写入终端类型
	writer.WriteUint16(entity.TerminalType)

	// 写入终端制造商
	writer.Write([]byte(entity.ManufactureID), 5)

	// 写入终端型号
	writer.Write([]byte(entity.Model), 20)

	// 写入终端ID
	writer.Write([]byte(entity.TerminalID), 7)

	// 写入终端Sim卡号
	writer.Write(stringToBCD(entity.Sim, 10))

	// 写入终端硬件版本
	hardwareVersion := []byte(entity.HardwareVersion)
	writer.WriteByte(byte(len(hardwareVersion)))
	writer.Write(hardwareVersion)

	// 写入终端固件版本号
	softwareVersion := []byte(entity.SoftwareVersion)
	writer.WriteByte(byte(len(softwareVersion)))
	writer.Write(softwareVersion)

	// 写入GNSS模块属性
	writer.WriteByte(entity.GNSSProperty)

	// 写入通信模块属性
	writer.WriteByte(entity.COMMProperty)
	return writer.Bytes(), nil
}

func (entity *T808_0x0107) Decode(data []byte) (int, error) {
	if len(data) < 48 {
		return 0, ErrInvalidBody
	}
	reader := NewReader(data)

	// 读取终端类型
	typ, err := reader.ReadUint16()
	if err != nil {
		return 0, err
	}

	// 读取终端制造商
	manufacture, err := reader.Read(5)
	if err != nil {
		return 0, err
	}

	// 读取终端型号
	model, err := reader.Read(20)
	if err != nil {
		return 0, err
	}

	// 读取终端ID
	terminalID, err := reader.Read(7)
	if err != nil {
		return 0, err
	}

	// 读取SIM卡号
	temp, err := reader.Read(10)
	if err != nil {
		return 0, err
	}
	sim := bcdToString(temp)

	// 读取终端硬件版本号长度
	size, err := reader.ReadByte()
	if err != nil {
		return 0, err
	}

	// 读取终端硬件版本号
	temp, err = reader.Read(int(size))
	if err != nil {
		return 0, err
	}
	hardwareVersion := bytesToString(temp)

	// 读取终端软件版本号长度
	size, err = reader.ReadByte()
	if err != nil {
		return 0, err
	}

	// 读取终端软件版本号
	temp, err = reader.Read(int(size))
	if err != nil {
		return 0, err
	}
	softwareVersion := bytesToString(temp[:size])

	// 读取GNSS模块属性
	gnssProperty, err := reader.ReadByte()
	if err != nil {
		return 0, err
	}

	// 读取通信模块属性
	commProperty, err := reader.ReadByte()
	if err != nil {
		return 0, err
	}

	entity.TerminalType = typ
	entity.ManufactureID = bytesToString(manufacture)
	entity.Model = bytesToString(model)
	entity.TerminalID = bytesToString(terminalID)
	entity.Sim = sim
	entity.HardwareVersion = hardwareVersion
	entity.SoftwareVersion = softwareVersion
	entity.GNSSProperty = gnssProperty
	entity.COMMProperty = commProperty
	return len(data) - reader.Len(), nil
}
