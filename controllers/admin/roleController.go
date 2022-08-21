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

func (ctl RoleController) Auth(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Query("id"))
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)
	roleAccess := []models.RoleAccess{}
	models.DB.Where("role_id=?", roleId).Find(&roleAccess)
	roleAccessMap := make(map[int]int, len(roleAccess))
	for _, v := range roleAccess {
		roleAccessMap[v.AccessId] = roleId
	}
	for i := 0; i < len(accessList); i++ {
		if _, ok := roleAccessMap[accessList[i].Id]; ok {
			accessList[i].Checked = true
		}
		for j := 0; j < len(accessList[i].AccessItem); j++ {
			if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
				accessList[i].AccessItem[j].Checked = true
			}
		}
	}
	c.HTML(http.StatusOK, "admin/roleAuth.html", gin.H{
		"roleId":     roleId,
		"accessList": accessList,
	})
}

func (ctl RoleController) DoAuth(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.PostForm("roleId"))
	accessIds := c.PostFormArray("access_node[]")

	roleAccess := models.RoleAccess{}

	models.DB.Where("role_id=?", roleId).Delete(&roleAccess)

	for _, v := range accessIds {
		roleAccess.RoleId = roleId
		roleAccess.AccessId, _ = strconv.Atoi(v)
		models.DB.Create(&roleAccess)
	}

	ctl.Success(c, "修改数据成功", "/admin/role")
}
