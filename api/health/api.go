package health

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/palantir/stacktrace"
	"log"
	"net/http"
)

type Api struct {
	mysql  *gorm.DB
	router *gin.RouterGroup
}

func New(router *gin.RouterGroup, mysql *gorm.DB) *Api {
	obj := Api{}
	obj.mysql = mysql
	obj.router = router

	obj.router.GET("/liveness", obj.getLiveness)
	obj.router.GET("/readiness", obj.getReadiness)

	return &obj
}

func (obj *Api) getLiveness(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (obj *Api) getReadiness(c *gin.Context) {
	if err := obj.mysql.DB().Ping(); err != nil {
		log.Print(stacktrace.Propagate(err, "failed to ping mysql database"))
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
