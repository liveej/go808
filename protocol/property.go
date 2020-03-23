package protocol

// 消息体属性
type Property uint16

// 获取消息体长度
func (property *Property) GetBodySize() uint16 {
	// 前十位表示消息体长度
	// 0x3ff == ‭001111111111‬
	val := uint16(*property)
	return ((val << 6) >> 6) & 0x3ff
}

// 设置消息体长度
func (property *Property) SetBodySize(size uint16) error {
	if size > 0x3ff {
		return ErrBodyTooLong
	}
	val := uint16(*property)
	*property = Property(((val >> 10) << 10) | size)
	return nil
}
