/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type UserInviteConsumptionModel struct {
}

type UserInviteConsumption struct {
	BaseModel
	Pid       int64   `json:"pid"`
	Uid       int64   `json:"uid"`
	Oid       string  `json:"oid"`
	Amount    float64 `json:"amount"`
	OldAmount float64 `json:"old_amount"`
	Status    int     `json:"status"`
}

func (UserInviteConsumption) TableName() string {
	return "sm_user_invite_consumption"
}
