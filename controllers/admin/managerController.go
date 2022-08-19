package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ManagerController struct{}

func (con ManagerController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/manager.html", nil)
}

func (con ManagerController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/managerAdd.html", nil)
}

func (con ManagerController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/managerEdit.html", nil)
}
