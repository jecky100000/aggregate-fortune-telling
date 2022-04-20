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

type AdvController struct {
}

// List 列表
func (con AdvController) List(c *gin.Context) {
	var data noticeListForm
	if err := c.ShouldBind(&data); err != nil {
		Json.Msg(400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		Json.Msg(401, "请登入", gin.H{})
		return
	}

	var list []models.Adv

	var count int64
	res := ay.Db.Table("sm_adv")

	row := res

	row.Count(&count)

	res.Order("created_at desc").
		Limit(data.PageSize).
		Offset((data.Page - 1) * data.PageSize).
		Find(&list)

	Json.Msg(200, "success", gin.H{
		"list":  list,
		"total": count,
	})
}

//type orderDetailForm struct {
//	Id int `form:"id"`
//}

// Detail 用户详情
func (con AdvController) Detail(c *gin.Context) {
	var data orderDetailForm
	if err := c.ShouldBind(&data); err != nil {
		Json.Msg(400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		Json.Msg(401, "请登入", gin.H{})
		return
	}

	var user models.Adv

	ay.Db.First(&user, data.Id)

	Json.Msg(200, "success", gin.H{
		"info": user,
	})
}

type advOptionForm struct {
	Id    int    `form:"id"`
	Link  string `form:"link"`
	Image string `form:"image"`
	Sort  int    `form:"sort"`
	Type  int    `form:"type"`
}

// Option 添加 编辑
func (con AdvController) Option(c *gin.Context) {
	var data advOptionForm
	if err := c.ShouldBind(&data); err != nil {
		Json.Msg(400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		Json.Msg(401, "请登入", gin.H{})
		return
	}

	var res models.Adv
	ay.Db.First(&res, data.Id)

	if data.Id != 0 {

		res.Link = data.Link
		res.Sort = data.Sort
		res.Image = data.Image
		res.Type = data.Type

		ay.Db.Save(&res)
		Json.Msg(200, "修改成功", gin.H{})
	} else {
		ay.Db.Create(&models.Adv{
			Sort:  data.Sort,
			Image: data.Image,
			Link:  data.Link,
			Type:  data.Type,
		})
		Json.Msg(200, "创建成功", gin.H{})

	}

}

func (con AdvController) Delete(c *gin.Context) {
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
		var order models.Adv
		ay.Db.Delete(&order, v)
	}

	Json.Msg(200, "删除成功", gin.H{})
}

func (con AdvController) Upload(c *gin.Context) {
	if Auth() == false {
		Json.Msg(401, "请登入", gin.H{})
		return
	}
	code, msg := Upload(c, "adv")

	if code != 200 {
		Json.Msg(400, msg, gin.H{})
	} else {
		Json.Msg(200, msg, gin.H{})
	}
}
