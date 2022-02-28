package service

import (
	"github.com/gin-gonic/gin"
)

func Set(r *gin.Engine) *gin.Engine {
	r.Use(Cors())
	r = SetSession(r)
	r = GetHeader(r)
	return r
}
