package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"os"
	"path"
	"reflect"
	"strconv"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
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
	if GetOssStatus() == 1 {
		return OssUploadImg(c, picName)
	} else {
		return LocalUploadImg(c, picName)
	}
}

func FormatImg(str string) string {
	if GetOssStatus() == 1 {
		return GetSettingFromColumn("OssDomain") + str
	} else {
		return "/" + str
	}
}

func LocalUploadImg(c *gin.Context, picName string) (string, error) {
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

func OssUploadImg(c *gin.Context, picName string) (string, error) {
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
	dir := "static/upload/" + day

	fileName := strconv.FormatInt(GetUnixNano(), 10) + extName

	dst := path.Join(dir, fileName)

	OssUpload(file, dst)

	return dst, nil
}

func OssUpload(file *multipart.FileHeader, dst string) (string, error) {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New("oss-cn-shanghai.aliyuncs.com", "LTAI5tSCs8BBQHpRFM9CRvQD", "jlEwaGnwVP7hFAnZTOKT5uIeANYI0K")
	if err != nil {
		return "", err
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket("ginshop-xwz")
	if err != nil {
		return "", err
	}

	// 填写本地文件的完整路径，例如D:\\localpath\\examplefile.txt。
	// fd, err := os.Open("D:\\localpath\\examplefile.txt")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(-1)
	// }
	// defer fd.Close()

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// 将文件流上传至exampledir目录下的exampleobject.txt文件。
	err = bucket.PutObject(dst, src)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	return dst, nil
}

// 通过列获取值
func GetSettingFromColumn(columnName string) string {
	//redis file
	setting := Setting{}
	DB.First(&setting)
	//反射来获取
	v := reflect.ValueOf(setting)
	val := v.FieldByName(columnName).String()
	return val
}

func GetOssStatus() int {
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	ossStatus := config.Section("oss").Key("status").String()
	status, _ := strconv.Atoi(ossStatus)
	return status
}

// string to html
func Str2Html(str string) template.HTML {
	return template.HTML(str)
}
