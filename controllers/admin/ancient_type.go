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

type AncientTypeController struct {
}

type ancientTypeListForm struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Key      string `form:"key"`
	Status   string `form:"status"`
	Type     string `form:"type"`
}

// List 列表
func (con AncientTypeController) List(c *gin.Context) {
	var data noticeTypeListForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
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

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list":  list,
		"total": count,
	})
}

//type orderDetailForm struct {
//	Id int `form:"id"`
//}

// Detail 用户详情
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

	var res models.NewsType

	ay.Db.First(&res, data.Id)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"info": res,
	})
}

type ancientTypeOptionForm struct {
	Id   int    `form:"id"`
	Name string `form:"name"`
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

	var res models.NewsType
	ay.Db.First(&res, data.Id)

	if data.Id != 0 {
		res.Name = data.Name

		ay.Db.Save(&res)
		ay.Json{}.Msg(c, 200, "修改成功", gin.H{})
	} else {

		ay.Db.Create(&models.NewsType{
			Name: data.Name,
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
		var res models.NewsType
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
