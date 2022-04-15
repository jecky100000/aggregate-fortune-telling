/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type UserCommentModel struct {
}

type UserComment struct {
	BaseModel
	Rate     float64 `json:"rate"`
	Content  string  `json:"content"`
	Uid      int64   `json:"uid"`
	MasterId int64   `json:"master_id"`
}

func (UserComment) TableName() string {
	return "sm_user_comment"
}
