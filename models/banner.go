/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package models

type Banner struct {
	BaseModel
	Image string `json:"image"`
	Url   string `json:"url"`
	Sort  int    `json:"-"`
}

func (Banner) TableName() string {
	return "sm_banner"
}
