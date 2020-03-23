package protocol

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// 终端鉴权
type T808_0x0102 struct {
	AuthKey string
}

// 获取类型
func (entity *T808_0x0102) Type() Type {
	return TypeT808_0x0102
}

// 消息编码
func (entity *T808_0x0102) Encode() ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if len(entity.AuthKey) > 0 {
		authKey, err := ioutil.ReadAll(transform.NewReader(
			bytes.NewReader([]byte(entity.AuthKey)), simplifiedchinese.GB18030.NewEncoder()))
		if err != nil {
			return nil, err
		}
		buffer.Write(authKey)
	}
	return buffer.Bytes(), nil
}

// 消息解码
func (entity *T808_0x0102) Decode(data []byte) (int, error) {
	if len(data) == 0 {
		return 0, ErrEntityDecode
	}

	reader := bytes.NewReader(data)
	authKey, err := ioutil.ReadAll(transform.NewReader(reader, simplifiedchinese.GB18030.NewDecoder()))
	if err != nil {
		return 0, ErrEntityDecode
	}
	entity.AuthKey = string(authKey)
	return len(data) - reader.Len(), nil
}
