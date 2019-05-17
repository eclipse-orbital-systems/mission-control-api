package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/eclipse-orbital-systems/mission-control-api/api/v1/echo"
	"github.com/eclipse-orbital-systems/mission-control-api/dal"
)

type Api struct {
	Echo   *echo.Api
}

func New(router *gin.RouterGroup, dalImpl *dal.Dal) *Api {
	obj := Api{}
	obj.Echo = echo.New(dalImpl)

	router.GET("/echo/:echoId", obj.Echo.Get)
	router.POST("/echo", obj.Echo.Post)

	return &obj
}
