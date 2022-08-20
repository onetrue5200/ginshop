package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (ctl BaseController) success(c *gin.Context, message string, redirectURL string) {
	c.HTML(http.StatusOK, "admin/success.html", gin.H{
		"message":     message,
		"redirectURL": redirectURL,
	})
}

func (ctl BaseController) error(c *gin.Context, message string, redirectURL string) {
	c.HTML(http.StatusOK, "admin/error.html", gin.H{
		"message":     message,
		"redirectURL": redirectURL,
	})
}
