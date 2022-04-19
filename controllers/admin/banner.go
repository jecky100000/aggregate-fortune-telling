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

type BannerController struct {
}

// List 列表
func (con BannerController) List(c *gin.Context) {
	var data noticeListForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, "400", ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, "401", "请登入", gin.H{})
		return
	}

	var list []models.Banner

	var count int64
	res := ay.Db.Table("sm_banner")

	row := res

	res.Order("created_at desc").
		Limit(data.PageSize).
		Offset((data.Page - 1) * data.PageSize).
		Find(&list)

	row.Count(&count)

	ay.Json{}.Msg(c, "200", "success", gin.H{
		"list":  list,
		"total": count,
	})
}

//type orderDetailForm struct {
//	Id int `form:"id"`
//}

// Detail 用户详情
func (con BannerController) Detail(c *gin.Context) {
	var data orderDetailForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, "400", ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, "401", "请登入", gin.H{})
		return
	}

	var user models.Banner

	ay.Db.First(&user, data.Id)

	ay.Json{}.Msg(c, "200", "success", gin.H{
		"info": user,
	})
}

type bannerOptionForm struct {
	Id    int    `form:"id"`
	Url   string `form:"url"`
	Image string `form:"image"`
	Sort  int    `form:"sort"`
}

// Option 添加 编辑
func (con BannerController) Option(c *gin.Context) {
	var data bannerOptionForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, "400", ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, "401", "请登入", gin.H{})
		return
	}

	var res models.Banner
	ay.Db.First(&res, data.Id)

	if data.Id != 0 {

		res.Url = data.Url
		res.Sort = data.Sort
		res.Image = data.Image

		ay.Db.Save(&res)
		ay.Json{}.Msg(c, "200", "修改成功", gin.H{})
	} else {
		ay.Db.Create(&models.Banner{
			Sort:  data.Sort,
			Image: data.Image,
			Url:   data.Url,
		})
		ay.Json{}.Msg(c, "200", "创建成功", gin.H{})

	}

}

func (con BannerController) Delete(c *gin.Context) {
	var data orderDeleteForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, "400", ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, "401", "请登入", gin.H{})
		return
	}

	idArr := strings.Split(data.Id, ",")

	for _, v := range idArr {
		var order models.Banner
		ay.Db.Delete(&order, v)
	}

	ay.Json{}.Msg(c, "200", "删除成功", gin.H{})
}

func (con BannerController) Upload(c *gin.Context) {

	code, msg := Upload(c, "banner")

	if code != 200 {
		ay.Json{}.Msg(c, "400", msg, gin.H{})
	} else {
		ay.Json{}.Msg(c, "200", msg, gin.H{})
	}
}
