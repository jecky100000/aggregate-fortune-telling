/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type NewsTypeModel struct {
}

type NewsType struct {
	Id    uint   `json:"id"`
	Pid   uint   `json:"-"`
	Name  string `json:"name"`
	Class int    `json:"-"`
}

func (NewsType) TableName() string {
	return "sm_notice_type"
}
