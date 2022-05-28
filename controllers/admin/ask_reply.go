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

type AskReplyController struct {
}

type askReplyListForm struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Key      string `form:"key"`
}

// List 列表
func (con AskReplyController) List(c *gin.Context) {
	var data askReplyListForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	type rj struct {
		Id          int64         `json:"id"`
		MasterPhone string        `json:"master_phone"`
		MasterName  string        `json:"master_name"`
		Adopt       int           `json:"adopt"`
		CreatedAt   models.MyTime `json:"created_at"`
		Content     string        `json:"content"`
	}

	var list []rj

	var count int64
	res := ay.Db.Table("sm_ask_reply").
		Select("sm_ask_reply.id,sm_ask_reply.content,sm_ask_reply.created_at,sm_ask_reply.adopt,sm_user.nickname as master_name,sm_user.phone as master_phone").
		Joins("left join sm_user on sm_ask_reply.master_id=sm_user.id").
		Where("sm_ask_reply.ask_id = ?", data.Key)

	row := res

	row.Count(&count)

	res.Order("sm_ask_reply.created_at desc").
		Limit(data.PageSize).
		Offset((data.Page - 1) * data.PageSize).
		Debug().
		Find(&list)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list":  list,
		"total": count,
	})
}

// Detail 用户详情
func (con AskReplyController) Detail(c *gin.Context) {
	var data orderDetailForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var user models.AskReply

	ay.Db.First(&user, data.Id)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"info": user,
	})
}

type askReplyOptionForm struct {
	Id       int    `form:"id"`
	AskId    string `form:"ask_id"`
	Content  string `form:"content"`
	MasterId int64  `form:"master_id"`
}

// Option 添加 编辑
func (con AskReplyController) Option(c *gin.Context) {
	var data askReplyOptionForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var res models.AskReply
	ay.Db.First(&res, data.Id)

	if data.Id != 0 {

		res.Content = data.Content
		res.AskId = data.AskId
		res.MasterId = data.MasterId

		if err := ay.Db.Save(&res).Error; err != nil {
			ay.Json{}.Msg(c, 400, "修改失败", gin.H{})
		} else {
			ay.Json{}.Msg(c, 200, "修改成功", gin.H{})
		}
	} else {
		ay.Db.Create(&models.AskReply{
			Content:  data.Content,
			MasterId: data.MasterId,
			AskId:    data.AskId,
		})
		ay.Json{}.Msg(c, 200, "创建成功", gin.H{})

	}

}

func (con AskReplyController) Delete(c *gin.Context) {
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
		var order models.AskReply
		ay.Db.Delete(&order, v)
	}

	ay.Json{}.Msg(c, 200, "删除成功", gin.H{})
}
