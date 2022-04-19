/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

import "gin/ay"

type ConsultModel struct {
}

type Consult struct {
	BaseModel
	Type int    `json:"type"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Sort int    `json:"sort"`
}

func (Consult) TableName() string {
	return "sm_consult"
}

func (con ConsultModel) GetType(val int) (res []Consult) {
	ay.Db.Where("type = ?", val).Order("sort asc").Find(&res)
	return
}
