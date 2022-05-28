/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package common

import (
	"aggregate-fortune-telling/ay"
	"aggregate-fortune-telling/models"
)

type NaYinModel struct {
}

var (
	gzu  = []string{"甲子", "丙寅", "戊辰", "庚午", "壬申", "甲戌", "丙子", "戊寅", "庚辰", "壬午", "甲申", "丙戌", "戊子", "庚寅", "壬辰", "甲午", "丙申", "戊戌", "庚子", "壬寅", "甲辰", "丙午", "戊申", "庚戌", "壬子", "甲寅", "丙辰", "戊午", "庚申", "壬戌"}
	zzu  = []string{"乙丑", "丁卯", "己巳", "辛未", "癸酉", "乙亥", "丁丑", "己卯", "辛巳", "癸未", "乙酉", "丁亥", "己丑", "辛卯", "癸巳", "乙未", "丁酉", "己亥", "辛丑", "癸卯", "乙巳", "丁未", "已酉", "辛亥", "癸丑", "乙卯", "丁巳", "己未", "辛酉", "癸亥"}
	nyzu = []string{"海中金", "炉中火", "大林木", "路旁土", "剑锋金", "山头火", "涧下水", "城头土", "白腊金", "杨柳木 ", "泉中水", "屋上土", "霹雳火", "松柏木", "长流水", "砂石金", "山下火", "平地木", "壁上土", "金薄金", "覆灯火", "天河水", "大驿土", "钗环金", "桑柘木", "大溪水", "沙中土", "天上火", "石榴木", "大海水"}
)

func (con NaYinModel) Get(gz string) string {

	var nayin models.JiaZi
	ay.Db.First(&nayin, "jiazi = ?", gz)

	// log.Println(nayin)
	return nayin.NaYin
	//z1 := 0
	//for k, v := range gzu {
	//	if v == gz {
	//		z1 = k
	//	}
	//}
	//if z1 == 0 {
	//	z2 := 0
	//	for k, v := range zzu {
	//		if v == gz {
	//			z2 = k
	//		}
	//	}
	//	return nyzu[z2]
	//} else {
	//	return nyzu[z1]
	//}
}
