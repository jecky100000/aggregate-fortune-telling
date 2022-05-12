/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package api

import (
	"gin/ay"
	"gin/models"
	"github.com/gin-gonic/gin"
)

type InviteController struct {
}

type InviteListForm struct {
	Page int `form:"page" binding:"required" label:"页码"`
}

// List 邀请记录
func (con InviteController) List(c *gin.Context) {
	var getForm InviteListForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	page := getForm.Page - 1

	var user models.User
	ay.Db.First(&user, "id = ?", GetToken(Token))

	if user.Id == 0 {
		ay.Json{}.Msg(c, 401, "Token错误", gin.H{})
		return
	}

	type returnList struct {
		Nickname  string        `json:"nickname"`
		Phone     string        `json:"phone"`
		Avatar    string        `json:"avatar"`
		CreatedAt models.MyTime `json:"-"`
		Uid       int64         `json:"-"`
		InviteAt  string        `json:"invite_at"`
	}

	var list []returnList

	ay.Db.Table("sm_user_invite").
		Select("sm_user_invite.created_at,sm_user.nickname,sm_user.phone,sm_user.avatar").
		Joins("left join sm_user on sm_user_invite.uid=sm_user.id").
		Where("sm_user_invite.pid = ?", user.Id).
		Order("sm_user_invite.created_at desc").
		Limit(10).
		Offset(page * 10).
		Find(&list)

	for k, v := range list {
		list[k].Avatar = ay.Yaml.GetString("domain") + v.Avatar
		list[k].InviteAt = v.CreatedAt.Format("2006/01/02")
	}

	if list != nil {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": list,
		})
	} else {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": []string{},
		})
	}

}

// Withdraw 下级用户消费记录
func (con InviteController) Withdraw(c *gin.Context) {
	var getForm InviteListForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	page := getForm.Page - 1

	var user models.User
	ay.Db.First(&user, "id = ?", GetToken(Token))

	if user.Id == 0 {
		ay.Json{}.Msg(c, 401, "Token错误", gin.H{})
		return
	}

	type returnList struct {
		Nickname  string        `json:"nickname"`
		Phone     string        `json:"phone"`
		Avatar    string        `json:"avatar"`
		Amount    float64       `json:"amount"`
		CreatedAt models.MyTime `json:"-"`
		Time      string        `json:"created_at"`
		Status    int           `json:"status"`
		Oid       string        `json:"oid"`
	}

	var list []returnList

	ay.Db.Table("sm_user_invite_consumption").
		Select("sm_user.nickname,sm_user.phone,sm_user.avatar,sm_user_invite_consumption.created_at,sm_user_invite_consumption.amount,sm_user_invite_consumption.status,sm_user_invite_consumption.oid").
		Joins("left join sm_user on sm_user_invite_consumption.uid=sm_user.id").
		Where("sm_user_invite_consumption.pid = ? and sm_user_invite_consumption.status != 3 ", user.Id).
		Order("sm_user_invite_consumption.created_at desc").
		Limit(10).
		Offset(page * 10).
		Find(&list)

	for k, v := range list {
		list[k].Avatar = ay.Yaml.GetString("domain") + v.Avatar
		list[k].Time = v.CreatedAt.Format("2006/01/02")
	}

	if list != nil {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": list,
		})
	} else {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": []string{},
		})
	}

}

type InviteDoForm struct {
	Amount float64 `form:"amount" binding:"required" label:"金额"`
}

// Do 钱包划转
func (con InviteController) Do(c *gin.Context) {
	var data InviteDoForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var user models.User
	ay.Db.First(&user, "id = ?", GetToken(Token))

	if user.Id == 0 {
		ay.Json{}.Msg(c, 401, "Token错误", gin.H{})
		return
	}

	//if data.Amount < 100 {
	//	ay.Json{}.Msg(c, 400, "提现金额需要大于100", gin.H{})
	//	return
	//}

	if user.InviteAmount < data.Amount {
		ay.Json{}.Msg(c, 400, "余额不足", gin.H{})
		return
	}
	OldAmount := user.Amount
	OldInviteAmount := user.InviteAmount

	nowInviteAmount := OldInviteAmount - data.Amount

	nowAmount := OldAmount + data.Amount

	user.Amount = nowAmount
	user.InviteAmount = nowInviteAmount

	if err := ay.Db.Save(&user); err != nil {
		// 划转记录
		ay.Db.Create(&models.UserTransfer{
			Amount:          data.Amount,
			OldAmount:       OldAmount,
			NowAmount:       nowAmount,
			OldInviteAmount: OldInviteAmount,
			NowInviteAmount: nowInviteAmount,
			Uid:             user.Id,
		})
		//
		ay.Json{}.Msg(c, 200, "提现成功", gin.H{})
	} else {
		ay.Json{}.Msg(c, 400, "提现失败", gin.H{})

	}

}
