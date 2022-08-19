package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainController struct{}

func (con MainController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", nil)
}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/welcome.html", nil)
}
