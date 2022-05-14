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

type HaulCasesController struct {
}

type haulCasesListForm struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Key      string `form:"key"`
	Type     string `form:"type"`
}

// List 列表
func (con HaulCasesController) List(c *gin.Context) {
	var data haulCasesListForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var list []models.HaulCases

	var count int64
	res := ay.Db.Table("sm_haul_cases")

	if data.Key != "" {
		res.Where("sm_haul_cases.name like ?", "%"+data.Key+"%")
	}

	if data.Type != "" {
		res.Where("sm_haul_cases.type = ?", data.Type)
	}

	row := res

	row.Count(&count)

	res.Order("sort asc").
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
func (con HaulCasesController) Detail(c *gin.Context) {
	var data orderDetailForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var user models.HaulCases

	ay.Db.First(&user, data.Id)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"info": user,
	})
}

type haulCasesOptionForm struct {
	Id    int    `form:"id"`
	Name  string `form:"name"`
	Link  string `form:"link"`
	Cover string `form:"cover"`
	Sort  int    `form:"sort"`
	Type  int    `form:"type"`
}

// Option 添加 编辑
func (con HaulCasesController) Option(c *gin.Context) {
	var data haulCasesOptionForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var res models.HaulCases
	ay.Db.First(&res, data.Id)

	if data.Id != 0 {

		res.Cover = data.Cover
		res.Sort = data.Sort
		res.Name = data.Name
		res.Type = data.Type
		res.Link = data.Link

		if err := ay.Db.Save(&res).Error; err != nil {
			ay.Json{}.Msg(c, 400, "修改失败", gin.H{})
		} else {
			ay.Json{}.Msg(c, 200, "修改成功", gin.H{})
		}
	} else {
		ay.Db.Create(&models.HaulCases{
			Name:  data.Name,
			Cover: data.Cover,
			Type:  data.Type,
			Link:  data.Link,
			Sort:  data.Sort,
		})
		ay.Json{}.Msg(c, 200, "创建成功", gin.H{})

	}

}

func (con HaulCasesController) Delete(c *gin.Context) {
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
		var order models.HaulCases
		ay.Db.Delete(&order, v)
	}

	ay.Json{}.Msg(c, 200, "删除成功", gin.H{})
}

func (con HaulCasesController) Upload(c *gin.Context) {

	code, msg := Upload(c, "haul_cases")

	if code != 200 {
		ay.Json{}.Msg(c, 400, msg, gin.H{})
	} else {
		ay.Json{}.Msg(c, 200, msg, gin.H{})
	}
}
