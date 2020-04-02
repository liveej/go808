package protocol

// 平台通用应答
type T808_0x8001 struct {
	ResponseMessageSerialNo uint16         // 对应的终端消息的流水号
	ResponseMsgID           MsgID          // 对应的终端消息的ID
	Result                  ResponseResult // 处理结果
}

func (entity *T808_0x8001) MsgID() MsgID {
	return MsgT808_0x8001
}

func (entity *T808_0x8001) Encode() ([]byte, error) {
	writer := NewWriter()

	// 写入流水号
	writer.WriteUint16(entity.ResponseMessageSerialNo)

	// 写入消息ID
	writer.WriteUint16(uint16(entity.ResponseMsgID))

	// 写入处理结果
	writer.WriteByte(byte(entity.Result))
	return writer.Bytes(), nil
}

func (entity *T808_0x8001) Decode(data []byte) (int, error) {
	if len(data) < 5 {
		return 0, ErrInvalidBody
	}
	reader := NewReader(data)

	// 读取流水号
	messageSerialNo, err := reader.ReadUint16()
	if err != nil {
		return 0, err
	}

	// 读取消息ID
	msgID, err := reader.ReadUint16()
	if err != nil {
		return 0, err
	}

	// 读取处理结果
	b, err := reader.ReadByte()
	if err != nil {
		return 0, err
	}

	entity.ResponseMsgID = MsgID(msgID)
	entity.Result = ResponseResult(b)
	entity.ResponseMessageSerialNo = messageSerialNo
	return len(data) - reader.Len(), nil
}
