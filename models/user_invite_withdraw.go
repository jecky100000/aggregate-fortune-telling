/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type UserInviteWithdrawModel struct {
}

type UserInviteWithdraw struct {
	BaseModel
	Uid       int64   `json:"uid"`
	Amount    float64 `json:"amount"`
	OldAmount float64 `json:"old_amount"`
	Status    int     `json:"status"`
}

func (UserInviteWithdraw) TableName() string {
	return "sm_user_invite_withdraw"
}
