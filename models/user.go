/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type UserModel struct {
}

type User struct {
	BaseModel
	Phone    string  `json:"phone"`
	Appid    int     `json:"appid"`
	Openid   string  `json:"openid"`
	Avatar   string  `gorm:"column:avatar" json:"avatar"`
	AreaId   int     `gorm:"column:area_id" json:"area_id"`
	Gender   int     `gorm:"column:gender" json:"gender"`
	NickName string  `gorm:"column:nickname" json:"nickname"`
	BirthDay string  `gorm:"column:birthday" json:"birthday"`
	Amount   float64 `gorm:"column:amount" json:"amount"`
	Type     int     `gorm:"column:type" json:"type"`
	MasterId int64   `gorm:"column:master_id" json:"master_id"`
}

func (User) TableName() string {
	return "sm_user"
}
