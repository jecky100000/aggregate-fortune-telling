/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package service

import (
	"gin/ay"
	"gin/controllers/admin"
	"gin/controllers/api"
	"github.com/gin-gonic/gin"
)

func Pretreatment() gin.HandlerFunc {
	return func(c *gin.Context) {
		api.Json = &ay.Json{Serve: c}
		admin.Json = &ay.Json{Serve: c}
		c.Next()
	}
}
