/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package task

import (
	"gin/ay"
	"gin/models"
	"log"
)

// AskAmount 提问消费退回
func AskAmount() {
	log.Println("开始计算提问48小时退回")
	var order []models.Order
	ay.Db.Where("status = 0 AND type = 3 AND now() >SUBDATE(created_at,interval -2 day)").Find(&order)
	for _, v := range order {
		var user models.User
		ay.Db.First(&user, v.Uid)
		user.Amount += v.Amount
		tx := ay.Db.Begin()
		if err := tx.Save(&user).Error; err != nil {
			tx.Rollback()
			log.Println("提问金额退款失败")
		}
		if err := tx.Model(models.Order{}).Where("id = ?", v.Id).UpdateColumn("status", 3).Error; err != nil {
			tx.Rollback()
			log.Println("提问金额退款失败")
		}
		tx.Commit()
	}
}
