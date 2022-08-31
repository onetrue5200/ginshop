package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GoodsController struct {
	BaseController
}

func (con GoodsController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goods.html", nil)
}

func (con GoodsController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goodsAdd.html", nil)
}
