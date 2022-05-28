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

type AncientController struct {
}

type ancientListForm struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Key      string `form:"key"`
	Status   string `form:"status"`
	Type     int    `form:"type"`
}

// List 列表
func (con AncientController) List(c *gin.Context) {
	var data ancientListForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	type returnList struct {
		models.Ancient
		TypeName string `json:"typeName"`
	}

	var list []returnList

	var count int64
	res := ay.Db.Table("sm_ancient").
		Select("sm_ancient.*,sm_ancient_type.name as type_name").
		Joins("left join sm_ancient_type on sm_ancient.vcid=sm_ancient_type.id")

	if data.Key != "" {
		res.Where("sm_ancient.title like ?", "%"+data.Key+"%")
	}

	if data.Type != 0 {
		res.Where("sm_ancient.type = ?", data.Type)
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
func (con AncientController) Detail(c *gin.Context) {
	var data orderDetailForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var user models.Ancient

	ay.Db.First(&user, data.Id)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"info": user,
	})
}

type ancientOptionForm struct {
	Id    int    `form:"id"`
	Title string `form:"title"`
	Cover string `form:"cover"`
	Cid   int    `form:"cid"`
	Vcid  int    `form:"vcid"`
}

// Option 添加 编辑
func (con AncientController) Option(c *gin.Context) {
	var data ancientOptionForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var res models.Ancient
	ay.Db.First(&res, data.Id)

	if data.Id != 0 {

		res.Title = data.Title
		res.Cover = data.Cover
		res.Cid = data.Cid
		res.Vcid = data.Vcid

		if err := ay.Db.Save(&res); err != nil {
			ay.Json{}.Msg(c, 400, "修改失败", gin.H{})
		} else {
			ay.Json{}.Msg(c, 200, "修改成功", gin.H{})
		}
	} else {
		if err := ay.Db.Create(&models.Ancient{
			Title: data.Title,
			Cover: data.Cover,
			Cid:   data.Cid,
			Vcid:  data.Vcid,
			Type:  2,
		}); err != nil {
			ay.Json{}.Msg(c, 400, "创建失败", gin.H{})
		} else {
			ay.Json{}.Msg(c, 200, "创建成功", gin.H{})
		}

	}

}

func (con AncientController) Delete(c *gin.Context) {
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
		var order models.Ancient
		ay.Db.Delete(&order, v)
	}

	ay.Json{}.Msg(c, 200, "删除成功", gin.H{})
}

func (con AncientController) Upload(c *gin.Context) {

	code, msg := Upload(c, "notice")

	if code != 200 {
		ay.Json{}.Msg(c, 400, msg, gin.H{})
	} else {
		ay.Json{}.Msg(c, 200, msg, gin.H{})
	}
}
