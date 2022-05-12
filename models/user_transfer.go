/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type UserTransferModel struct {
}

type UserTransfer struct {
	BaseModel
	Uid             int64   `json:"uid"`
	Amount          float64 `json:"amount"`
	OldInviteAmount float64 `json:"old_invite_amount"`
	NowInviteAmount float64 `json:"now_invite_amount"`
	OldAmount       float64 `json:"old_amount"`
	NowAmount       float64 `json:"now_amount"`
}

func (UserTransfer) TableName() string {
	return "sm_user_transfer"
}
