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
	"strings"
)

type AccountController struct {
}

type listForm struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Key      string `form:"key"`
}

// List 用户列表
func (con AccountController) List(c *gin.Context) {
	var data listForm
	if err := c.ShouldBind(&data); err != nil {
		Json.Msg(400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		Json.Msg(401, "请登入", gin.H{})
		return
	}

	var user []models.User

	var count int64
	if data.Key != "" {
		ay.Db.Model(models.User{}).
			Where("phone like ?", "%"+data.Key+"%").
			Count(&count)
		ay.Db.Model(models.User{}).
			Where("phone like ?", "%"+data.Key+"%").
			Order("created_at desc").
			Limit(data.PageSize).
			Offset((data.Page - 1) * data.PageSize).
			Find(&user)
	} else {
		ay.Db.Model(models.User{}).
			Count(&count)
		ay.Db.Model(models.User{}).
			Limit(data.PageSize).
			Order("created_at desc").
			Offset((data.Page - 1) * data.PageSize).
			Find(&user)
	}

	Json.Msg(200, "success", gin.H{
		"list":  user,
		"total": count,
	})
}

type detailForm struct {
	Id int `form:"id"`
}

// Detail 用户详情
func (con AccountController) Detail(c *gin.Context) {
	var data detailForm
	if err := c.ShouldBind(&data); err != nil {
		Json.Msg(400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		Json.Msg(401, "请登入", gin.H{})
		return
	}

	var user models.User

	ay.Db.First(&user, data.Id)

	Json.Msg(200, "success", gin.H{
		"info": user,
	})
}

type optionForm struct {
	Id       int     `form:"id"`
	Type     int     `form:"type"`
	Phone    string  `form:"phone"`
	Nickname string  `form:"nickname"`
	Amount   float64 `form:"amount"`
}

// Option 添加 编辑
func (con AccountController) Option(c *gin.Context) {
	var data optionForm
	if err := c.ShouldBind(&data); err != nil {
		Json.Msg(400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		Json.Msg(401, "请登入", gin.H{})
		return
	}

	var user models.User
	ay.Db.First(&user, data.Id)

	if data.Id != 0 {
		if user.Phone != data.Phone {
			// 手机号变动
			var phoneNum int64
			ay.Db.Model(&models.User{}).Where("id != ? AND phone = ?", data.Id, data.Phone).Count(&phoneNum)
			if phoneNum != 0 {
				Json.Msg(400, "手机已存在", gin.H{})
				return
			}
		}

		user.Phone = data.Phone
		user.Amount = data.Amount
		user.Type = data.Type
		user.NickName = data.Nickname

		ay.Db.Save(&user)
		Json.Msg(200, "修改成功", gin.H{})
	} else {
		var phoneNum int64
		ay.Db.Model(&models.User{}).Where("phone = ?", data.Phone).Count(&phoneNum)
		if phoneNum != 0 {
			Json.Msg(400, "手机已存在", gin.H{})
			return
		}
		ay.Db.Create(&models.User{
			Type:     data.Type,
			Amount:   data.Amount,
			Phone:    data.Phone,
			Avatar:   "/static/user/default.png",
			NickName: data.Nickname,
		})
		Json.Msg(200, "创建成功", gin.H{})

	}

}

type deleteForm struct {
	Id string `form:"id"`
}

func (con AccountController) Delete(c *gin.Context) {
	var data deleteForm
	if err := c.ShouldBind(&data); err != nil {
		Json.Msg(400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		Json.Msg(401, "请登入", gin.H{})
		return
	}

	idArr := strings.Split(data.Id, ",")

	for _, v := range idArr {
		var user models.User
		ay.Db.Delete(&user, v)
	}

	Json.Msg(200, "删除成功", gin.H{})
}
