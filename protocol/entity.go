package protocol

// 消息实体
type Entity interface {
	Type() Type
	Encode() ([]byte, error)
	Decode([]byte) (int, error)
}
