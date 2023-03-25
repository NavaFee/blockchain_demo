package utils

import (
	"blockchain/conf"
	"bytes"
	"encoding/binary"
	"log"
	"os"
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

// 判断数据库是否存在
func DBExists() bool {
	if _, err := os.Stat(conf.DBNAME); os.IsNotExist(err) {
		return false
	}
	return true
}
