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

type NoticeController struct {
}

type noticeListForm struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Key      string `form:"key"`
	Status   string `form:"status"`
	Type     string `form:"type"`
}

// List 列表
func (con NoticeController) List(c *gin.Context) {
	var data noticeListForm
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

	var list []returnList

	var count int64
	res := ay.Db.Table("sm_notice").
		Select("sm_notice.*,sm_notice_type.name as type_name").
		Joins("left join sm_notice_type on sm_notice.type=sm_notice_type.id")

	if data.Key != "" {
		res.Where("sm_notice.title like ? OR sm_notice.keywords like ? OR sm_notice.content like ? OR sm_notice.description like ?", "%"+data.Key+"%", "%"+data.Key+"%", "%"+data.Key+"%", "%"+data.Key+"%")
	}

	if data.Status != "" {
		res.Where("sm_notice.status = ?", data.Status)
	}

	if data.Type != "" {
		res.Where("sm_notice.type = ?", data.Type)
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

//type orderDetailForm struct {
//	Id int `form:"id"`
//}

// Detail 用户详情
func (con NoticeController) Detail(c *gin.Context) {
	var data orderDetailForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var user models.NewsNotice

	ay.Db.First(&user, data.Id)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"info": user,
	})
}

type noticeOptionForm struct {
	Id          int    `form:"id"`
	Title       string `form:"title"`
	Keywords    string `form:"keywords"`
	Description string `form:"description"`
	Cover       string `form:"cover"`
	Content     string `form:"content"`
	Status      int    `form:"status"`
	Type        int64  `form:"type"`
}

// Option 添加 编辑
func (con NoticeController) Option(c *gin.Context) {
	var data noticeOptionForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var res models.NewsNotice
	ay.Db.First(&res, data.Id)

	if data.Id != 0 {

		res.Type = data.Type
		res.Title = data.Title
		res.Description = data.Description
		res.Content = data.Content
		res.Cover = data.Cover
		res.Keywords = data.Keywords
		res.Status = data.Status

		if err := ay.Db.Save(&res).Error; err != nil {
			ay.Json{}.Msg(c, 400, "修改失败", gin.H{})
		} else {
			ay.Json{}.Msg(c, 200, "修改成功", gin.H{})
		}
	} else {
		ay.Db.Create(&models.NewsNotice{
			Title:       data.Title,
			Keywords:    data.Keywords,
			Cover:       data.Cover,
			Content:     data.Content,
			Description: data.Description,
			Status:      data.Status,
			Type:        data.Type,
		})
		ay.Json{}.Msg(c, 200, "创建成功", gin.H{})

	}

}

func (con NoticeController) Delete(c *gin.Context) {
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
		var order models.NewsNotice
		ay.Db.Delete(&order, v)
	}

	ay.Json{}.Msg(c, 200, "删除成功", gin.H{})
}

func (con NoticeController) Upload(c *gin.Context) {

	code, msg := Upload(c, "notice")

	if code != 200 {
		ay.Json{}.Msg(c, 400, msg, gin.H{})
	} else {
		ay.Json{}.Msg(c, 200, msg, gin.H{})
	}
}
