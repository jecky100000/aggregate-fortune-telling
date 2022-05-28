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
)

type AncientTypeController struct {
}

// List 列表
func (con AncientTypeController) List(c *gin.Context) {

	type listForm struct {
		Page     int    `form:"page"`
		PageSize int    `form:"pageSize"`
		Key      string `form:"key"`
	}

	var data listForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	type rList struct {
		models.AncientType
		Children []models.AncientType `json:"children"`
	}
	var list []rList

	var count int64
	ay.Db.Model(&models.AncientType{}).
		Where("pid = 0").
		Order("id desc").
		Limit(data.PageSize).
		Offset((data.Page - 1) * data.PageSize).
		Find(&list)

	for k, v := range list {
		var xj []models.AncientType
		ay.Db.Where("pid = ?", v.Id).Find(&xj)
		list[k].Children = xj
	}

	ay.Db.Model(&models.AncientType{}).
		Where("pid = 0").
		Count(&count)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list":  list,
		"total": count,
	})
}

// Detail 详情
func (con AncientTypeController) Detail(c *gin.Context) {
	var data orderDetailForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var res models.AncientType

	ay.Db.First(&res, data.Id)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"info": res,
	})
}

type ancientTypeOptionForm struct {
	Id   int64  `form:"id"`
	Name string `form:"name"`
	Pid  int64  `form:"pid"`
}

// Option 添加 编辑
func (con AncientTypeController) Option(c *gin.Context) {
	var data ancientTypeOptionForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var res models.AncientType
	ay.Db.First(&res, data.Id)

	if data.Id != 0 {
		res.Name = data.Name
		res.Pid = data.Pid

		if err := ay.Db.Save(&res).Error; err != nil {
			ay.Json{}.Msg(c, 400, "修改失败", gin.H{})
		} else {
			ay.Json{}.Msg(c, 200, "修改成功", gin.H{})
		}
	} else {

		ay.Db.Create(&models.AncientType{
			Name: data.Name,
			Pid:  data.Pid,
		})
		ay.Json{}.Msg(c, 200, "创建成功", gin.H{})

	}

}

func (con AncientTypeController) Delete(c *gin.Context) {
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
		var res models.AncientType
		ay.Db.Delete(&res, v)
	}

	ay.Json{}.Msg(c, 200, "删除成功", gin.H{})
}

func (con AncientTypeController) All(c *gin.Context) {
	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}
	pid := c.Query("pid")

	type list struct {
		Label string `gorm:"column:name" json:"label"`
		Value int64  `gorm:"column:id" json:"value"`
	}
	var l []list
	res := ay.Db.Table("sm_ancient_type")

	if pid != "" {
		res.Where("pid = ? ", pid)
	} else {
		res.Where("pid = ? ", 0)

	}

	res.Find(&l)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list": l,
	})
}
