/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package service

import (
	"gin/controllers/api"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, appid := 0, 0
		for name, values := range c.Request.Header {
			// Loop over all values for the name.
			for _, value := range values {
				if name == "Authorization" {
					api.Token = value
					token = 1
				}
				if name == "From" {
					api.Appid, _ = strconv.Atoi(value)
					appid = 1
				}
			}
		}
		if token == 0 {
			api.Token = ""
		}
		if appid == 0 {
			api.Appid = 0
		}

		c.Next()
	}
}
