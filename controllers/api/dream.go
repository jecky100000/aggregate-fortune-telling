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

type DreamController struct {
}

// Main 首页
func (con DreamController) Main(c *gin.Context) {
	res := models.DreamTypeModel{}.GetAllType()
	recommend := models.DreamModel{}.GetList("0", 3)

	for k, v := range recommend {
		recommend[k].Message = ay.Summary(v.Message, 50)
	}

	Json.Msg(200, "success", gin.H{
		"list":      res,
		"recommend": recommend,
		//"notice":    notice,
	})
}

type GetDreamForm struct {
	Title string `form:"title" binding:"required"`
}

// Search 搜索
func (con DreamController) Search(c *gin.Context) {
	var getForm GetDreamForm
	if err := c.ShouldBind(&getForm); err != nil {
		Json.Msg(400, ay.Validator{}.Translate(err), gin.H{})
		return
	}
	res := models.DreamModel{}.GetList(getForm.Title, 10)

	for k, v := range res {
		res[k].Message = ay.Summary(v.Message, 50)
	}

	Json.Msg(200, "success", gin.H{
		"list": res,
	})
}

type GetDreamDetailForm struct {
	Id int `form:"id" binding:"required"`
}

// Detail 详情
func (con DreamController) Detail(c *gin.Context) {
	var getForm GetDreamDetailForm
	if err := c.ShouldBind(&getForm); err != nil {
		Json.Msg(400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	res := models.DreamModel{}.GetDetail(getForm.Id)

	total := models.DreamTotalModel{}.GetTotal(getForm.Id)

	Json.Msg(200, "success", gin.H{
		"info":  res,
		"total": total.Num,
		//"notice": notice,
	})
}
