/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package common

import (
	"math"
	"strings"
)

type HexapodModel struct {
}

type HexapodList struct {
	Id       int `gorm:"primaryKey"`
	Title    string
	Content  string
	Handount string `gorm:"column:handout"`
}

var HexapodName = [64]string{
	"乾", "坤", "屯", "蒙", "需", "讼", "师", "比",
	"小畜", "履", "泰", "否", "同人", "大有", "谦", "豫",
	"随", "蛊", "临", "观", "噬嗑", "贲", "剥", "复",
	"无妄", "大畜", "颐", "大过", "坎", "离", "咸", "恒", "遁",
	"大壮", "晋", "明夷", "家人", "睽", "蹇", "解", "损",
	"益", "夬", "姤", "萃", "升", "困", "井", "革",
	"鼎", "震", "艮", "渐", "归妹", "丰", "旅",
	"巽", "兑", "涣", "节", "中孚", "小过", "既济", "未济",
}

var HexapodWuName = [64]string{
	"坤", "复", "师", "临", "谦", "明夷", "升", "泰",
	"豫", "震", "解", "归妹", "小过", "丰", "恒", "大壮",
	"比", "屯", "坎", "节", "蹇", "既济", "井", "需",
	"萃", "随", "困", "兑", "咸", "大过", "革", "夬",
	"剥", "颐", "蒙", "损", "艮", "贲", "蛊", "大畜",
	"晋", "噬嗑", "未济", "睽", "旅", "离", "鼎", "大有",
	"观", "益", "涣", "中孚", "渐", "家人", "巽", "小畜",
	"否", "无妄", "讼", "履", "遁", "同人", "姤", "乾",
}

func (con HexapodModel) Hexapod(data string) (int, string) {
	num := 0
	arr := strings.Split(data, ",")

	for k, v := range arr {
		if v == "1" {
			num += int(math.Pow(2, float64(k)))
		}
	}

	return num, HexapodWuName[num]

}
