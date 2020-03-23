package protocol

// 设置终端参数
type T808_0x8103 struct {
	//Parameters parameter.List
}

// 获取类型
func (entity *T808_0x8103) Type() Type {
	return TypeT808_0x8103
}

// 消息编码
func (entity *T808_0x8103) Encode() ([]byte, error) {
	//// 写入参数数量
	//var tmp [4]byte
	//buffer := bytes.NewBuffer(nil)
	//buffer.WriteByte(byte(len(entity.Parameters)))
	//
	//// 写入参数列表
	//for _, param := range entity.Parameters {
	//	// 写入参数ID
	//	binary.BigEndian.PutUint32(tmp[:], param.ID)
	//	buffer.Write(tmp[:])
	//
	//	// 写入参数长度
	//	buffer.WriteByte(byte(len(param.Value)))
	//
	//	// 写入参数数据
	//	if len(param.Value) > 0 {
	//		buffer.Write(param.Value)
	//	}
	//}
	//return buffer.Bytes(), nil
	return nil, nil
}

// 消息解码
func (entity *T808_0x8103) Decode(data []byte) (int, error) {
	return 0, ErrEntityDecode
}
