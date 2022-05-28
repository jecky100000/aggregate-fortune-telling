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

// AskAmount 提问消费退回
func AskAmount() {
	log.Println("开始计算提问48小时关闭")
	var order []models.Order
	ay.Db.Where("status = 0 AND type = 3 AND now() >SUBDATE(created_at,interval -2 day)").Find(&order)
	for _, v := range order {
		if err := ay.Db.Model(models.Order{}).Where("id = ?", v.Id).UpdateColumn("status", 1).Error; err != nil {
			log.Println("提问订单关闭失败")
		} else {
			log.Println("提问订单关闭成功")
		}
	}
}
