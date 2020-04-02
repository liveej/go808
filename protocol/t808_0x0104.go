package protocol

import (
	"go808/errors"
)

// 查询终端参数应答
type T808_0x0104 struct {
	ResponseMessageSerialNo uint16
	Params                  []*Param
}

func (entity *T808_0x0104) MsgID() MsgID {
	return MsgT808_0x0104
}

func (entity *T808_0x0104) Encode() ([]byte, error) {
	writer := NewWriter()

	// 写入消息序列号
	writer.WriteUint16(entity.ResponseMessageSerialNo)

	// 写入参数个数
	writer.WriteByte(byte(len(entity.Params)))

	// 写入参数列表
	for _, param := range entity.Params {
		// 写入参数ID
		writer.WriteUint32(param.id)

		// 写入参数长度
		writer.WriteByte(byte(len(param.serialized)))

		// 写入参数数据
		writer.WriteBytes(param.serialized)
	}
	return writer.Bytes(), nil
}

func (entity *T808_0x0104) Decode(data []byte) (int, error) {
	if len(data) <= 3 {
		return 0, errors.ErrEntityDecodeFail
	}
	reader := NewReader(data)

	// 读取消息序列号
	responseMessageSerialNo, err := reader.ReadUint16()
	if err != nil {
		return 0, errors.ErrEntityDecodeFail
	}

	// 读取参数个数
	paramNums, err := reader.ReadByte()
	if err != nil {
		return 0, errors.ErrEntityDecodeFail
	}

	// 读取参数信息
	params := make([]*Param, 0, paramNums)
	for i := 0; i < int(paramNums); i++ {
		// 读取参数ID
		id, err := reader.ReadUint32()
		if err != nil {
			return 0, errors.ErrEntityDecodeFail
		}

		// 读取数据长度
		size, err := reader.ReadByte()
		if err != nil {
			return 0, errors.ErrEntityDecodeFail
		}

		// 读取数据内容
		value, err := reader.ReadBytes(int(size))
		if err != nil {
			return 0, errors.ErrEntityDecodeFail
		}
		params = append(params, &Param{
			id:         id,
			serialized: value,
		})
	}

	entity.Params = params
	entity.ResponseMessageSerialNo = responseMessageSerialNo
	return len(data) - reader.Len(), nil
}
