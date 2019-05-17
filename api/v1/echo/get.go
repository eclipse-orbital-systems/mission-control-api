package echo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

type getByIdRequest struct {
	Id string `binding:"required,ulid"`
}

func (obj *Api) Get(c *gin.Context) {
	req, _ := httputil.DumpRequest(c.Request, false)
	c.String(http.StatusOK, string(req))
}
