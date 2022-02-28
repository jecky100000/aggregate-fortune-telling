/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package models

type HomeModel struct {
}

type Home struct {
	Id    int64  `gorm:"primaryKey" json:"-"`
	Image string `gorm:"column:image" json:"image"`
	Link  string `gorm:"column:link" json:"link"`
}

func (Home) TableName() string {
	return "sm_home"
}
