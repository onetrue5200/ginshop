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

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
