/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package models

type HaulAmountModel struct {
}

type HaulAmount struct {
	Id     int64   `gorm:"primaryKey" json:"id"`
	Amount float64 `gorm:"column:amount" json:"amount"`
	Sort   string  `gorm:"column:sort" json:"sort"`
}

func (HaulAmount) TableName() string {
	return "sm_haul_amount"
}
