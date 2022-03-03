/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package admin

import (
	"gin/ay"
	"gin/models"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
}

type GetAccountListForm struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Key      string `form:"key"`
}

func (con AccountController) List(c *gin.Context) {
	var data GetAccountListForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, "400", ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, "401", "请登入", gin.H{})
		return
	}

	var user []models.User

	var count int64
	if data.Key != "" {
		ay.Db.Model(models.User{}).Where("phone like ?", "%"+data.Key+"%").
			Count(&count).Find(&user)
		ay.Db.Limit(data.PageSize).Offset((data.Page-1)*data.PageSize).
			Where("phone like ?", "%"+data.Key+"%")
	} else {
		ay.Db.Model(models.User{}).Where("phone like ?", "%"+data.Key+"%").
			Count(&count).Find(&user)
		ay.Db.Limit(data.PageSize).Offset((data.Page - 1) * data.PageSize)
	}

	ay.Json{}.Msg(c, "200", "success", gin.H{
		"list":  user,
		"total": count,
	})
}
