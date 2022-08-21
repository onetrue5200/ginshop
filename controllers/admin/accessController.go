package admin

import (
	"ginshop/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type AccessController struct {
	BaseController
}

func (ctl AccessController) Index(c *gin.Context) {
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)
	c.HTML(http.StatusOK, "admin/access.html", gin.H{
		"accessList": accessList,
	})
}
func (ctl AccessController) Add(c *gin.Context) {
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Find(&accessList)
	c.HTML(http.StatusOK, "admin/access/add.html", gin.H{
		"accessList": accessList,
	})
}
func (ctl AccessController) DoAdd(c *gin.Context) {
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	accessType, _ := strconv.Atoi(c.PostForm("type"))
	actionName := c.PostForm("action_name")
	url := c.PostForm("url")
	moduleId, _ := strconv.Atoi(c.PostForm("module_id"))
	sort, _ := strconv.Atoi(c.PostForm("sort"))
	status, _ := strconv.Atoi(c.PostForm("status"))
	description := c.PostForm("description")

	if moduleName == "" {
		ctl.Error(c, "模块名不能为空", "/admin/access/add")
		return
	}

	access := models.Access{
		ModuleName:  moduleName,
		Type:        accessType,
		ActionName:  actionName,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Description: description,
		Status:      status,
	}

	err := models.DB.Create(&access).Error
	if err != nil {
		ctl.Error(c, "增加数据失败", "admin/access/add")
	} else {
		ctl.Success(c, "增加数据成功", "admin/access")
	}
}

func (ctl AccessController) Edit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	access := models.Access{Id: id}
	models.DB.Find(&access)
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)
	c.HTML(http.StatusOK, "admin/access/edit.html", gin.H{
		"access":     access,
		"accessList": accessList,
	})
}
func (ctl AccessController) DoEdit(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	accessType, _ := strconv.Atoi(c.PostForm("type"))
	actionName := c.PostForm("action_name")
	url := c.PostForm("url")
	moduleId, _ := strconv.Atoi(c.PostForm("module_id"))
	sort, _ := strconv.Atoi(c.PostForm("sort"))
	status, _ := strconv.Atoi(c.PostForm("status"))
	description := c.PostForm("description")

	access := models.Access{Id: id}
	models.DB.Find(&access)
	access.ModuleName = moduleName
	access.Type = accessType
	access.ActionName = actionName
	access.Url = url
	access.ModuleId = moduleId
	access.Sort = sort
	access.Status = status
	access.Description = description

	err := models.DB.Save(&access).Error
	if err != nil {
		ctl.Error(c, "修改数据失败", "/admin/access/edit?id="+c.PostForm("id"))
	} else {
		ctl.Success(c, "修改数据成功", "/admin/access")
	}
}
func (ctl AccessController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	access := models.Access{Id: id}
	if access.ModuleId == 0 {
		accessList := []models.Access{}
		models.DB.Where("module_id=?", access.Id).Find(&accessList)
		if len(accessList) > 0 {
			ctl.Error(c, "当前模块下有菜单或操作，不可删除", "/admin/access")
			return
		}
	}
	models.DB.Delete(&access)
	ctl.Success(c, "删除数据成功", "/admin/access")
}
