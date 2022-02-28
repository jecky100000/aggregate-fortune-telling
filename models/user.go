/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package models

type UserModel struct {
}

type User struct {
	BaseModel
	Phone    string
	Appid    int
	Openid   string
	Avatar   string  `gorm:"column:avatar"`
	AreaId   int     `gorm:"column:area_id"`
	Gender   int     `gorm:"column:gender"`
	NickName string  `gorm:"column:nickname"`
	BirthDay string  `gorm:"column:birthday"`
	Amount   float64 `gorm:"column:amount"`
	Type     int     `gorm:"column:type"`
	MasterId int64   `gorm:"column:master_id"`
}

func (User) TableName() string {
	return "sm_user"
}
