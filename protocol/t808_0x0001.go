package protocol

// 终端应答
type T808_0x0001 struct {
	ResponseMessageSerialNo uint16
	ResponseMessageID       uint16
	ResponseResult          ResponseResult
}

func (entity *T808_0x0001) MsgID() MsgID {
	return MsgT808_0x0001
}

func (entity *T808_0x0001) Encode() ([]byte, error) {
	writer := NewWriter()

	// 写入消息序列号
	writer.WriteUint16(entity.ResponseMessageSerialNo)

	// 写入响应消息ID
	writer.WriteUint16(entity.ResponseMessageID)

	// 写入响应结果
	writer.WriteByte(byte(entity.ResponseResult))
	return writer.Bytes(), nil
}

func (entity *T808_0x0001) Decode(data []byte) (int, error) {
	if len(data) < 5 {
		return 0, ErrInvalidBody
	}
	reader := NewReader(data)

	// 读取消息序列号
	responseMessageSerialNo, err := reader.ReadUint16()
	if err != nil {
		return 0, err
	}

	// 读取响应消息ID
	responseMessageID, err := reader.ReadUint16()
	if err != nil {
		return 0, err
	}

	// 读取响应结果
	result, err := reader.ReadByte()
	if err != nil {
		return 0, err
	}

	entity.ResponseMessageSerialNo = responseMessageSerialNo
	entity.ResponseMessageID = responseMessageID
	entity.ResponseResult = ResponseResult(result)
	return len(data) - reader.Len(), nil
}
