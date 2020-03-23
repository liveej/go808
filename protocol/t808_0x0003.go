package protocol

// 终端注销
type T808_0x0003 struct {
}

// 获取类型
func (entity *T808_0x0003) Type() Type {
	return TypeT808_0x0003
}

// 消息编码
func (entity *T808_0x0003) Encode() ([]byte, error) {
	return nil, nil
}

// 消息解码
func (entity *T808_0x0003) Decode(data []byte) (int, error) {
	return 0, nil
}
