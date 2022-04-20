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

type NewsController struct {
}

type GetNewsDetailForm struct {
	Id int `form:"id" binding:"required"`
}

// Type 获取所有文章类型
func (con NewsController) Type(c *gin.Context) {
	var getForm GetNewsDetailForm
	if err := c.ShouldBind(&getForm); err != nil {
		Json.Msg(400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var res []models.NewsType
	ay.Db.Find(&res, "pid = ?", getForm.Id)
	Json.Msg(200, "success", gin.H{
		"list": res,
	})
}

type GetNewsAllForm struct {
	Id   int64 `form:"id" binding:"required"`
	Page int   `form:"id" binding:"required" label:"页码"`
}

// All 获取分类下所有文章
func (con NewsController) All(c *gin.Context) {
	var getForm GetNewsAllForm
	if err := c.ShouldBind(&getForm); err != nil {
		Json.Msg(400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	page := getForm.Page - 1

	var res []models.NewsNotice
	ay.Db.Limit(10).Offset(page*10).Order("id desc").Find(&res, "type = ? and status = 1", getForm.Id)
	Json.Msg(200, "success", gin.H{
		"list": res,
	})
}

// Detail 文章详情
func (con NewsController) Detail(c *gin.Context) {
	var getForm GetNewsDetailForm
	if err := c.ShouldBind(&getForm); err != nil {
		Json.Msg(400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	res := models.NewsModel{}.GetDetail(getForm.Id)
	Json.Msg(200, "success", gin.H{
		"info": res,
	})
}

type GetNewsRecommendForm struct {
	Type int `form:"type" binding:"required" label:"页码"`
}

// Recommend 文章详情
func (con NewsController) Recommend(c *gin.Context) {
	var getForm GetNewsRecommendForm
	if err := c.ShouldBind(&getForm); err != nil {
		Json.Msg(400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	notice := models.NewsModel{}.GetList(getForm.Type)

	for k, v := range notice {
		notice[k].Time = ay.LastTime(int(v.CreatedAt.Unix()))
	}

	Json.Msg(200, "success", gin.H{
		"info": notice,
	})
}
