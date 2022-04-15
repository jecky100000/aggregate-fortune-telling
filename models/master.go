/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type MasterModel struct {
}

type Master struct {
	BaseModel
	Type      string  `json:"type"`
	Name      string  `json:"name"`
	Sign      string  `json:"sign"`
	Label     string  `json:"label"`
	Introduce string  `json:"introduce"`
	Avatar    string  `json:"avatar"`
	Years     int     `json:"years"`
	Online    int     `json:"online"`
	Rate      float64 `json:"rate"`
	AskNum    int     `json:"ask_num"`
	Uid       int64   `json:"-"`
	Fans      int64   `json:"fans"`
	Reply     int64   `json:"reply"`
}

func (Master) TableName() string {
	return "sm_master"
}
