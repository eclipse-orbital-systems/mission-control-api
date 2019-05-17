package echo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func (obj *Api) Post(c *gin.Context) {
	req, _ := httputil.DumpRequest(c.Request, false)
	c.String(http.StatusOK, string(req))
}
