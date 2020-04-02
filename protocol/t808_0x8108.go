package protocol

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

// 下发终端升级包
type T808_0x8108 struct {
	Type          byte
	ManufactureID string
	Version       string
	Size          uint32
	Data          []byte
}

func (entity *T808_0x8108) MsgID() MsgID {
	return MsgT808_0x8108
}

func (entity *T808_0x8108) Encode() ([]byte, error) {
	writer := NewWriter()

	// 写入升级类型
	writer.WriteByte(entity.Type)

	// 写入制造商
	writer.WriteBytes([]byte(entity.ManufactureID), 5)

	// 转换版本编码
	reader := bytes.NewReader([]byte(entity.Version))
	version, err := ioutil.ReadAll(
		transform.NewReader(reader, simplifiedchinese.GB18030.NewEncoder()))
	if err != nil {
		return nil, err
	}

	// 写入版本长度
	writer.WriteByte(byte(len(version)))

	// 写入版本信息
	writer.WriteBytes(version)

	// 写入升级包长度
	writer.WriteUint32(entity.Size)

	// 写入升级包数据
	writer.WriteBytes(entity.Data)
	return writer.Bytes(), nil
}

func (entity *T808_0x8108) Decode(data []byte) (int, error) {
	if len(data) < 11 {
		return 0, ErrInvalidBody
	}
	reader := NewReader(data)

	// 读取升级类型
	var err error
	entity.Type, err = reader.ReadByte()
	if err != nil {
		return 0, err
	}

	// 读取制造商
	manufacture, err := reader.ReadBytes(5)
	if err != nil {
		return 0, err
	}
	entity.ManufactureID = bytesToString(manufacture)

	// 读取版本长度
	versionSize, err := reader.ReadByte()
	if err != nil {
		return 0, err
	}

	// 读取版本信息
	entity.Version, err = reader.ReadString(int(versionSize))
	if err != nil {
		return 0, err
	}

	// 读取升级包长度
	entity.Size, err = reader.ReadUint32()
	if err != nil {
		return 0, err
	}

	// 读取升级包数据
	entity.Data, err = reader.ReadBytes()
	if err != nil {
		return 0, err
	}
	return len(data) - reader.Len(), nil
}
