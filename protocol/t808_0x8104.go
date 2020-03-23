package protocol

// 查询终端参数
type T808_0x8104 struct {
}

// 获取类型
func (entity *T808_0x8104) Type() Type {
	return TypeT808_0x8104
}

// 消息编码
func (entity *T808_0x8104) Encode() ([]byte, error) {
	return nil, nil
}

// 消息解码
func (entity *T808_0x8104) Decode(data []byte) (int, error) {
	return 0, nil
}
