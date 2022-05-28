/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

import (
	"aggregate-fortune-telling/ay"
	"log"
	"strings"
)

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

func (con AreaModel) GetP(id int64) string {
	if id == 0 {
		return ""
	}
	var res Area
	ay.Db.First(&res, id)

	log.Println(res.MergeName)
	arr := strings.Split(res.MergeName, ",")

	return arr[1] + arr[2] + arr[3]
}
