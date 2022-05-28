/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package api

import (
	"aggregate-fortune-telling/ay"
	"aggregate-fortune-telling/models/common"
	"github.com/gin-gonic/gin"
)

type CalenderController struct {
}

type GetCalenderGetForm struct {
	Y int `form:"y" binding:"required" label:"年份"`
	M int `form:"m" binding:"required" label:"月份"`
	D int `form:"d" binding:"required" label:"日"`
}

func (con CalenderController) Get(c *gin.Context) {
	var getForm GetCalenderGetForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}
	if getForm.M > 12 || getForm.M < 1 {
		ay.Json{}.Msg(c, 400, "请输入正确的月份", gin.H{})
		return
	}
	if getForm.D > 31 || getForm.D < 1 {
		ay.Json{}.Msg(c, 400, "请输入正确的天数", gin.H{})
		return
	}

	arr := common.CalendarModel{}.Get(getForm.Y, getForm.M, getForm.D)
	ay.Json{}.Msg(c, 200, "success", arr)
}
