package protocol

// 查询车辆位置
type T808_0x8201 struct {
}

// 获取类型
func (entity *T808_0x8201) Type() Type {
	return TypeT808_0x8201
}

// 消息编码
func (entity *T808_0x8201) Encode() ([]byte, error) {
	return nil, nil
}

// 消息解码
func (entity *T808_0x8201) Decode(data []byte) (int, error) {
	return 0, nil
}
