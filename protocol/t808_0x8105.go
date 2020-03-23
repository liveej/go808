package protocol

import (
	"bytes"
)

// 终端控制
type T808_0x8105 struct {
	Cmd    T808_0x8105Command
	Data string
}

// 终端控制命令
type T808_0x8105Command byte

var (
	// 升级
	T808_0x8105CommandUpgrade T808_0x8105Command = 1
	// 设置服务器
	T808_0x8105CommandHost T808_0x8105Command = 2
	// 关机
	T808_0x8105CommandShutdown T808_0x8105Command = 3
	// 重置
	T808_0x8105CommandReset T808_0x8105Command = 4
	// 恢复出厂设置
	T808_0x8105CommandFactoryReset T808_0x8105Command = 5
	// 关闭无线网络
	T808_0x8105CommandCloseNetwork T808_0x8105Command = 6
	// 关闭所有网络
	T808_0x8105CommandCloseAllNetwork T808_0x8105Command = 7
)

// 获取类型
func (entity *T808_0x8105) Type() Type {
	return TypeT808_0x8105
}

// 消息编码
func (entity *T808_0x8105) Encode() ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteByte(byte(entity.Cmd))
	if len(entity.Data) > 0 {
		buffer.WriteString(entity.Data)
	}
	return buffer.Bytes(), nil
}

// 消息解码
func (entity *T808_0x8105) Decode(data []byte) (int, error) {
	return 0, nil
}
