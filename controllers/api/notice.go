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
	"strconv"
)

type NoticeController struct {
}

type NoticeSearch struct {
	Id        int           `json:"id"`
	Type      int           `json:"type"`
	Title     string        `json:"title"`
	Cover     string        `json:"cover"`
	CreatedAt models.MyTime `json:"created_at"`
	Collect   int           `json:"collect"`
	View      int64         `json:"view"`
	Time      string        `json:"time"`
}

type GetNoticeSearchForm struct {
	Type  int    `form:"type" binding:"required" label:"类型"`
	Title string `form:"title"`
	Page  int    `form:"page" binding:"required" label:"分页"`
}

// Search 搜索
func (con NoticeController) Search(c *gin.Context) {

	var getForm GetNoticeSearchForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	uid := GetToken(Token)

	sql := ""
	page := getForm.Page

	if page != 0 {
		page = (page - 1) * 10
	}

	page1 := strconv.Itoa(page)
	title := getForm.Title
	field := "id,type,title,cover,created_at,view"

	if getForm.Type == 3 {
		if getForm.Title == "" {
			sql = "(SELECT " + field + " FROM sm_ancient LIMIT " + page1 + ",10) UNION (SELECT " + field + " FROM sm_baike LIMIT " + page1 + ",10)"
		} else {
			sql = "(SELECT " + field + " FROM sm_ancient WHERE title LIKE '%" + title + "%' LIMIT " + page1 + ",10) UNION (SELECT " + field + " FROM sm_baike WHERE title LIKE '%" + title + "%' LIMIT " + page1 + ",10)"
		}
	} else if getForm.Type == 1 {
		if getForm.Title == "" {
			sql = "SELECT " + field + " FROM sm_baike ORDER BY id DESC LIMIT " + strconv.Itoa((getForm.Page-1)*20) + ",20"
		} else {
			sql = "SELECT " + field + " FROM sm_baike WHERE title LIKE '%" + title + "%'  ORDER BY id DESC LIMIT " + strconv.Itoa((getForm.Page-1)*20) + ",20"
		}
	} else {
		if title == "" {
			sql = "SELECT " + field + " FROM sm_ancient LIMIT " + strconv.Itoa((getForm.Page-1)*20) + ",20"
		} else {
			sql = "SELECT " + field + " FROM sm_ancient WHERE title LIKE '%" + title + "%' LIMIT " + strconv.Itoa((getForm.Page-1)*20) + ",20"
		}
	}

	var search []NoticeSearch

	ay.Db.Raw(sql).Scan(&search)

	for k, v := range search {
		var collect models.Collect
		ay.Db.First(&collect, "type = ? and uid = ? and cid = ?", v.Type+1, uid, v.Id)
		if collect.Id == 0 {
			search[k].Collect = 0
		} else {
			search[k].Collect = 1
		}
		search[k].Time = ay.LastTime(int(v.CreatedAt.Unix()))
		search[k].Cover = ay.Yaml.GetString("domain") + v.Cover
	}

	if search == nil {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": []string{},
		})
	} else {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": search,
		})
	}

}

type GetAncientDetailForm struct {
	Aid string `form:"aid" binding:"required" label:"古籍id"`
	//Page int    `form:"page" binding:"required"`
}

// Detail 古籍详情
func (con NoticeController) Detail(c *gin.Context) {

	var getForm GetAncientDetailForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var res models.Ancient
	ay.Db.First(&res, "id = ?", getForm.Aid)
	if res.Id == 0 {
		ay.Json{}.Msg(c, 400, "古籍不存在", gin.H{})
		return
	}
	res.View = res.View + 1
	ay.Db.Save(&res)

	var ancient []models.AncientClass

	ay.Db.Where("aid = ?", getForm.Aid).Order("sort asc").Find(&ancient)

	for k, v := range ancient {
		ancient[k].Link = ay.Yaml.GetString("domain") + v.Link
	}

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list": ancient,
	})
}

type GetBaiKeForm struct {
	Id int `form:"id" binding:"required"`
}

// BaiKe 百科详情
func (con NoticeController) BaiKe(c *gin.Context) {

	var getForm GetBaiKeForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var baike models.BaiKe

	ay.Db.First(&baike, "id = ?", getForm.Id)

	if baike.Id == 0 {
		ay.Json{}.Msg(c, 400, "数据有误", gin.H{})
	} else {
		baike.View = baike.View + 1
		ay.Db.Save(&baike)
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"info": baike,
		})
	}

}
