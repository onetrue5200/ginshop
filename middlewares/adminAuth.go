package middlewares

import (
	"encoding/json"
	"ginshop/models"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
			return
		}
	}
	if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
		c.Redirect(302, "/admin/login")
	}
}
