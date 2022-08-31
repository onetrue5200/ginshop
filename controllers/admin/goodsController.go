package admin

import (
	"ginshop/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GoodsController struct {
	BaseController
}

func (con GoodsController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goods.html", nil)
}

func (con GoodsController) Add(c *gin.Context) {
	// 获取商品分类
	goodsCateList := []models.GoodsCate{}
	models.DB.Where("pid=0").Preload("GoodsCateItems").Find(&goodsCateList)
	// 获取商品颜色信息
	goodsColorList := []models.GoodsColor{}
	models.DB.Find(&goodsColorList)
	// 获取商品规格包装
	goodsTypeList := []models.GoodsType{}
	models.DB.Find(&goodsTypeList)

	c.HTML(http.StatusOK, "admin/goodsAdd.html", gin.H{
		"goodsCateList":  goodsCateList,
		"goodsColorList": goodsColorList,
		"goodsTypeList":  goodsTypeList,
	})
}

func (con GoodsController) DoAdd(c *gin.Context) {
	attrIdList := c.PostFormArray("attr_id_list")
	attrValueList := c.PostFormArray("attr_value_list")
	goodsImageList := c.PostFormArray("goods_image_list")
	c.JSON(200, gin.H{
		"attrIdList":     attrIdList,
		"attrValueList":  attrValueList,
		"goodsImageList": goodsImageList,
	})
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

func (con GoodsController) GoodsTypeAttribute(c *gin.Context) {
	cateId, _ := strconv.Atoi(c.Query("cateId"))
	goodsTypeAttributeList := []models.GoodsTypeAttribute{}
	err := models.DB.Where("cate_id = ?", cateId).Find(&goodsTypeAttributeList).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"result":  "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"result":  goodsTypeAttributeList,
		})
	}
}
