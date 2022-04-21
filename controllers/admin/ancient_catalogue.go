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
	"log"
	"path"
	"strings"
)

type AncientCatalogueController struct {
}

type ancientCatalogueListForm struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Key      string `form:"key"`
	Id       int    `form:"id"`
}

// List 列表
func (con AncientCatalogueController) List(c *gin.Context) {
	var data ancientCatalogueListForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}
	log.Println(data.Id)

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var list []models.AncientClass

	var count int64
	ay.Db.Order("sort asc").
		Where("aid = ?", data.Id).
		Limit(data.PageSize).
		Offset((data.Page - 1) * data.PageSize).
		Find(&list)

	ay.Db.Table("sm_ancient_class").
		Where("aid = ?", data.Id).
		Count(&count)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list":  list,
		"total": count,
	})
}

//type orderDetailForm struct {
//	Id int `form:"id"`
//}

// Detail 用户详情
func (con AncientCatalogueController) Detail(c *gin.Context) {
	var data orderDetailForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var res models.AncientClass

	ay.Db.First(&res, data.Id)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"info": res,
	})
}

type ancientCatalogueOptionForm struct {
	Id      int    `form:"id"`
	Name    string `form:"name"`
	Content string `form:"content"`
	Link    string `form:"link"`
	Sort    int    `form:"sort"`
	Aid     int64  `form:"aid"`
}

// Option 添加 编辑
func (con AncientCatalogueController) Option(c *gin.Context) {
	var data ancientCatalogueOptionForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var res models.AncientClass
	ay.Db.First(&res, data.Id)

	vType := 0
	ext := path.Ext(data.Link)
	if ext == "doc" || ext == "docx" {
		vType = 3
	} else if ext == "pdf" {
		vType = 2
	} else {
		vType = 1
	}

	if data.Id != 0 {
		res.Name = data.Name
		res.Link = data.Link
		res.Content = data.Content
		res.Sort = data.Sort
		res.Type = vType

		ay.Db.Save(&res)
		ay.Json{}.Msg(c, 200, "修改成功", gin.H{})
	} else {

		ay.Db.Create(&models.AncientClass{
			Name:    data.Name,
			Link:    data.Link,
			Content: data.Content,
			Sort:    data.Sort,
			Aid:     data.Aid,
			Type:    vType,
		})
		ay.Json{}.Msg(c, 200, "创建成功", gin.H{})

	}

}

func (con AncientCatalogueController) Delete(c *gin.Context) {
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
		var res models.AncientClass
		ay.Db.Delete(&res, v)
	}

	ay.Json{}.Msg(c, 200, "删除成功", gin.H{})
}

func (con AncientCatalogueController) Upload(c *gin.Context) {

	code, msg := Upload(c, "ancient_catalogue")

	if code != 200 {
		ay.Json{}.Msg(c, 400, msg, gin.H{})
	} else {
		ay.Json{}.Msg(c, 200, msg, gin.H{})
	}
}
