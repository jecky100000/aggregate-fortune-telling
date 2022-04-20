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

type AreaController struct {
}

type GetAreaGetForm struct {
	Id int `form:"id"`
}

func (con AreaController) Get(c *gin.Context) {
	var getForm GetAreaGetForm
	if err := c.ShouldBind(&getForm); err != nil {
		Json.Msg(400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	//id := getForm.Id
	//if id == 0 {
	//	id = 0
	//}
	var area []models.Area

	if err := ay.Db.Where("pid = ?", getForm.Id).Find(&area).Error; err != nil {
		Json.Msg(400, "查询失败", gin.H{})
	} else {
		Json.Msg(200, "success", gin.H{
			"list": area,
		})
	}

}
