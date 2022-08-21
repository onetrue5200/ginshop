package admin

import (
	"ginshop/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {
	managerList := []models.Manager{}
	models.DB.Preload("Role").Find(&managerList)
	c.HTML(http.StatusOK, "admin/manager.html", gin.H{
		"managerList": managerList,
	})
}

func (con ManagerController) Add(c *gin.Context) {
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/managerAdd.html", gin.H{
		"roleList": roleList,
	})
}

func (con ManagerController) DoAdd(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.PostForm("role_id"))
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	managerList := []models.Manager{}
	models.DB.Where("username=?", username).Find(&managerList)
	if len(managerList) > 0 {
		con.Error(c, "此管理员已存在", "/admin/manager/add")
		return
	}
	manager := models.Manager{
		Username: username,
		Password: models.Md5(password),
		Email:    email,
		Mobile:   mobile,
		RoleId:   roleId,
		AddTime:  int(models.GetUnix()),
		Status:   1,
	}
	err := models.DB.Create(&manager).Error
	if err != nil {
		con.Error(c, "增加管理员失败", "/admin/manager/add")
	} else {
		con.Success(c, "增加管理员成功", "/admin/manager")
	}
}

func (con ManagerController) Edit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	manager := models.Manager{Id: id}
	models.DB.Find(&manager)
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/managerEdit.html", gin.H{
		"manager":  manager,
		"roleList": roleList,
	})
}

func (con ManagerController) DoEdit(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	role_id, _ := strconv.Atoi(c.PostForm("role_id"))
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	manager := models.Manager{Id: id}
	manager.Username = username
	manager.Email = email
	manager.Mobile = mobile
	manager.RoleId = role_id
	if password != "" {
		manager.Password = models.Md5(password)
	}
	err := models.DB.Save(&manager).Error
	if err != nil {
		con.Error(c, "修改数据失败", "/admin/manager/edit?id="+c.PostForm(("id")))
	} else {
		con.Success(c, "修改数据成功", "/admin/manager")
	}
}

func (con ManagerController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	manager := models.Manager{Id: id}
	models.DB.Delete(&manager)
	con.Success(c, "删除数据成功", "/admin/role")
}
