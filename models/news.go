/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

import (
	"gin/ay"
)

type NewsModel struct {
}

type NewsNotice struct {
	BaseModel
	// Class       int    `gorm:"column:class" json:"-"`
	Type        int64  `gorm:"column:type" json:"type"`
	Title       string `gorm:"column:title" json:"title"`
	Keywords    string `gorm:"column:keywords" json:"keywords"`
	Description string `gorm:"column:description" json:"description"`
	Cover       string `gorm:"column:cover" json:"cover"`
	Content     string `gorm:"column:content" json:"content"`
	View        int    `gorm:"column:view" json:"view"`
	Status      int    `gorm:"column:status" json:"status"`
}

func (NewsNotice) TableName() string {
	return "sm_notice"
}

type returnNotice struct {
	NewsNotice
	Time string `json:"time"`
}

func (con NewsModel) GetList(class int) (res []returnNotice) {
	var noticeType NewsType
	ay.Db.First(&noticeType, "class = ?", class)

	ay.Db.Where("status = ? and type = ?", noticeType.Id, class).Select("id,title,keywords,description,cover,view,created_time").Order("RAND()").Limit(10).Find(&res)
	return
}

func (con NewsModel) GetDetail(id int) (res NewsNotice) {
	ay.Db.Where("id", id).Find(&res)
	con.UpdateTotal(&res)
	return
}

func (con NewsModel) UpdateTotal(res *NewsNotice) bool {
	res.View = res.View + 1
	if err := ay.Db.Table("sm_notice").Save(res).Error; err != nil {
		return false
	} else {
		return true
	}
}
