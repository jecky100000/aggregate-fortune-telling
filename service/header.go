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
		token, appid := "", 0
		for name, values := range c.Request.Header {
			// Loop over all values for the name.
			for _, value := range values {
				if name == "Authorization" {
					token = value
				}
				if name == "From" {
					appid, _ = strconv.Atoi(value)
				}
			}
		}
		api.Token = token

		api.Appid = int64(appid)

		c.Next()
	}
}
