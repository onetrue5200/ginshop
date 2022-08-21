package admin

import (
	"ginshop/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	BaseController
}

func (ctl RoleController) Index(c *gin.Context) {
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/role.html", gin.H{
		"roleList": roleList,
	})
}
func (ctl RoleController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/roleAdd.html", nil)
}
func (ctl RoleController) DoAdd(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	if title == "" {
		ctl.Error(c, "角色标题不能为空", "admin/role/add")
	}
	role := models.Role{}
	role.Title = title
	role.Description = description
	role.Status = 1
	role.AddTime = int(models.GetUnix())
	err := models.DB.Create(&role).Error
	if err != nil {
		ctl.Error(c, "增加角色失败，请重试", "/admin/role/add")
	} else {
		ctl.Success(c, "增加角色成功", "/admin/role")
	}
}
func (ctl RoleController) Edit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	role := models.Role{Id: id}
	models.DB.Find(&role)
	c.HTML(http.StatusOK, "admin/roleEdit.html", gin.H{
		"role": role,
	})
}
func (ctl RoleController) DoEdit(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	if title == "" {
		ctl.Error(c, "角色标题不能为空", "admin/role/edit")
	}
	role := models.Role{Id: id}
	models.DB.Find(&role)
	role.Title = title
	role.Description = description
	err := models.DB.Save(&role).Error
	if err != nil {
		ctl.Error(c, "修改数据失败", "/admin/role?id="+c.PostForm("id"))
	} else {
		ctl.Success(c, "修改数据成功", "/admin/role")
	}
}
func (ctl RoleController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	role := models.Role{Id: id}
	models.DB.Delete(&role)
	ctl.Success(c, "删除数据成功", "/admin/role")
}
