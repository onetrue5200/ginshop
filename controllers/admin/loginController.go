package admin

import (
	"encoding/json"
	"fmt"
	"ginshop/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (ctl LoginController) Index(c *gin.Context) {
	fmt.Println(models.Md5("123456"))
	c.HTML(http.StatusOK, "admin/login.html", nil)
}

func (ctl LoginController) Captcha(c *gin.Context) {
	id, b64s, err := models.MakeCaptcha()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}

func (ctl LoginController) DoLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	captchaId := c.PostForm("captcha-id")
	verifyCode := c.PostForm("verify-code")
	if flag := models.VerifyCaptcha(captchaId, verifyCode); flag {
		userinfo := []models.Manager{}
		password = models.Md5(password)
		models.DB.Where("username=? AND password=?", username, password).Find(&userinfo)
		if len(userinfo) > 0 {
			session := sessions.Default(c)
			userinfoSlice, _ := json.Marshal(userinfo)
			session.Set("userinfo", string(userinfoSlice))
			session.Save()
			ctl.success(c, "login successfully", "/admin")
		} else {
			ctl.error(c, "wrong username or password", "/admin/login")
		}
	} else {
		ctl.error(c, "verify code unmatch", "/admin/login")
	}
}

func (ctl LoginController) DoLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userinfo")
	session.Save()
	ctl.success(c, "logout", "/admin/login")
}
