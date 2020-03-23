package protocol

import (
	"bytes"
	"encoding/hex"
	"errors"
	"time"
)

// 字节置零
func ZeroBytes(data []byte) []byte {
	for i := range data {
		data[i] = 0x00
	}
	return data
}

// 二进制转字符串
func BytesToString(data []byte) string {
	n := bytes.IndexByte(data, 0)
	if n == -1 {
		return string(data)
	}
	return string(data[:n])
}

// ToBCDTime 转为BCD时间
func ToBCDTime(t time.Time) ([]byte, error) {
	t = time.Unix(t.Unix(), 0)
	s := t.Format("20060102150405")[2:]
	return hex.DecodeString(s)
}

// FromBCDTime 转为time.Time
func FromBCDTime(bcd []byte) (time.Time, error) {
	if len(bcd) != 6 {
		return time.Time{}, errors.New("format mismatch")
	}
	t, err := time.ParseInLocation(
		"20060102150405", "20"+hex.EncodeToString(bcd), time.Local)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
