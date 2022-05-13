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

type CollectController struct {
}

type GetCollectSetForm struct {
	Type int   `form:"type" binding:"required" label:"类型"`
	Id   int64 `form:"id" binding:"required" `
}

func (con CollectController) Set(c *gin.Context) {
	var getForm GetCollectSetForm
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

	var collect models.Collect

	ay.Db.Select("id").First(&collect, "uid = ? and type = ? and cid = ?", user.Id, getForm.Type, getForm.Id)

	if collect.Id == 0 {

		if getForm.Type == 1 {

			rj, _, _ := models.MasterModel{}.IsMaser(getForm.Id)
			if !rj {
				ay.Json{}.Msg(c, 400, "大师不存在", gin.H{})
				return
			}

		} else if getForm.Type == 2 {
			var baike models.BaiKe
			ay.Db.Select("id").First(&baike, "id = ?", getForm.Id)
			if baike.Id == 0 {
				ay.Json{}.Msg(c, 400, "百科不存在", gin.H{})
				return
			}
		} else {
			var ancient models.Ancient
			ay.Db.Select("id").First(&ancient, "id = ?", getForm.Id)
			if ancient.Id == 0 {
				ay.Json{}.Msg(c, 400, "古籍不存在", gin.H{})
				return
			}
		}
		ss := models.Collect{
			Uid:  user.Id,
			Type: getForm.Type,
			Cid:  getForm.Id,
			//CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			//CreatedAt: models.JsonTime{},
		}
		ay.Db.Create(&ss)
		if ss.Id != 0 {
			ay.Json{}.Msg(c, 200, "收藏成功", gin.H{})
		} else {
			ay.Json{}.Msg(c, 400, "收藏失败", gin.H{})
		}
	} else {
		ay.Db.Delete(&collect)
		if collect.Id != 0 {
			ay.Json{}.Msg(c, 200, "已从收藏中移除", gin.H{})
		} else {
			ay.Json{}.Msg(c, 400, "移除失败", gin.H{})
		}
	}

}
