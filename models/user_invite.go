/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type UserInviteModel struct {
}

type UserInvite struct {
	BaseModel
	Pid int64 `json:"pid"`
	Uid int64 `json:"uid"`
}

func (UserInvite) TableName() string {
	return "sm_user_invite"
}
