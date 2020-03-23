package protocol

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// 文本信息下发
type T808_0x8300 struct {
	Flag byte
	Text string
}

// 获取类型
func (entity *T808_0x8300) Type() Type {
	return TypeT808_0x8300
}

// 消息编码
func (entity *T808_0x8300) Encode() ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteByte(entity.Flag)
	if len(entity.Text) > 0 {
		text, err := ioutil.ReadAll(transform.NewReader(
			bytes.NewReader([]byte(entity.Text)), simplifiedchinese.GB18030.NewEncoder()))
		if err == nil {
			buffer.Write(text)
		}
	}
	return buffer.Bytes(), nil
}

// 消息解码
func (entity *T808_0x8300) Decode(data []byte) (int, error) {
	return 0, nil
}
