/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

import "aggregate-fortune-telling/ay"

type UserModel struct {
}

type User struct {
	BaseModel
	Phone        string  `json:"phone"`
	Avatar       string  `gorm:"column:avatar" json:"avatar"`
	AreaId       int     `gorm:"column:area_id" json:"area_id"`
	Gender       int     `gorm:"column:gender" json:"gender"`
	NickName     string  `gorm:"column:nickname" json:"nickname"`
	BirthDay     string  `gorm:"column:birthday" json:"birthday"`
	InviteAmount float64 `gorm:"column:invite_amount" json:"invite_amount"`
	Amount       float64 `gorm:"column:amount" json:"amount"`
	Type         int     `gorm:"column:type" json:"type"`
	MasterId     int64   `gorm:"column:master_id" json:"master_id"`
	Pid          int64   `gorm:"column:pid" json:"pid"`
	Aff          string  `gorm:"column:aff" json:"aff"`
}

func (User) TableName() string {
	return "sm_user"
}

func (con UserModel) GetPhone(phone string) (res User) {
	ay.Db.Where("phone = ?", phone).First(&res)
	return
}
