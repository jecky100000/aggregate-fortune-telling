/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package admin

import (
	"gin/ay"
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

type GetControllerLoginForm struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (con Controller) Login(c *gin.Context) {
	var getForm GetControllerLoginForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, "400", ay.Validator{}.Translate(err), gin.H{})
		return
	}
}
