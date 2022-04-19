/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type AncientClassModel struct {
}

type AncientClass struct {
	//gorm.Model
	BaseModel
	Aid     int64  `json:"aid"`
	Name    string `json:"name"`
	Type    int    `json:"type"`
	Link    string `json:"link"`
	Content string `json:"content"`
	Sort    int    `json:"sort"`
}

func (AncientClass) TableName() string {
	return "sm_ancient_class"
}
