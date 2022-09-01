package admin

import (
	"ginshop/models"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

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
	// 1. 获取表单提交过来的数据进行判断
	title := c.PostForm("title")
	subTitle := c.PostForm("sub_title")
	goodsSn := c.PostForm("goods_sn")
	cateId, _ := strconv.Atoi(c.PostForm("cate_id"))
	goodsNumber, _ := strconv.Atoi(c.PostForm("goods_number"))
	//注意小数点
	marketPrice, _ := strconv.ParseFloat(c.PostForm("market_price"), 64)
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)

	relationGoods := c.PostForm("relation_goods")
	goodsAttr := c.PostForm("goods_attr")
	goodsVersion := c.PostForm("goods_version")
	goodsGift := c.PostForm("goods_gift")
	goodsFitting := c.PostForm("goods_fitting")
	//获取的是切片
	goodsColorArr := c.PostFormArray("goods_color")

	goodsKeywords := c.PostForm("goods_keywords")
	goodsDesc := c.PostForm("goods_desc")
	goodsContent := c.PostForm("goods_content")
	isDelete, _ := strconv.Atoi(c.PostForm("is_delete"))
	isHot, _ := strconv.Atoi(c.PostForm("is_hot"))
	isBest, _ := strconv.Atoi(c.PostForm("is_best"))
	isNew, _ := strconv.Atoi(c.PostForm("is_new"))
	goodsTypeId, _ := strconv.Atoi(c.PostForm("goods_type_id"))
	sort, _ := strconv.Atoi(c.PostForm("sort"))
	status, _ := strconv.Atoi(c.PostForm("status"))
	addTime := int(models.GetUnix())
	// 将颜色切片转换为字符串
	goodsColorStr := strings.Join(goodsColorArr, ",")
	// 上传图片 生成缩略图
	goodsImg, _ := models.UploadImg(c, "goods_img")
	// 增加商品数据
	goods := models.Goods{
		Title:         title,
		SubTitle:      subTitle,
		GoodsSn:       goodsSn,
		CateId:        cateId,
		ClickCount:    100,
		GoodsNumber:   goodsNumber,
		MarketPrice:   marketPrice,
		Price:         price,
		RelationGoods: relationGoods,
		GoodsAttr:     goodsAttr,
		GoodsVersion:  goodsVersion,
		GoodsGift:     goodsGift,
		GoodsFitting:  goodsFitting,
		GoodsKeywords: goodsKeywords,
		GoodsDesc:     goodsDesc,
		GoodsContent:  goodsContent,
		IsDelete:      isDelete,
		IsHot:         isHot,
		IsBest:        isBest,
		IsNew:         isNew,
		GoodsTypeId:   goodsTypeId,
		Sort:          sort,
		Status:        status,
		AddTime:       addTime,
		GoodsColor:    goodsColorStr,
		GoodsImg:      goodsImg,
	}
	err := models.DB.Create(&goods).Error
	if err != nil {
		con.Error(c, "增加失败", "/admin/goods/add")
	}
	// 增加图库信息
	wg.Add(1)
	go func() {
		goodsImageList := c.PostFormArray("goods_image_list")
		for _, v := range goodsImageList {
			goodsImgObj := models.GoodsImage{}
			goodsImgObj.GoodsId = goods.Id
			goodsImgObj.ImgUrl = v
			goodsImgObj.Sort = 10
			goodsImgObj.Status = 1
			goodsImgObj.AddTime = int(models.GetUnix())
			models.DB.Create(&goodsImgObj)
		}
		wg.Done()
	}()
	//6、增加规格包装
	wg.Add(1)
	go func() {
		attrIdList := c.PostFormArray("attr_id_list")
		attrValueList := c.PostFormArray("attr_value_list")
		for i := 0; i < len(attrIdList); i++ {
			goodsTypeAttributeId, attributeIdErr := strconv.Atoi(attrIdList[i])
			if attributeIdErr == nil {
				//获取商品类型属性的数据
				goodsTypeAttributeObj := models.GoodsTypeAttribute{Id: goodsTypeAttributeId}
				models.DB.Find(&goodsTypeAttributeObj)
				//给商品属性里面增加数据  规格包装
				goodsAttrObj := models.GoodsAttr{}
				goodsAttrObj.GoodsId = goods.Id
				goodsAttrObj.AttributeTitle = goodsTypeAttributeObj.Title
				goodsAttrObj.AttributeType = goodsTypeAttributeObj.AttrType
				goodsAttrObj.AttributeId = goodsTypeAttributeObj.Id
				goodsAttrObj.AttributeCateId = goodsTypeAttributeObj.CateId
				goodsAttrObj.AttributeValue = attrValueList[i]
				goodsAttrObj.Status = 1
				goodsAttrObj.Sort = 10
				goodsAttrObj.AddTime = int(models.GetUnix())
				models.DB.Create(&goodsAttrObj)
			}

		}
		wg.Done()
	}()
	wg.Wait()
	con.Success(c, "增加数据成功", "/admin/goods")
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
