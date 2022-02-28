/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package models

type Consult struct {
	BaseModel
	Type int    `json:"-"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Sort int    `json:"-"`
}

func (Consult) TableName() string {
	return "sm_consult"
}
