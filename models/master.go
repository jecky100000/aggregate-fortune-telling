/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

import (
	"aggregate-fortune-telling/ay"
	"strings"
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

type UserMaster struct {
	Master
	Avatar   string   `json:"avatar"`
	Nickname string   `json:"nickname"`
	Phone    string   `json:"phone"`
	TypeName []string `json:"type_name"`
}

func (con MasterModel) GetMaster(userId int64) (bool, UserMaster) {

	var row UserMaster

	ay.Db.Table("sm_user").
		Select("sm_master.*,sm_user.avatar,sm_user.nickname,sm_user.phone,sm_user.id").
		Joins("left join sm_master on sm_user.master_id=sm_master.id").
		Where("sm_user.id = ?", userId).
		First(&row)

	if row.Id == 0 {
		return false, UserMaster{}
	} else {

		for _, v := range strings.Split(row.Type, ",") {
			var masterType MasterType
			ay.Db.First(&masterType, "id = ?", v)
			if masterType.Name != "" {
				row.TypeName = append(row.TypeName, masterType.Name)
			}

		}
		row.Avatar = ay.Yaml.GetString("domain") + row.Avatar
		row.BackImage = ay.Yaml.GetString("domain") + row.BackImage

		return true, row
	}

}

func (con MasterModel) GetMasterPage(page int, isRecommend int, isType int) []UserMaster {

	var row []UserMaster

	res := ay.Db.Table("sm_user").
		Select("sm_master.*,sm_user.avatar,sm_user.nickname,sm_user.phone,sm_user.id").
		Joins("left join sm_master on sm_user.master_id=sm_master.id")

	if isRecommend == 1 {
		res.Where("is_recommend = 1")
	} else {
		res.Where("FIND_IN_SET(?,sm_master.type) and sm_user.type = 1", isType)
	}

	res.Limit(10).
		Offset(page * 10).
		Order("sm_user.id desc").
		Find(&row)

	for k, v1 := range row {
		for _, v := range strings.Split(v1.Type, ",") {
			var masterType MasterType
			ay.Db.First(&masterType, "id = ?", v)
			if masterType.Name != "" {
				row[k].TypeName = append(row[k].TypeName, masterType.Name)
			}
		}
		row[k].Avatar = ay.Yaml.GetString("domain") + v1.Avatar
		row[k].BackImage = ay.Yaml.GetString("domain") + v1.BackImage
	}

	return row

}
