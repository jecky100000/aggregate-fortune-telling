/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type AncientModel struct {
}

type Ancient struct {
	BaseModel
	Type   int    `json:"type"`
	Title  string `json:"title"`
	Cover  string `json:"cover"`
	Author string `json:"author"`
	Cid    string `json:"cid"`
	Vcid   string `gorm:"column:vcid" json:"vcid"`
	View   int64  `gorm:"column:view" json:"view"`
}

func (Ancient) TableName() string {
	return "sm_ancient"
}
