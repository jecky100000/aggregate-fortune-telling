/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package models

type TenGodsModel struct {
}

type TenGods struct {
	Id      int64  `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"column:name" json:"name"`
	Content string `gorm:"column:content" json:"content"`
}

func (TenGods) TableName() string {
	return "sm_ten_gods"
}
