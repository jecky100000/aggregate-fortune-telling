package ay

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Json struct {
	Serve *gin.Context
}

func (class Json) Msg(c *gin.Context, code int, msg string, data map[string]interface{}) {
	res := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	c.JSON(http.StatusOK, res)
	c.Abort()

}
