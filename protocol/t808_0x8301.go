package protocol

import (
	"bytes"
	"go808/errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

// 事件设置
type T808_0x8301 struct {
	Type   byte
	Events []T808_0x8301_Event
}

// 事件信息
type T808_0x8301_Event struct {
	EventID byte
	Content string
}

func (entity *T808_0x8301) MsgID() MsgID {
	return MsgT808_0x8301
}

func (entity *T808_0x8301) Encode() ([]byte, error) {
	writer := NewWriter()

	// 写入设置类型
	writer.WriteByte(entity.Type)

	// 写入事件数量
	writer.WriteByte(byte(len(entity.Events)))

	// 写入事件列表
	for _, event := range entity.Events {
		// 写入事件ID
		writer.WriteByte(event.EventID)

		// 写入内容长度
		reader := bytes.NewReader([]byte(event.Content))
		content, err := ioutil.ReadAll(
			transform.NewReader(reader, simplifiedchinese.GB18030.NewEncoder()))
		if err != nil {
			return nil, err
		}
		writer.WriteByte(byte(len(content)))

		// 写入事件内容
		writer.WriteBytes(content)
	}
	return writer.Bytes(), nil
}

func (entity *T808_0x8301) Decode(data []byte) (int, error) {
	if len(data) < 2 {
		return 0, errors.ErrEntityDecodeFail
	}
	reader := NewReader(data)

	// 读取设置类型
	var err error
	entity.Type, err = reader.ReadByte()
	if err != nil {
		return 0, errors.ErrEntityDecodeFail
	}

	// 读取事件数量
	count, err := reader.ReadByte()
	if err != nil {
		return 0, errors.ErrEntityDecodeFail
	}

	// 读取事件列表
	for i := 0; i < int(count); i++ {
		var event T808_0x8301_Event

		// 读取事件ID
		event.EventID, err = reader.ReadByte()
		if err != nil {
			return 0, errors.ErrEntityDecodeFail
		}

		// 读取内容长度
		size, err := reader.ReadByte()
		if err != nil {
			return 0, errors.ErrEntityDecodeFail
		}

		// 读取事件内容
		event.Content, err = reader.ReadString(int(size))
		if err != nil {
			return 0, errors.ErrEntityDecodeFail
		}
		entity.Events = append(entity.Events, event)
	}
	return len(data) - reader.Len(), nil
}
