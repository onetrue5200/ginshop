package admin

import (
	"fmt"
	"ginshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
}

func (ctl LoginController) Index(c *gin.Context) {
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
	captchaId := c.PostForm("captcha-id")
	verifyCode := c.PostForm("verify-code")
	if flag := models.VerifyCaptcha(captchaId, verifyCode); flag == true {
		c.String(http.StatusOK, "success")
	} else {
		c.String(http.StatusOK, "fail")
	}
}
