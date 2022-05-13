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
	"strings"
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

type Master struct {
	Id int64
	models.Master
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
}

// List 获取类型下大师
func (con MasterController) List(c *gin.Context) {
	var getForm GetMasterListForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}
	page := getForm.Page - 1

	//field := "id,name,sign,type,years,online,avatar,rate"
	//var res []models.Master
	//ay.Db.Where("FIND_IN_SET(?,type)", getForm.Type).Select(field).Limit(10).Offset(page * 10).Order("id desc").Find(&res)

	var res []Master
	ay.Db.Table("sm_user").
		Select("sm_user.id,sm_user.phone,sm_user.nickname,sm_master.sign,sm_master.type,sm_master.years,sm_master.online,sm_user.avatar,sm_master.rate").
		Joins("left join sm_master on sm_user.master_id=sm_master.id").
		Where("FIND_IN_SET(?,sm_master.type) and sm_user.type=1", getForm.Type).
		Limit(10).
		Offset(page * 10).
		Order("sm_user.id desc").
		Find(&res)

	var row []map[string]interface{}

	for _, v := range res {
		var typeName []string

		for _, v := range strings.Split(v.Master.Type, ",") {
			var masterType models.MasterType
			ay.Db.First(&masterType, "id = ?", v)
			if masterType.Name != "" {
				typeName = append(typeName, masterType.Name)
			}

		}

		row = append(row, map[string]interface{}{
			"id":        v.Id,
			"name":      v.Nickname,
			"sign":      v.Master.Sign,
			"years":     v.Master.Years,
			"online":    v.Master.Online,
			"avatar":    ay.Yaml.GetString("domain") + v.Avatar,
			"rate":      v.Master.Rate,
			"type_name": typeName,
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

	//field := "id,name,sign,type,years,online,avatar,rate"
	//var res []models.Master
	//ay.Db.Where("FIND_IN_SET(?,type)", getForm.Type).Select(field).Limit(10).Offset(page * 10).Order("id desc").Find(&res)
	var res []Master
	ay.Db.Table("sm_user").
		Select("sm_user.id,sm_user.phone,sm_user.nickname,sm_master.sign,sm_master.type,sm_master.years,sm_master.online,sm_user.avatar,sm_master.rate").
		Joins("left join sm_master on sm_user.master_id=sm_master.id").
		Where("is_recommend = 1").
		Limit(10).
		Offset(page * 10).
		Order("sm_user.id desc").
		Find(&res)

	var row []map[string]interface{}

	for _, v := range res {
		var typeName []string

		for _, v := range strings.Split(v.Master.Type, ",") {
			var masterType models.MasterType
			ay.Db.First(&masterType, "id = ?", v)
			if masterType.Name != "" {
				typeName = append(typeName, masterType.Name)
			}

		}

		row = append(row, map[string]interface{}{
			"id":        v.Id,
			"name":      v.Nickname,
			"sign":      v.Master.Sign,
			"years":     v.Master.Years,
			"online":    v.Master.Online,
			"avatar":    ay.Yaml.GetString("domain") + v.Avatar,
			"rate":      v.Master.Rate,
			"type_name": typeName,
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

	var optionUser models.User
	ay.Db.First(&optionUser, "id = ?", GetToken(Token))
	models.UserMasterLogModel{}.Save(optionUser.Id, user.Id)

	type master struct {
		models.Master
		//ImageS   []string `json:"images"`
		Avatar   string   `json:"avatar"`
		Name     string   `json:"name"`
		TypeName []string `json:"type_name"`
		Phone    string   `json:"phone"`
	}
	var res master
	//ay.Db.Model(&models.Master{}).Where("id = ?", user.MasterId).First(&res)

	ay.Db.Table("sm_master").
		Select("sm_master.*,sm_user.nickname as name,sm_user.avatar,sm_user.phone").
		Joins("left join sm_user on sm_master.id=sm_user.master_id").
		Where("sm_master.id = ?", user.MasterId).
		First(&res)

	var typeName []string

	for _, v := range strings.Split(res.Type, ",") {
		var masterType models.MasterType
		ay.Db.First(&masterType, "id = ?", v)
		if masterType.Name != "" {
			typeName = append(typeName, masterType.Name)
		}
	}
	res.TypeName = typeName

	if res.Id == 0 {
		ay.Json{}.Msg(c, 400, "大师不存在", gin.H{})
		return
	}

	res.Avatar = ay.Yaml.GetString("domain") + res.Avatar
	res.BackImage = ay.Yaml.GetString("domain") + res.BackImage

	res.Id = user.Id

	//var image []string
	//
	//json.Unmarshal([]byte(res.Image), &image)
	//
	//for k, v := range image {
	//	image[k] = ay.Yaml.GetString("domain") + v
	//}
	//
	//res.ImageS = image
	//res.Image = ""

	// 粉丝
	var count int64
	ay.Db.Model(models.Collect{}).Where("type = 1 and cid = ?", res.Id).Count(&count)
	res.Fans = 50000 + count

	// 回复
	var count1 int64
	ay.Db.Model(models.AskReply{}).Where("master_id = ?", res.Id).Count(&count1)
	res.Reply = 20000 + count1

	// 当前用户
	uid := GetToken(Token)

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
		"info": res,
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
