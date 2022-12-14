package admin

import (
	"encoding/json"
	"ginshop/models"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MainController struct{}

func (con MainController) Index(c *gin.Context) {
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	userinfoStr, ok := userinfo.(string)
	if ok {
		var userinfoStruct []models.Manager
		json.Unmarshal([]byte(userinfoStr), &userinfoStruct)

		roleId := userinfoStruct[0].RoleId

		accessList := []models.Access{}
		models.DB.Where("module_id=?", 0).Preload("AccessItem", func(db *gorm.DB) *gorm.DB {
			return db.Order("access.sort DESC")
		}).Order("sort DESC").Find(&accessList)
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

		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"username":   userinfoStruct[0].Username,
			"accessList": accessList,
			"isSuper":    userinfoStruct[0].IsSuper,
		})
	} else {
		c.Redirect(302, "/admin/login")
	}
}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/welcome.html", nil)
}

func (con MainController) ChangeStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	table := c.Query("table")
	field := c.Query("field")
	err := models.DB.Exec("update "+table+" set "+field+" = 1 - "+field+" where id = ?", id).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改失败，请重试",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改数据成功",
	})
}

func (con MainController) ChangeNum(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	table := c.Query("table")
	field := c.Query("field")
	num := c.Query("num")
	err := models.DB.Exec("update "+table+" set "+field+" = "+num+" where id = ?", id).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改失败，请重试",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改数据成功",
	})
}
