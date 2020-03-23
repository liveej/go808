package protocol

import (
	"bytes"
	"encoding/binary"
)

// 查询终端属性应答
type T808_0x0107 struct {
	TerminalType    uint16 // 终端类型
	ManufactureID   string // 制造商
	Model           string // 终端型号
	TerminalID      string // 终端ID
	HardwareVersion string // 终端硬件版本
	SoftwareVersion string // 终端固件版本号
	GNSSProperty    byte   // GNSS模块属性
	COMMProperty    byte   // 通信模块属性
}

// 获取类型
func (entity *T808_0x0107) Type() Type {
	return TypeT808_0x0107
}

// 消息编码
func (entity *T808_0x0107) Encode() ([]byte, error) {
	var temp [256]byte
	buffer := bytes.NewBuffer(nil)

	// 写入终端类型
	binary.BigEndian.PutUint16(temp[:2], entity.TerminalType)
	buffer.Write(temp[:2])

	// 写入终端制造商
	ZeroBytes(temp[:5])
	copy(temp[:5], entity.ManufactureID)
	buffer.Write(temp[:5])

	// 写入终端型号
	ZeroBytes(temp[:20])
	copy(temp[:20], entity.Model)
	buffer.Write(temp[:20])

	// 写入终端ID
	ZeroBytes(temp[:7])
	copy(temp[:7], entity.TerminalID)
	buffer.Write(temp[:7])

	// 写入终端硬件版本
	hardwareVersion := []byte(entity.HardwareVersion)
	buffer.WriteByte(byte(len(hardwareVersion)))
	buffer.Write(hardwareVersion)

	// 写入终端固件版本号
	softwareVersion := []byte(entity.SoftwareVersion)
	buffer.WriteByte(byte(len(softwareVersion)))
	buffer.Write(softwareVersion)

	// 写入GNSS模块属性
	buffer.WriteByte(entity.GNSSProperty)

	// 写入通信模块属性
	buffer.WriteByte(entity.COMMProperty)

	return buffer.Bytes(), nil
}

// 消息解码
func (entity *T808_0x0107) Decode(data []byte) (int, error) {
	var temp [256]byte
	reader := bytes.NewReader(data)

	// 读取终端类型
	count, err := reader.Read(temp[:2])
	if err != nil || count != 2 {
		return 0, ErrEntityDecode
	}
	typ := binary.BigEndian.Uint16(temp[:2])

	// 读取终端制造商
	count, err = reader.Read(temp[:5])
	if err != nil || count != 5 {
		return 0, ErrEntityDecode
	}
	manufacture := BytesToString(temp[:5])

	// 读取终端型号
	count, err = reader.Read(temp[:20])
	if err != nil || count != 20 {
		return 0, ErrEntityDecode
	}
	model := BytesToString(temp[:20])

	// 读取终端ID
	count, err = reader.Read(temp[:7])
	if err != nil || count != 7 {
		return 0, ErrEntityDecode
	}
	terminalID := BytesToString(temp[:7])

	// 读取终端硬件版本号长度
	count, err = reader.Read(temp[:1])
	if err != nil || count != 1 {
		return 0, ErrEntityDecode
	}
	size := int(temp[0])

	// 读取终端硬件版本号
	count, err = reader.Read(temp[:size])
	if err != nil || count != size {
		return 0, ErrEntityDecode
	}
	hardwareVersion := BytesToString(temp[:size])

	// 读取终端软件版本号长度
	count, err = reader.Read(temp[:1])
	if err != nil || count != 1 {
		return 0, ErrEntityDecode
	}
	size = int(temp[0])

	// 读取终端软件版本号
	count, err = reader.Read(temp[:size])
	if err != nil || count != size {
		return 0, ErrEntityDecode
	}
	softwareVersion := BytesToString(temp[:size])

	// 读取GNSS模块属性
	count, err = reader.Read(temp[:1])
	if err != nil || count != 1 {
		return 0, ErrEntityDecode
	}
	gnssProperty := temp[0]

	// 读取通信模块属性
	count, err = reader.Read(temp[:1])
	if err != nil || count != 1 {
		return 0, ErrEntityDecode
	}
	commProperty := temp[0]

	// 更新Entity信息
	entity.TerminalType = typ
	entity.ManufactureID = manufacture
	entity.Model = model
	entity.TerminalID = terminalID
	entity.HardwareVersion = hardwareVersion
	entity.SoftwareVersion = softwareVersion
	entity.GNSSProperty = gnssProperty
	entity.COMMProperty = commProperty
	return len(data) - reader.Len(), nil
}
