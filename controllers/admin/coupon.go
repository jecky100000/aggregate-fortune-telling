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
	"log"
	"strings"
	"time"
)

type CouponController struct {
}

type couponListForm struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Key      string `form:"key"`
	Status   string `form:"status"`
	Type     string `form:"type"`
	Uid      int64  `form:"uid"`
}

// List 用户列表
func (con CouponController) List(c *gin.Context) {
	var data couponListForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	type returnList struct {
		models.Coupon
		Phone    string `json:"phone"`
		TypeName string `json:"type_name"`
	}

	var list []returnList

	var count int64
	res := ay.Db.Table("sm_user_coupon").
		Select("sm_user.phone,sm_user_coupon.*").
		Joins("left join sm_user on sm_user_coupon.uid=sm_user.id").
		Where("sm_user_coupon.uid =?", data.Uid)

	row := res

	row.Count(&count)

	res.Order("created_at desc").
		Limit(data.PageSize).
		Offset((data.Page - 1) * data.PageSize).
		Find(&list)

	for k, _ := range list {
		list[k].TypeName = "测算"
	}

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list":  list,
		"total": count,
	})
}

// Detail 详情
func (con CouponController) Detail(c *gin.Context) {
	var data orderDetailForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var user models.Coupon

	ay.Db.First(&user, data.Id)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"info": user,
	})
}

type couponOption struct {
	Id          int64   `form:"id"`
	Uid         int64   `form:"uid"`
	Name        string  `form:"name"`
	SubName     string  `form:"sub_name"`
	Prohibit    string  `form:"prohibit"`
	UsedAt      string  `form:"used_at"`
	Effective   string  `form:"effective"`
	EffectiveAt string  `form:"effective_at"`
	Amount      float64 `form:"amount"`
	Product     string  `form:"product"`
	Status      int     `form:"status"`
	AmountThan  float64 `form:"amount_than"`
}

// Option 添加 编辑
func (con CouponController) Option(c *gin.Context) {
	var data couponOption
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}
	log.Println(data)

	var res models.Coupon
	ay.Db.First(&res, data.Id)

	stamp, _ := time.ParseInLocation("2006-01-02 15:04:05", data.EffectiveAt, time.Local)

	if data.Id != 0 {

		res.Id = data.Id
		res.Uid = data.Uid
		res.Status = data.Status
		res.Name = data.Name
		res.SubName = data.SubName
		res.Effective = data.Effective
		//res.EffectiveAt = data.EffectiveAt
		res.EffectiveAt = models.MyTime{Time: stamp}
		res.Product = data.Product
		res.AmountThan = data.AmountThan
		res.Amount = data.Amount
		res.Prohibit = data.Prohibit

		if err := ay.Db.Save(&res).Error; err != nil {
			ay.Json{}.Msg(c, 400, "修改失败", gin.H{})
		} else {
			ay.Json{}.Msg(c, 200, "修改成功", gin.H{})
		}
	} else {
		ay.Db.Create(&models.Coupon{
			Uid:         data.Uid,
			Name:        data.Name,
			SubName:     data.SubName,
			Prohibit:    data.Prohibit,
			Effective:   data.Effective,
			EffectiveAt: models.MyTime{Time: stamp},
			Amount:      data.Amount,
			Product:     data.Product,
			Status:      data.Status,
			AmountThan:  data.AmountThan,
		})
		ay.Json{}.Msg(c, 200, "创建成功", gin.H{})

	}

}

func (con CouponController) Delete(c *gin.Context) {
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
		var order models.Coupon
		ay.Db.Delete(&order, v)
	}

	ay.Json{}.Msg(c, 200, "删除成功", gin.H{})
}
