package util

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

// int64tobyte[]
func Int64ToBytes(value int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(value))
	return buf
}

// bytes2int64
func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

// string2byte[]
func EncodeString(value string) (ret []byte) {
	return []byte(value)
}

// byte[]2string
func DecodeString(value []byte) string  {
	return string(value[:])
}

// String 转 int64
func String2Int64(value string) int64 {
	res, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		fmt.Errorf("转换错误", err)
	}
	return res
}

// String 转 int
func String2Int(value string) int {
	res, err := strconv.Atoi(value)
	if err != nil {
		fmt.Errorf("转换错误", err)
	}
	return res
}

// Float64 转 string
func Float64ToString(value float64) string {
	return strconv.FormatFloat(value, 'E', -1, 64)
}
