package ay

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Json struct {
	Serve *gin.Context
}

func (class Json) Msg(code int, msg string, data map[string]interface{}) {
	res := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	class.Serve.JSON(http.StatusOK, res)
	class.Serve.Abort()
}

func (class Json) Msg1(c *gin.Context, code string, msg string, data map[string]interface{}) {
	res := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	c.JSON(http.StatusOK, res)
	c.Abort()

}
