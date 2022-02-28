/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package models

type AreaModel struct {
}

type Area struct {
	Id        int     `json:"id"`
	Pid       int     `json:"-"`
	ShortName string  `json:"-"`
	Name      string  `json:"name"`
	MergeName string  `json:"-"`
	Level     string  `json:"-"`
	Pinyin    string  `json:"-"`
	Code      string  `json:"-"`
	Zip       string  `json:"-"`
	First     string  `json:"-"`
	Lng       float64 `json:"-"`
	Lat       float64 `json:"-"`
}

func (Area) TableName() string {
	return "sm_area"
}
