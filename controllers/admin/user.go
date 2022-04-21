/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package admin

import (
	"gin/ay"
	"gin/models"
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
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var admin models.Admin
	ay.Db.First(&admin, "account = ?", getForm.Account)

	if admin.Id == 0 {
		ay.Json{}.Msg(c, 400, "账号不存在", gin.H{})
		return
	}

	if admin.Password != ay.MD5(getForm.Password) {
		//log.Println(ay.MD5(getForm.Password))
		ay.Json{}.Msg(c, 400, "密码错误", gin.H{})
		return
	}

	token := ay.AuthCode(admin.Account, "ENCODE", "", 0)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"token": token,
	})
}
