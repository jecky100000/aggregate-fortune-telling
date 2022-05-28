/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package task

import (
	"aggregate-fortune-telling/ay"
	"aggregate-fortune-telling/models"
	"log"
)

func InviteAmount() {
	log.Println("开始计算邀请用户收益")
	var consumption []models.UserInviteConsumption
	ay.Db.Where("status = 0 AND now() >SUBDATE(created_at,interval -3 day)").Find(&consumption)
	for _, v := range consumption {
		var pUser models.User
		ay.Db.First(&pUser, v.Pid)
		pUser.InviteAmount += v.Amount
		if err := ay.Db.Save(&pUser).Error; err == nil {
			ay.Db.Model(models.UserInviteConsumption{}).Where("id = ?", v.Id).UpdateColumn("status", 1)
			//ay.Db.Create(&models.Order{
			//	Uid:    pUser.Id,
			//	Des:    strconv.FormatInt(v.Id, 10),
			//	Status: 1,
			//	Amount: v.Amount,
			//	Op:     1,
			//	Type:   4,
			//	Line:   strconv.FormatInt(v.Uid, 10),
			//})
		}
	}
}
