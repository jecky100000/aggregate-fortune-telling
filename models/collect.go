/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package models

type CollectModel struct {
}

type Collect struct {
	BaseModel
	Uid  int64 `json:"uid"`
	Type int   `json:"type"`
	Cid  int64 `json:"cid"`
}

func (Collect) TableName() string {
	return "sm_collect"
}
