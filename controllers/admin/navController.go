package admin

import (
	"fmt"
	"ginshop/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NavController struct {
	BaseController
}

func (con NavController) Index(c *gin.Context) {
	//当前页
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	fmt.Println(page)
	//每页显示的数量
	pageSize := 8
	//获取数据
	navList := []models.Nav{}
	models.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&navList)

	//获取总数量
	var count int64
	models.DB.Table("nav").Count(&count)
	c.HTML(http.StatusOK, "admin/navIndex.html", gin.H{
		"navList": navList,
		//注意float64类型
		"totalPages": math.Ceil(float64(count) / float64(pageSize)),
		"page":       page,
	})
}
func (con NavController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/navAdd.html", gin.H{})
}
func (con NavController) DoAdd(c *gin.Context) {
	title := c.PostForm("title")
	link := c.PostForm("link")
	position, _ := strconv.Atoi(c.PostForm("position"))
	isOpennew, _ := strconv.Atoi(c.PostForm("is_opennew"))
	relation := c.PostForm("relation")
	sort, _ := strconv.Atoi(c.PostForm("sort"))
	status, _ := strconv.Atoi(c.PostForm("status"))
	if title == "" {
		con.Error(c, "标题不能为空", "/admin/nav/add")
		return
	}

	nav := models.Nav{
		Title:     title,
		Link:      link,
		Position:  position,
		IsOpennew: isOpennew,
		Relation:  relation,
		Sort:      sort,
		Status:    status,
		AddTime:   int(models.GetUnix()),
	}
	err := models.DB.Create(&nav).Error
	if err != nil {
		con.Error(c, "增加导航失败 请重试", "/admin/nav/add")
	} else {
		con.Success(c, "增加导航成功", "/admin/nav")
	}

}
func (con NavController) Edit(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/nav")
	} else {
		nav := models.Nav{Id: id}
		models.DB.Find(&nav)
		fmt.Println(nav)
		c.HTML(http.StatusOK, "admin/navEdit.html", gin.H{
			"nav": nav,
		})
	}

}
func (con NavController) DoEdit(c *gin.Context) {

	id, err1 := strconv.Atoi(c.PostForm("id"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/nav")
		return
	}

	title := c.PostForm("title")
	link := c.PostForm("link")
	position, _ := strconv.Atoi(c.PostForm("position"))
	isOpennew, _ := strconv.Atoi(c.PostForm("is_opennew"))
	relation := c.PostForm("relation")
	sort, _ := strconv.Atoi(c.PostForm("sort"))
	status, _ := strconv.Atoi(c.PostForm("status"))
	if title == "" {
		con.Error(c, "标题不能为空", "/admin/nav/add")
		return
	}

	nav := models.Nav{Id: id}
	models.DB.Find(&nav)
	nav.Title = title
	nav.Link = link
	nav.Position = position
	nav.IsOpennew = isOpennew
	nav.Relation = relation
	nav.Sort = sort
	nav.Status = status
	err2 := models.DB.Save(&nav).Error
	if err2 != nil {
		con.Error(c, "修改数据失败", "/admin/nav/edit?id="+c.PostForm("id"))
	} else {
		con.Success(c, "修改数据成功", "/admin/nav")
	}

}
func (con NavController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/nav")
	} else {
		nav := models.Nav{Id: id}
		models.DB.Delete(&nav)
		con.Success(c, "删除数据成功", "/admin/nav")
	}
}
