/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package models

type BaiKeModel struct {
}

type BaiKe struct {
	BaseModel
	Type    int    `json:"-"`
	Title   string `json:"title"`
	Cover   string `json:"cover"`
	Content string `json:"content"`
	//Cid   string `json:"-"`
	VType int `json:"-"`
	//Vcid  string `json:"-"`
	View int64 `json:"view"`
}

func (BaiKe) TableName() string {
	return "sm_baike"
}
