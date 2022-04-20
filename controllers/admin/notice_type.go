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

type NoticeTypeController struct {
}

type noticeTypeListForm struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Key      string `form:"key"`
	Status   string `form:"status"`
	Type     string `form:"type"`
}

// List 列表
func (con NoticeTypeController) List(c *gin.Context) {
	var data noticeTypeListForm
	if err := c.ShouldBind(&data); err != nil {
		Json.Msg(400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		Json.Msg(401, "请登入", gin.H{})
		return
	}

	type returnList struct {
		models.NewsNotice
		TypeName string `json:"typeName"`
	}

	var list []models.NewsType

	var count int64
	ay.Db.Table("sm_notice_type").
		Select("sm_notice_type.*").
		Order("id desc").
		Limit(data.PageSize).
		Offset((data.Page - 1) * data.PageSize).
		Find(&list)

	ay.Db.Table("sm_notice_type").Count(&count)

	Json.Msg(200, "success", gin.H{
		"list":  list,
		"total": count,
	})
}

//type orderDetailForm struct {
//	Id int `form:"id"`
//}

// Detail 用户详情
func (con NoticeTypeController) Detail(c *gin.Context) {
	var data orderDetailForm
	if err := c.ShouldBind(&data); err != nil {
		Json.Msg(400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		Json.Msg(401, "请登入", gin.H{})
		return
	}

	var res models.NewsType

	ay.Db.First(&res, data.Id)

	Json.Msg(200, "success", gin.H{
		"info": res,
	})
}

type noticeTypeOptionForm struct {
	Id   int    `form:"id"`
	Name string `form:"name"`
}

// Option 添加 编辑
func (con NoticeTypeController) Option(c *gin.Context) {
	var data noticeTypeOptionForm
	if err := c.ShouldBind(&data); err != nil {
		Json.Msg(400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		Json.Msg(401, "请登入", gin.H{})
		return
	}

	var res models.NewsType
	ay.Db.First(&res, data.Id)

	if data.Id != 0 {
		res.Name = data.Name

		ay.Db.Save(&res)
		Json.Msg(200, "修改成功", gin.H{})
	} else {

		ay.Db.Create(&models.NewsType{
			Name: data.Name,
		})
		Json.Msg(200, "创建成功", gin.H{})

	}

}

func (con NoticeTypeController) Delete(c *gin.Context) {
	var data orderDeleteForm
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
		var res models.NewsType
		ay.Db.Delete(&res, v)
	}

	Json.Msg(200, "删除成功", gin.H{})
}

func (con NoticeTypeController) All(c *gin.Context) {
	if Auth() == false {
		Json.Msg(401, "请登入", gin.H{})
		return
	}
	type list struct {
		Label string `gorm:"column:name" json:"label"`
		Value int64  `gorm:"column:id" json:"value"`
	}
	var l []list
	ay.Db.Table("sm_notice_type").Find(&l)

	Json.Msg(200, "success", gin.H{
		"list": l,
	})
}
