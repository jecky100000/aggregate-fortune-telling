/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package admin

import (
	"aggregate-fortune-telling/ay"
	"aggregate-fortune-telling/controllers/api"
	"aggregate-fortune-telling/models"
	"github.com/gin-gonic/gin"
	"strings"
)

type AccountController struct {
}

type listForm struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Key      string `form:"key"`
	Type     string `form:"type"`
}

// List 用户列表
func (con AccountController) List(c *gin.Context) {
	var data listForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var count int64
	var user []models.User

	res := ay.Db.Table("sm_user")

	if data.Key != "" {
		res.Where("phone like ? or nickname like ?", "%"+data.Key+"%", "%"+data.Key+"%")
	}

	if data.Type != "" {
		res.Where("type = ?", data.Type)
	}

	res.Count(&count)

	res.Order("created_at desc").
		Limit(data.PageSize).
		Offset((data.Page - 1) * data.PageSize).
		Find(&user)

	ay.Json{}.Msg(c, 200, "success", gin.H{
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
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var user models.User

	ay.Db.First(&user, data.Id)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"info": user,
	})
}

type optionForm struct {
	Id       int     `form:"id"`
	Type     int     `form:"type"`
	Phone    string  `form:"phone"`
	Nickname string  `form:"nickname"`
	Amount   float64 `form:"amount"`
	Avatar   string  `form:"avatar"`
}

// Option 添加 编辑
func (con AccountController) Option(c *gin.Context) {
	var data optionForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
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
				ay.Json{}.Msg(c, 400, "手机已存在", gin.H{})
				return
			}
		}

		user.Phone = data.Phone
		user.Amount = data.Amount
		user.Type = data.Type
		user.NickName = data.Nickname
		user.Avatar = data.Avatar

		if err := ay.Db.Save(&user).Error; err != nil {
			ay.Json{}.Msg(c, 400, "请联系管理员", gin.H{})
			return
		} else {
			ay.Json{}.Msg(c, 200, "修改成功", gin.H{})
			return
		}

	} else {
		var phoneNum int64
		ay.Db.Model(&models.User{}).Where("phone = ?", data.Phone).Count(&phoneNum)
		if phoneNum != 0 {
			ay.Json{}.Msg(c, 400, "手机已存在", gin.H{})
			return
		}
		if err := ay.Db.Create(&models.User{
			Type:     data.Type,
			Amount:   data.Amount,
			Phone:    data.Phone,
			Avatar:   data.Avatar,
			NickName: data.Nickname,
		}).Error; err != nil {
			ay.Json{}.Msg(c, 400, "请联系管理员", gin.H{})
			return
		} else {
			r, rs := api.LoginController{}.MakeImAccount(data.Phone)
			if r != 1 {
				ay.Json{}.Msg(c, 400, rs, gin.H{})
				return
			}
			ay.Json{}.Msg(c, 200, "创建成功", gin.H{})
			return
		}

	}

}

type deleteForm struct {
	Id string `form:"id"`
}

func (con AccountController) Delete(c *gin.Context) {
	var data deleteForm
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
		var user models.User
		ay.Db.First(&user, v)
		ay.Db.Delete(&user)
		ay.Db.Delete(&models.Master{}, user.MasterId)
	}

	ay.Json{}.Msg(c, 200, "删除成功", gin.H{})
}

func (con AccountController) Upload(c *gin.Context) {

	code, msg := Upload(c, "user")

	if code != 200 {
		ay.Json{}.Msg(c, 400, msg, gin.H{})
	} else {
		ay.Json{}.Msg(c, 200, msg, gin.H{})
	}
}
