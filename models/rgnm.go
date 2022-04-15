/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type RgnmModel struct {
}

type Rgnm struct {
	Id   int64  `gorm:"primaryKey" json:"id"`
	Rgz  string `gorm:"column:rgz" json:"rgz"`
	Rgxx string `gorm:"column:rgxx" json:"rgxx"`

	Rgcz  string `gorm:"column:rgcz" json:"rgcz"`
	Rgzfx string `gorm:"column:rgzfx" json:"rgzfx"`
	Xgfx  string `gorm:"column:xgfx" json:"xgfx"`
	Aqfx  string `gorm:"column:aqfx" json:"aqfx"`
	Syfx  string `gorm:"column:syfx" json:"syfx"`
	Cyfx  string `gorm:"column:cyfx" json:"cyfx"`
	Jkfx  string `gorm:"column:jkfx" json:"jkfx"`
}

func (Rgnm) TableName() string {
	return "sm_rgnm"
}
