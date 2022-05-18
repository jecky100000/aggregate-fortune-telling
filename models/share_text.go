/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

import "gin/ay"

type ShareTextModel struct {
}

type ShareText struct {
	BaseModel
	Content string `json:"content"`
}

func (ShareText) TableName() string {
	return "sm_share_text"
}

func (ShareTextModel) GetAll() (res []ShareText) {
	ay.Db.Find(&res)
	return
}
