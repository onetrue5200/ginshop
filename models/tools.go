package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"html/template"
	"io"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 获取时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

// 获取时间戳(纳秒)
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func UploadImg(c *gin.Context, picName string) (string, error) {
	file, err := c.FormFile(picName)
	if err != nil {
		return "", err
	}
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("图片后缀不合法")
	}

	day := GetDay()
	dir := "./static/upload/" + day
	err = os.MkdirAll(dir, 0750)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fileName := strconv.FormatInt(GetUnixNano(), 10) + extName

	dst := path.Join(dir, fileName)
	c.SaveUploadedFile(file, dst)
	return dst, nil
}

// string to html
func Str2Html(str string) template.HTML {
	return template.HTML(str)
}
