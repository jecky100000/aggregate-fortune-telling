/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

import "aggregate-fortune-telling/ay"

type UserOpenidModel struct {
}

type UserOpenid struct {
	BaseModel
	Openid string `json:"openid"`
	Uid    int64  `json:"uid"`
	Appid  int64  `json:"appid"`
}

func (UserOpenid) TableName() string {
	return "sm_user_openid"
}

func (con UserOpenidModel) Get(appid int64, openid string) (res UserOpenid) {
	ay.Db.Where("openid = ? AND appid = ?", openid, appid).First(&res)
	return
}
