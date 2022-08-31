package admin

import (
	"ginshop/models"
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

func (con GoodsController) ImageUpload(c *gin.Context) {
	imgDir, err := models.UploadImg(c, "file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"link": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"link": "/" + imgDir,
		})
	}
}
