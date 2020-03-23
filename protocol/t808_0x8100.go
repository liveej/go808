package protocol

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// 终端应答结果
type T808_0x8100_Result byte

const (
	// 成功
	T808_0x8100_ResultSuccess = 0
	// 车辆已被注册
	T808_0x8100_ResultCarRegistered = 1
	// 数据库中无该车辆
	T808_0x8100_ResultCarNotFound = 2
	// 终端已被注册
	T808_0x8100_ResultTerminalRegistered = 3
	// 数据库中无该终端
	T808_0x8100_ResultTerminalNotFound = 4
)

// 终端应答
type T808_0x8100 struct {
	MessageSerialNo uint16             // 对应的终端注册消息的流水号
	Result          T808_0x8100_Result // 结果
	AuthKey         string             // 鉴权码
}

// 获取类型
func (entity *T808_0x8100) Type() Type {
	return TypeT808_0x8100
}

// 消息编码
func (entity *T808_0x8100) Encode() ([]byte, error) {
	var tmp [2]byte
	buffer := bytes.NewBuffer(nil)

	// 写入消息流水号
	binary.BigEndian.PutUint16(tmp[:2], entity.MessageSerialNo)
	buffer.Write(tmp[:2])

	// 写入响应结果
	buffer.WriteByte(byte(entity.Result))

	// 写入鉴权码
	if len(entity.AuthKey) > 0 {
		authKey, err := ioutil.ReadAll(transform.NewReader(
			bytes.NewReader([]byte(entity.AuthKey)), simplifiedchinese.GB18030.NewEncoder()))
		if err != nil {
			return nil, err
		}
		buffer.Write(authKey)
	}
	return buffer.Bytes(), nil
}

// 消息解码
func (entity *T808_0x8100) Decode(data []byte) (int, error) {
	if len(data) < 3 {
		return 0, ErrEntityDecode
	}

	var temp [2]byte
	reader := bytes.NewReader(data)

	// 读取消息流水号
	count, err := reader.Read(temp[:2])
	if err != nil || count != 2 {
		return 0, ErrEntityDecode
	}
	messageSerialNo := binary.BigEndian.Uint16(temp[:2])

	// 读取响应结果
	count, err = reader.Read(temp[:1])
	if err != nil || count != 1 {
		return 0, ErrEntityDecode
	}
	entity.Result = T808_0x8100_Result(temp[0])

	// 读取鉴权码
	if reader.Len() > 0 {
		authKey, err := ioutil.ReadAll(transform.NewReader(reader, simplifiedchinese.GB18030.NewDecoder()))
		if err != nil {
			return 0, err
		}
		entity.AuthKey = string(authKey)
	}
	entity.MessageSerialNo = messageSerialNo
	return len(data) - reader.Len(), nil
}
