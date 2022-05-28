/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

import (
	"aggregate-fortune-telling/ay"
)

type DreamModel struct {
}

type Dream struct {
	Id      int    `gorm:"primaryKey" json:"id"`
	Title   string `gorm:"column:title" json:"title"`
	Message string `gorm:"column:message" json:"message"`
	Biglx   string `gorm:"column:biglx" json:"-"`
	Smalllx string `gorm:"column:smalllx" json:"-"`
	Zm      string `gorm:"column:zm" json:"-"`
}

type DreamList struct {
	Id      int    `gorm:"primaryKey" json:"id"`
	Title   string `gorm:"column:title" json:"title"`
	Message string `gorm:"column:message" json:"message"`
}

func (Dream) TableName() string {
	return "sm_dream"
}

func (con DreamModel) GetList(title string, limit int) (res []DreamList) {
	if title == "0" {
		ay.Db.Table("sm_dream").Order("RAND()").Limit(limit).Select([]string{"id", "title", "message"}).Scan(&res)
	} else {
		ay.Db.Table("sm_dream").Where("title LIKE ?", "%"+title+"%").Order("RAND()").Limit(limit).Select([]string{"id", "title", "message"}).Scan(&res)
	}

	return
}

func (con DreamModel) GetDetail(id int) (res Dream) {
	ay.Db.Table("sm_dream").Where("id", id).Find(&res)
	return
}
