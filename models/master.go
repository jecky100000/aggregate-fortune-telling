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

type MasterModel struct {
}

type Master struct {
	BaseModel
	Type        string  `json:"type"`
	Sign        string  `json:"sign"`
	Label       string  `json:"label"`
	Introduce   string  `json:"introduce"`
	Years       int     `json:"years"`
	Online      int     `json:"online"`
	Rate        float64 `json:"rate"`
	AskNum      int     `json:"ask_num"`
	BackImage   string  `json:"back_image"`
	IsRecommend int     `json:"is_recommend"`
}

func (Master) TableName() string {
	return "sm_master"
}

func (con MasterModel) IsMaser(id int64) (bool, User, Master) {
	var user User

	var master Master

	ay.Db.First(&user, "id = ?", id)
	if user.MasterId == 0 {
		return false, user, master
	}
	ay.Db.Where("id = ?", user.MasterId).First(&master)

	if master.Id == 0 {
		return false, user, master
	}

	return true, user, master
}
