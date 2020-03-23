package protocol

import (
	"bytes"
	"encoding/binary"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

// 终端注册
type T808_0x0100 struct {
	ProvinceID    uint16 // 省份
	CityID        uint16 // 城市
	ManufactureID string // 制造商
	Model         string // 终端型号
	TerminalID    string // 终端ID
	PlateColor    byte   // 车牌颜色
	LicenseNo     string // 车辆标识
}

// 获取类型
func (entity *T808_0x0100) Type() Type {
	return TypeT808_0x0100
}

// 消息编码
func (entity *T808_0x0100) Encode() ([]byte, error) {
	var temp [2]byte
	buffer := bytes.NewBuffer(nil)

	// 写入省份ID
	binary.BigEndian.PutUint16(temp[:], entity.ProvinceID)
	buffer.Write(temp[:])

	// 写入城市ID
	binary.BigEndian.PutUint16(temp[:], entity.CityID)
	buffer.Write(temp[:])

	// 写入制造商
	var manufacturer [5]byte
	copy(manufacturer[:], entity.ManufactureID)
	buffer.Write(manufacturer[:])

	// 写入终端型号
	var model [20]byte
	copy(model[:], entity.Model)
	buffer.Write(model[:])

	// 写入终端ID
	var terminalID [7]byte
	copy(terminalID[:], entity.TerminalID)
	buffer.Write(terminalID[:])

	// 写入车牌颜色
	buffer.WriteByte(entity.PlateColor)

	// 写入车辆标识
	if len(entity.LicenseNo) > 0 {
		reader := bytes.NewReader([]byte(entity.LicenseNo))
		licenseNo, err := ioutil.ReadAll(transform.NewReader(reader, simplifiedchinese.GB18030.NewEncoder()))
		if err == nil {
			buffer.Write(licenseNo)
		}
	}
	return buffer.Bytes(), nil
}

// 消息解码
func (entity *T808_0x0100) Decode(data []byte) (int, error) {
	if len(data) < 37 {
		return 0, ErrEntityDecode
	}

	var temp [2]byte
	reader := bytes.NewReader(data)

	// 读取省份ID
	count, err := reader.Read(temp[:2])
	if err != nil || count != 2 {
		return 0, ErrEntityDecode
	}
	province := binary.BigEndian.Uint16(temp[:2])

	// 读取城市ID
	count, err = reader.Read(temp[:2])
	if err != nil || count != 2 {
		return 0, ErrEntityDecode
	}
	city := binary.BigEndian.Uint16(temp[:2])

	// 读取制造商
	var manufacturer [5]byte
	count, err = reader.Read(manufacturer[:])
	if err != nil || count != len(manufacturer) {
		return 0, ErrEntityDecode
	}

	// 读取终端型号
	var model [20]byte
	count, err = reader.Read(model[:])
	if err != nil || count != len(model) {
		return 0, ErrEntityDecode
	}

	// 读取终端ID
	var terminalID [7]byte
	count, err = reader.Read(terminalID[:])
	if err != nil || count != len(terminalID) {
		return 0, ErrEntityDecode
	}

	// 读取车牌颜色
	count, err = reader.Read(temp[:1])
	if err != nil || count != 1 {
		return 0, ErrEntityDecode
	}
	color := temp[0]

	entity.ProvinceID = province
	entity.CityID = city
	entity.ManufactureID = BytesToString(manufacturer[:])
	entity.Model = BytesToString(model[:])
	entity.TerminalID = BytesToString(terminalID[:])
	entity.PlateColor = color

	// 读取车辆标识
	if reader.Len() > 0 {
		licenseNo, err := ioutil.ReadAll(transform.NewReader(reader, simplifiedchinese.GB18030.NewDecoder()))
		if err == nil {
			entity.LicenseNo = string(licenseNo)
		}
	}
	return len(data) - reader.Len(), nil
}
