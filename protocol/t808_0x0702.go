package protocol

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"time"
)

// 驾驶员身份信息采集上报
type T808_0x0702 struct {
	State        byte
	Time         time.Time
	ICCardResult byte
	DriverName   string
	Number       string
	CompanyName  string
	ExpiryDate   time.Time
}

func (entity *T808_0x0702) MsgID() MsgID {
	return MsgT808_0x0702
}

func (entity *T808_0x0702) Encode() ([]byte, error) {
	writer := NewWriter()

	// 写入状态
	writer.WriteByte(entity.State)

	// 写入时间
	writer.WriteBcdTime(entity.Time)

	// 写入IC卡结果
	writer.WriteByte(entity.ICCardResult)

	// 写入驾驶员姓名
	reader := bytes.NewReader([]byte(entity.DriverName))
	driverName, err := ioutil.ReadAll(
		transform.NewReader(reader, simplifiedchinese.GB18030.NewEncoder()))
	if err != nil {
		return nil, err
	}
	writer.WriteByte(byte(len(driverName)))
	writer.Write(driverName)

	// 写入从业资格编码
	reader = bytes.NewReader([]byte(entity.Number))
	number, err := ioutil.ReadAll(
		transform.NewReader(reader, simplifiedchinese.GB18030.NewEncoder()))
	if err != nil {
		return nil, err
	}
	writer.Write(number, 20)

	// 写入发行机构名称
	reader = bytes.NewReader([]byte(entity.CompanyName))
	companyName, err := ioutil.ReadAll(
		transform.NewReader(reader, simplifiedchinese.GB18030.NewEncoder()))
	if err != nil {
		return nil, err
	}
	writer.WriteByte(byte(len(companyName)))
	writer.Write(companyName)

	// 写入证件有效期
	writer.Write(stringToBCD(entity.ExpiryDate.Format("20060102"), 4))
	return writer.Bytes(), nil
}

func (entity *T808_0x0702) Decode(data []byte) (int, error) {
	if len(data) < 34 {
		return 0, ErrInvalidBody
	}
	reader := NewReader(data)

	// 读取状态
	var err error
	entity.State, err = reader.ReadByte()
	if err != nil {
		return 0, err
	}

	// 读取时间
	entity.Time, err = reader.ReadBcdTime()
	if err != nil {
		return 0, err
	}

	// 读取IC卡结果
	entity.ICCardResult, err = reader.ReadByte()
	if err != nil {
		return 0, err
	}

	// 读取驾驶员姓名
	size, err := reader.ReadByte()
	if err != nil {
		return 0, err
	}
	entity.DriverName, err = reader.ReadString(int(size))
	if err != nil {
		return 0, err
	}

	// 读取从业资格编码
	entity.Number, err = reader.ReadString(20)
	if err != nil {
		return 0, err
	}

	// 发行机构名称
	size, err = reader.ReadByte()
	if err != nil {
		return 0, err
	}
	entity.CompanyName, err = reader.ReadString(int(size))
	if err != nil {
		return 0, err
	}

	// 读取证件有效期
	bcd, err := reader.Read(4)
	if err != nil {
		return 0, err
	}
	entity.ExpiryDate, err = time.ParseInLocation(
		"20060102150405", bcdToString(bcd)+"000000", time.Local)
	if err != nil {
		return 0, err
	}
	return len(data) - reader.Len(), nil
}
