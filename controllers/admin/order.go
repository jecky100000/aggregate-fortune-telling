/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package admin

import (
	"aggregate-fortune-telling/ay"
	"aggregate-fortune-telling/models"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type OrderController struct {
}

type orderListForm struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Key      string `form:"key"`
	Status   string `form:"status"`
	Type     string `form:"type"`
}

type returnList struct {
	Id           int64         `json:"id"`
	Oid          string        `json:"oid"`
	Type         int           `json:"type"`
	Uid          int64         `json:"uid"`
	Username     string        `json:"username"`
	Gender       int           `json:"gender"`
	Amount       float64       `json:"amount"`
	OldAmount    float64       `json:"old_amount"`
	Coupon       int64         `json:"coupon"`
	PayType      int           `json:"pay_type"`
	TradeNo      string        `json:"trade_no"`
	Status       int           `json:"status"`
	Discount     float64       `json:"discount"`
	Phone        string        `json:"phone"`
	CreatedAt    models.MyTime `json:"created_at"`
	CouponAmount float64       `json:"coupon_amount"`
	Json         string        `json:"json"`
	PayTime      string        `json:"pay_time"`
	Birthday     string        `json:"birthday"`
	Content      string        `json:"content"`
}

// List 用户列表
func (con OrderController) List(c *gin.Context) {
	var data orderListForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var list []returnList

	var count int64
	res := ay.Db.Table("sm_order").
		Select("sm_user.phone,sm_order.*,sm_user_coupon.amount as coupon_amount").
		Joins("left join sm_user on sm_order.uid=sm_user.id").
		Joins("left join sm_user_coupon on sm_order.coupon=sm_user_coupon.id")

	if data.Key != "" {
		res.Where("sm_order.oid like ? OR sm_order.trade_no like ? OR sm_order.username like ? OR sm_user.phone like ?", "%"+data.Key+"%", "%"+data.Key+"%", "%"+data.Key+"%", "%"+data.Key+"%")
	}

	if data.Status != "" {
		res.Where("sm_order.status = ?", data.Status)
	}

	if data.Type != "" {
		res.Where("sm_order.type = ?", data.Type)
	}

	row := res

	row.Count(&count)

	res.Order("created_at desc").
		Limit(data.PageSize).
		Offset((data.Page - 1) * data.PageSize).
		Find(&list)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list":  list,
		"total": count,
	})
}

type orderDetailForm struct {
	Id int `form:"id"`
}

// Detail 用户详情
func (con OrderController) Detail(c *gin.Context) {
	var data orderDetailForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var user models.Order

	ay.Db.First(&user, data.Id)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"info": user,
	})
}

type orderOptionForm struct {
	Id     int `form:"id"`
	Status int `form:"status"`
}

// Option 添加 编辑
func (con OrderController) Option(c *gin.Context) {
	var data orderOptionForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var order models.Order
	ay.Db.First(&order, data.Id)

	if order.Status == 3 {
		ay.Json{}.Msg(c, 400, "请勿重新修改退款状态", gin.H{})
		return
	}

	if data.Status == 3 {
		var user models.User
		ay.Db.First(&user, order.Uid)
		user.Amount = user.Amount + order.Amount
		ay.Db.Save(&user)
	}

	order.Status = data.Status
	order.PayTime = time.Now().Format("2006-01-02 15:04:05")

	if err := ay.Db.Save(&order).Error; err != nil {
		ay.Json{}.Msg(c, 400, "修改失败", gin.H{})
	} else {
		ay.Json{}.Msg(c, 200, "修改成功", gin.H{})
	}

}

type orderDeleteForm struct {
	Id string `form:"id"`
}

func (con OrderController) Delete(c *gin.Context) {
	var data orderDeleteForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	idArr := strings.Split(data.Id, ",")

	for _, v := range idArr {
		var order models.Order
		ay.Db.Delete(&order, v)
	}

	ay.Json{}.Msg(c, 200, "删除成功", gin.H{})
}
