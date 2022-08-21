package models

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"
)

// 获取时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
