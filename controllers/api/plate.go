/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package api

import (
	"aggregate-fortune-telling/ay"
	"aggregate-fortune-telling/models"
	"aggregate-fortune-telling/models/common"
	"fmt"
	"github.com/6tail/lunar-go/calendar"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type PlateController struct {
}

type GetPlateSubmitForm struct {
	UserName string `form:"username" binding:"required" label:"名称"`
	Gender   int    `form:"gender" binding:"required" label:"性别"`
	Y        int    `form:"y" binding:"required" label:"年份"`
	M        int    `form:"m" binding:"required" label:"月份"`
	D        int    `form:"d" binding:"required" label:"日"`
	H        int    `form:"h"`
	I        int    `form:"i"`
	AreaId   int    `form:"area_id"`
}

// Submit 生成订单
func (con PlateController) Submit(c *gin.Context) {
	var getForm GetPlateSubmitForm
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

	if getForm.H < -1 || getForm.H > 24 {
		ay.Json{}.Msg(c, 400, "请输入正确的时间", gin.H{})
		return
	}

	if getForm.I < -1 || getForm.I > 60 {
		ay.Json{}.Msg(c, 400, "请输入正确的分钟", gin.H{})
		return
	}

	oid := ay.MakeOrder(time.Now())

	uid, _ := strconv.Atoi(ay.AuthCode(Token, "DECODE", "", 0))

	if uid == 0 {
		ay.Json{}.Msg(c, 401, "token错误", gin.H{})
		return
	}

	des := getForm.UserName + "的排盘"

	order := &models.Order{
		Oid:        oid,
		Type:       2,
		Ip:         GetRequestIP(c),
		Des:        des,
		Amount:     0,
		Uid:        int64(uid),
		Status:     0,
		UserName:   getForm.UserName,
		Gender:     getForm.Gender,
		Appid:      Appid,
		PayType:    0,
		OutTradeNo: oid,
		Y:          getForm.Y,
		M:          getForm.M,
		D:          getForm.D,
		H:          getForm.H,
		I:          getForm.I,
		AreaId:     getForm.AreaId,
		Op:         2,
		OldAmount:  0,
	}

	ay.Db.Create(order)

	if order.Id == 0 {
		ay.Json{}.Msg(c, 400, "数据错误，请联系管理员", gin.H{})
	} else {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"oid": oid,
		})
	}
}

// Detail 排盘详情
func (con PlateController) Detail(c *gin.Context) {
	var getForm GetHaulDetailForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var order models.Order

	ay.Db.First(&order, "oid = ? and type = 2", getForm.Oid)

	if order.Id == 0 {
		ay.Json{}.Msg(c, 400, "订单不存在", gin.H{})
		return
	}

	t := DateFormat(order.Y, order.M, order.D, order.H, order.I, 0)
	tai := con.GetTai(t, order.AreaId)

	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", tai, loc)
	y := theTime.Year()
	m := theTime.Month()
	d := theTime.Day()
	h := theTime.Hour()
	i := theTime.Minute()

	rs := common.PlateModel{}.Detail(y, int(m), d, h, i, 0, order.Gender)

	rs["info"] = map[string]interface{}{
		"y":        order.Y,
		"m":        order.M,
		"d":        order.D,
		"h":        order.H,
		"i":        order.I,
		"gender":   order.Gender,
		"username": order.UserName,
	}
	ay.Json{}.Msg(c, 200, "success", rs)
}

type GetPlateInfoForm struct {
	Y      int `form:"y" binding:"required" label:"年份"`
	M      int `form:"m" binding:"required" label:"月份"`
	D      int `form:"d" binding:"required" label:"日"`
	H      int `form:"h"`
	I      int `form:"i"`
	AreaId int `form:"area_id"`
}

// Info 排盘首页
func (con PlateController) Info(c *gin.Context) {
	var getForm GetPlateInfoForm
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

	if getForm.H < -1 || getForm.H > 24 {
		ay.Json{}.Msg(c, 400, "请输入正确的时间", gin.H{})
		return
	}

	if getForm.I < -1 || getForm.I > 60 {
		ay.Json{}.Msg(c, 400, "请输入正确的分钟", gin.H{})
		return
	}

	t := DateFormat(getForm.Y, getForm.M, getForm.D, getForm.H, getForm.I, 0)
	tai := con.GetTai(t, getForm.AreaId)

	solar := calendar.NewSolar(getForm.Y, getForm.M, getForm.D, getForm.H, getForm.I, 0)
	lunar := solar.GetLunar()
	BaZi := lunar.GetEightChar()
	BaZi.SetSect(1)

	// 获取八字
	bz := common.HaulModel{}.GetBZ(BaZi)

	hour := ""
	vh := getForm.H

	if vh == -1 {
		hour = "未知"
	} else {
		hour = common.LunarModel{}.HourChinese(vh)
	}

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"bz":    bz,
		"tai":   tai,
		"date":  t,
		"lunar": fmt.Sprintf("%s %s时", lunar, hour),
	})

}

type GetPlateYearForm struct {
	Oid string `form:"oid" binding:"required" label:"订单号"`
	Key int    `form:"key"`
}

// Year 获取流年
func (con PlateController) Year(c *gin.Context) {
	var getForm GetPlateYearForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var order models.Order

	ay.Db.First(&order, "oid = ? and type = 2", getForm.Oid)

	if order.Id == 0 {
		ay.Json{}.Msg(c, 400, "订单不存在", gin.H{})
		return
	}

	t := DateFormat(order.Y, order.M, order.D, order.H, order.I, 0)
	tai := con.GetTai(t, order.AreaId)

	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", tai, loc)
	y := theTime.Year()
	m := theTime.Month()
	d := theTime.Day()
	h := theTime.Hour()
	i := theTime.Minute()

	rs := common.PlateModel{}.GetLiuNian(y, int(m), d, h, i, 0, order.Gender, getForm.Key)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list": rs,
	})

}

type GetPlateMonthForm struct {
	Oid   string `form:"oid" binding:"required" label:"订单号"`
	Key   int    `form:"key"`
	Index int    `form:"index"`
}

// Month 获取流月
func (con PlateController) Month(c *gin.Context) {
	var getForm GetPlateMonthForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var order models.Order

	ay.Db.First(&order, "oid = ? and type = 2", getForm.Oid)

	if order.Id == 0 {
		ay.Json{}.Msg(c, 400, "订单不存在", gin.H{})
		return
	}

	t := DateFormat(order.Y, order.M, order.D, order.H, order.I, 0)
	tai := con.GetTai(t, order.AreaId)

	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", tai, loc)
	y := theTime.Year()
	m := theTime.Month()
	d := theTime.Day()
	h := theTime.Hour()
	i := theTime.Minute()

	rs := common.PlateModel{}.GetLiuYue(y, int(m), d, h, i, 0, order.Gender, getForm.Key, getForm.Index)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list": rs,
	})

}

// GetTai  获取太阳时
func (con PlateController) GetTai(t string, id int) string {
	var area models.Area
	ay.Db.First(&area, "id = ?", id)

	if area.Id == 0 {
		return ""
	}

	pinTime := common.TaiModel{}.Pin(area.Lng, t)
	return pinTime

}
