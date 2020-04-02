package protocol

import (
	"go808/errors"
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

func (entity *T808_0x8100) MsgID() MsgID {
	return MsgT808_0x8100
}

func (entity *T808_0x8100) Encode() ([]byte, error) {
	writer := NewWriter()

	// 写入消息流水号
	writer.WriteUint16(entity.MessageSerialNo)

	// 写入响应结果
	writer.WriteByte(byte(entity.Result))

	// 写入鉴权码
	if len(entity.AuthKey) > 0 {
		if err := writer.WritString(entity.AuthKey); err != nil {
			return nil, err
		}
	}
	return writer.Bytes(), nil
}

func (entity *T808_0x8100) Decode(data []byte) (int, error) {
	if len(data) < 3 {
		return 0, errors.ErrEntityDecodeFail
	}
	reader := NewReader(data)

	// 读取流水号
	messageSerialNo, err := reader.ReadUint16()
	if err != nil {
		return 0, errors.ErrEntityDecodeFail
	}

	// 读取响应结果
	temp, err := reader.ReadByte()
	if err != nil {
		return 0, errors.ErrEntityDecodeFail
	}

	// 读取鉴权码
	if reader.Len() > 0 {
		entity.AuthKey, err = reader.ReadString()
		if err != nil {
			return 0, errors.ErrEntityDecodeFail
		}
	}

	entity.Result = T808_0x8100_Result(temp)
	entity.MessageSerialNo = messageSerialNo
	return len(data) - reader.Len(), nil
}
