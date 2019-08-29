package util

import (
	"crypto/md5"
	"fmt"
)

func Md5(data []byte) (result string) {
	md5Sum := md5.Sum(data)
	result = fmt.Sprintf("%x", md5Sum) //把数组格式化成16进制，因为数组中有二进制的数据，不可读
	return
}
