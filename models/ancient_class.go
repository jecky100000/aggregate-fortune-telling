/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package models

type AncientClassModel struct {
}

type AncientClass struct {
	//gorm.Model
	BaseModel
	Aid  int    `json:"-"`
	Name string `json:"name"`
	Type string `json:"type"`
	Link string `json:"link"`
	Sort string `json:"-"`
}

func (AncientClass) TableName() string {
	return "sm_ancient_class"
}
