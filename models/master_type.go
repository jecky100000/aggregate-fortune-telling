/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type MasterTypeModel struct {
}

type MasterType struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (MasterType) TableName() string {
	return "sm_master_type"
}
