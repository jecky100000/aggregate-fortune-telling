/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

import (
	"fmt"
	"gin/ay"
	"strconv"
)

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

func (con UserInviteConsumptionModel) Set(uid, pid int64, amount float64, oid string) {
	config := ConfigModel{}.GetId(1)
	// 获取上级
	var pUser User
	ay.Db.First(&pUser, "id = ?", pid)
	if pUser.Id != 0 {
		inviteAmount, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", config.InviteRate*amount), 64)

		// 消费记录
		ay.Db.Create(&UserInviteConsumption{
			Pid:       pUser.Id,
			Uid:       uid,
			Amount:    inviteAmount,
			OldAmount: amount,
			Status:    0,
			Oid:       oid,
		})
	}
}
