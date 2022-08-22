package admin

import (
	"fmt"
	"ginshop/models"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FocusController struct {
	BaseController
}

func (con FocusController) Index(c *gin.Context) {
	focusList := []models.Focus{}
	models.DB.Find(&focusList)
	c.HTML(http.StatusOK, "admin/focus.html", gin.H{
		"focusList": focusList,
	})
}

func (con FocusController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focusAdd.html", nil)
}

func (con FocusController) DoAdd(c *gin.Context) {
	title := c.PostForm("title")
	focusType, _ := strconv.Atoi(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort, _ := strconv.Atoi(c.PostForm("sort"))
	status, _ := strconv.Atoi(c.PostForm("status"))

	focusImgSrc, err := models.UploadImg(c, "focus_img")
	if err != nil {
		fmt.Println(err)
		con.Error(c, "上传图片失败", "/admin/focus/add")
		return
	}

	focus := models.Focus{
		Title:     title,
		FocusType: focusType,
		FocusImg:  focusImgSrc,
		Link:      link,
		Sort:      sort,
		Status:    status,
		AddTime:   int(models.GetUnix()),
	}
	err = models.DB.Create(&focus).Error
	if err != nil {
		con.Error(c, "增加数据失败", "/admin/focus/add")
	} else {
		con.Success(c, "增加数据成功", "/admin/focus")
	}
}

func (con FocusController) Edit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	c.HTML(http.StatusOK, "admin/focusEdit.html", gin.H{
		"focus": focus,
	})
}

func (con FocusController) DoEdit(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	title := c.PostForm("title")
	focusType, _ := strconv.Atoi(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort, _ := strconv.Atoi(c.PostForm("sort"))
	status, _ := strconv.Atoi(c.PostForm("status"))

	focusImgSrc, _ := models.UploadImg(c, "focus_img")
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	focus.Title = title
	focus.FocusType = focusType
	focus.Link = link
	focus.Sort = sort
	focus.Status = status
	if focusImgSrc != "" {
		os.Remove(focus.FocusImg)
		focus.FocusImg = focusImgSrc
	}
	fmt.Printf("2: %#v\n", focus)
	err := models.DB.Save(&focus).Error
	if err != nil {
		con.Error(c, "修改数据失败", "/admin/focus/edit?id="+c.PostForm("id"))
	} else {
		con.Success(c, "修改数据成功", "/admin/focus")
	}
}

func (con FocusController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	os.Remove(focus.FocusImg)
	models.DB.Delete(&focus)
	con.Success(c, "删除数据成功", "/admin/role")
}
