package middlewares

import (
	"encoding/json"
	"fmt"
	"ginshop/models"
	"os"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

func InitAdminAuthMiddleware(c *gin.Context) {
	pathname := strings.Split(c.Request.URL.String(), "?")[0]
	session := sessions.Default(c)
	userinfoJson := session.Get("userinfo")
	userinfoStr, ok := userinfoJson.(string)
	if ok {
		var userinfo []models.Manager
		json.Unmarshal([]byte(userinfoStr), &userinfo)
		if len(userinfo) > 0 && userinfo[0].Username != "" {
			urlPath := strings.Replace(pathname, "/admin", "", 1)
			if userinfo[0].IsSuper == 1 || excludeAuthPath(urlPath) {
				return
			}
			roleId := userinfo[0].RoleId
			roleAccess := []models.RoleAccess{}
			models.DB.Where("role_id=?", roleId).Find(&roleAccess)
			roleAccessMap := make(map[int]int, len(roleAccess))
			for _, v := range roleAccess {
				roleAccessMap[v.AccessId] = roleId
			}
			access := models.Access{}
			models.DB.Where("url=?", urlPath).Find(&access)

			if _, ok := roleAccessMap[access.Id]; !ok {
				c.String(200, "没有权限")
				c.Abort()
			}

			for _, v := range roleAccess {
				roleAccessMap[v.AccessId] = roleId
			}
			return
		}
	}
	if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
		c.Redirect(302, "/admin/login")
	}
}

func excludeAuthPath(urlPath string) bool {
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	excludeAuthPath := config.Section("").Key("excludeAuthPath").String()
	excludeAuthPathSlice := strings.Split(excludeAuthPath, ",")
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}
