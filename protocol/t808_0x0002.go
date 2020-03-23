package protocol

// 终端心跳
type T808_0x0002 struct {
}

// 获取类型
func (entity *T808_0x0002) Type() Type {
	return TypeT808_0x0002
}

// 消息编码
func (entity *T808_0x0002) Encode() ([]byte, error) {
	return nil, nil
}

// 消息解码
func (entity *T808_0x0002) Decode(data []byte) (int, error) {
	return 0, nil
}
