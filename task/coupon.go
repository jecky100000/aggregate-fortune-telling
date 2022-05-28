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

// CouponAmount 优惠卷设置过期
func CouponAmount() {
	log.Println("开始优惠卷设置过期")
	var res []models.Coupon
	ay.Db.Where("status = 0 AND now() > effective_at").Find(&res)
	for _, v := range res {
		ay.Db.Model(models.Coupon{}).Where("id = ?", v.Id).UpdateColumn("status", 3)
	}
}
