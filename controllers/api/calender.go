/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package api

import (
	"gin/ay"
	"gin/models/common"
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
		Json.Msg(400, ay.Validator{}.Translate(err), gin.H{})
		return
	}
	if getForm.M > 12 || getForm.M < 1 {
		Json.Msg(400, "请输入正确的月份", gin.H{})
		return
	}
	if getForm.D > 31 || getForm.D < 1 {
		Json.Msg(400, "请输入正确的天数", gin.H{})
		return
	}

	arr := common.CalendarModel{}.Get(getForm.Y, getForm.M, getForm.D)
	Json.Msg(200, "success", arr)
}
