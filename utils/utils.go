package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

// 将int64 转化为 []byte 切片
func IntToHex(num int64) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buf.Bytes()
}
