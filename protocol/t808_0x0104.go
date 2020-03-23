package protocol

// 终端参数
type Param interface {
	ID() uint32
	Value() (interface{}, error)
	Encode() (int, error)
	Decode(data []byte, size byte) (int, error)
}

// 查询终端参数应答
type T808_0x0104 struct {
	ResponseMessageSerialNo uint16
	Params                  []Param
}

// 获取类型
func (entity *T808_0x0104) Type() Type {
	return TypeT808_0x0104
}

// 消息编码
func (entity *T808_0x0104) Encode() ([]byte, error) {
	return nil, nil
}

// 消息解码
func (entity *T808_0x0104) Decode(data []byte) (int, error) {
	return 0, nil
	//if len(data) <= 3 {
	//	return 0, ErrT808_0x0104
	//}
	//
	//// 读取消息序列号
	//buffer := make([]byte, 255)
	//reader := bytes.NewReader(data)
	//count, err := reader.Read(buffer[:2])
	//if err != nil || count != 2 {
	//	return 0, ErrT808_0x0104
	//}
	//responseMessageSerialNo := binary.BigEndian.Uint16(buffer[:2])
	//
	//// 读取参数个数
	//count, err = reader.Read(buffer[:1])
	//if err != nil || count != 1 {
	//	return 0, ErrT808_0x0104
	//}
	//paramCount := int(buffer[0])
	//
	//// 读取参数信息
	//params := make([]Param, 0, paramCount)
	//for i := 0; i < paramCount; i++ {
	//	// 读取参数ID
	//	count, err = reader.Read(buffer[:4])
	//	if err != nil || count != 4 {
	//		return 0, ErrT808_0x0104
	//	}
	//	id := binary.BigEndian.Uint32(buffer[:4])
	//
	//	// 读取数据长度
	//	count, err = reader.Read(buffer[:1])
	//	if err != nil || count != 1 {
	//		return 0, ErrT808_0x0104
	//	}
	//	size := buffer[0]
	//
	//	// 读取数据内容
	//	count, err = reader.Read(buffer[:int(size)])
	//	if err != nil || count != int(size) {
	//		return 0, ErrT808_0x0104
	//	}
	//
	//	value := make([]byte, int(size))
	//	copy(value, buffer[:int(size)])
	//
	//	params = append(params, &parameter.Parameter{
	//		ID:    id,
	//		Value: value,
	//	})
	//}
	//
	//// 更新Entity信息
	////entity.Parameters = parameters
	//entity.ResponseMessageSerialNo = responseMessageSerialNo
	//return len(data) - reader.Len(), nil
}
