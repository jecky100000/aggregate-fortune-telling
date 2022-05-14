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

type MasterController struct {
}

// Type 获取所有类型
func (con MasterController) Type(c *gin.Context) {

	var t []models.MasterType
	ay.Db.Find(&t)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list": t,
	})
}

type GetMasterListForm struct {
	Type int `form:"type"`
	Page int `form:"page"`
}

// List 获取类型下大师
func (con MasterController) List(c *gin.Context) {
	var getForm GetMasterListForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}
	page := getForm.Page - 1

	var row []gin.H

	res := models.MasterModel{}.GetMasterPage(page, 0, getForm.Type)

	for _, v := range res {
		row = append(row, map[string]interface{}{
			"id":        v.Id,
			"name":      v.Nickname,
			"sign":      v.Master.Sign,
			"years":     v.Master.Years,
			"online":    v.Master.Online,
			"avatar":    v.Avatar,
			"rate":      v.Master.Rate,
			"type_name": v.TypeName,
			"phone":     v.Phone,
		})
	}

	if row == nil {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": []string{},
		})
	} else {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": row,
		})
	}
}

// GetRecommend 获取推荐下大师
func (con MasterController) GetRecommend(c *gin.Context) {
	var getForm GetMasterListForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}
	page := getForm.Page - 1

	var row []gin.H

	res := models.MasterModel{}.GetMasterPage(page, 1, 0)

	for _, v := range res {
		row = append(row, map[string]interface{}{
			"id":        v.Id,
			"name":      v.Nickname,
			"sign":      v.Master.Sign,
			"years":     v.Master.Years,
			"online":    v.Master.Online,
			"avatar":    v.Avatar,
			"rate":      v.Master.Rate,
			"type_name": v.TypeName,
			"phone":     v.Phone,
		})
	}

	if row == nil {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": []string{},
		})
	} else {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": row,
		})
	}
}

type GetMasterDetailForm struct {
	Id int64 `form:"id"`
}

// Detail 获取大师详情
func (con MasterController) Detail(c *gin.Context) {
	var getForm GetMasterDetailForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	rj, user, _ := models.MasterModel{}.IsMaser(getForm.Id)
	if !rj {
		ay.Json{}.Msg(c, 400, "大师不存在", gin.H{})
		return
	}

	// 当前用户
	uid := GetToken(Token)

	var optionUser models.User
	ay.Db.First(&optionUser, "id = ?", uid)
	models.UserMasterLogModel{}.Save(optionUser.Id, user.Id)

	isMaster, res := models.MasterModel{}.GetMaster(getForm.Id)

	if !isMaster {
		ay.Json{}.Msg(c, 400, "大师不存在", gin.H{})
		return
	}
	// 粉丝
	var count int64
	ay.Db.Model(models.Collect{}).Where("type = 1 and cid = ?", res.Id).Count(&count)
	fans := 50000 + count

	// 回复
	var count1 int64
	ay.Db.Model(models.AskReply{}).Where("master_id = ?", res.Id).Count(&count1)
	reply := 20000 + count1

	// 推荐
	var recommendMaster models.UserRecommendMaster
	ay.Db.Where("uid = ? AND master_id = ?", uid, res.Id).First(&recommendMaster)

	isRecommend, isCollect := 0, 0

	if recommendMaster.Id != 0 {
		isRecommend = 1
	}

	// 收藏
	var collect models.Collect
	ay.Db.First(&collect, "uid = ? and type = ? and cid = ?", uid, 1, res.Id)

	if collect.Id != 0 {
		isCollect = 1
	}

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"info": gin.H{
			"ask_num":      res.AskNum,
			"avatar":       res.Avatar,
			"back_image":   res.BackImage,
			"created_at":   res.CreatedAt,
			"fans":         fans,
			"id":           res.Id,
			"introduce":    res.Introduce,
			"is_recommend": res.IsRecommend,
			"label":        res.Label,
			"name":         res.Nickname,
			"online":       res.Online,
			"phone":        res.Phone,
			"rate":         res.Rate,
			"reply":        reply,
			"sign":         res.Sign,
			"type_name":    res.TypeName,
			"years":        res.Years,
		},
		"user": gin.H{
			"is_recommend": isRecommend,
			"is_collect":   isCollect,
		},
	})
}

// Recommend 用户大师推荐
func (con MasterController) Recommend(c *gin.Context) {
	var getForm GetMasterDetailForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var user models.User
	ay.Db.First(&user, "id = ?", GetToken(Token))

	if user.Id == 0 {
		ay.Json{}.Msg(c, 401, "Token错误", gin.H{})
		return
	}

	rj, _, _ := models.MasterModel{}.IsMaser(getForm.Id)
	if !rj {
		ay.Json{}.Msg(c, 400, "大师不存在", gin.H{})
		return
	}

	var masterRecommend models.UserRecommendMaster
	ay.Db.First(&masterRecommend, "uid = ? AND master_id = ?", user.Id, getForm.Id)

	if masterRecommend.Id == 0 {
		ss := models.UserRecommendMaster{
			Uid:      user.Id,
			MasterId: getForm.Id,
		}
		ay.Db.Create(&ss)
		ay.Json{}.Msg(c, 200, "推荐成功", gin.H{})
	} else {
		ay.Db.Delete(&masterRecommend)
		ay.Json{}.Msg(c, 200, "取消推荐", gin.H{})
	}
}
