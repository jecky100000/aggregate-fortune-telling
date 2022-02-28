/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package common

type XingModel struct {
}

// Cai 财星
func (con XingModel) Cai(bz []string) int {
	rg := bz[4]
	num := 0
	bzTg := []string{
		bz[0], bz[2], bz[4], bz[6],
	}
	bzDz := []string{
		bz[1], bz[3], bz[5], bz[7],
	}

	tg := map[string]interface{}{
		"甲": []string{"戌", "己"},
		"乙": []string{"戌", "己"},
		"丙": []string{"庚", "辛"},
		"丁": []string{"庚", "辛"},
		"戊": []string{"壬", "癸"},
		"己": []string{"壬", "癸"},
		"庚": []string{"甲", "乙"},
		"辛": []string{"甲", "乙"},
		"壬": []string{"丙", "丁"},
		"癸": []string{"丙", "丁"},
	}

	dz := map[string]interface{}{
		"甲": []string{"辰", "戌", "丑", "未"},
		"乙": []string{"辰", "戌", "丑", "未"},
		"丙": []string{"申", "酉"},
		"丁": []string{"申", "酉"},
		"戊": []string{"子", "亥"},
		"己": []string{"子", "亥"},
		"庚": []string{"寅", "卯"},
		"辛": []string{"寅", "卯"},
		"壬": []string{"巳", "午"},
		"癸": []string{"巳", "午"},
	}

	for _, v := range tg[rg].([]string) {
		for _, v1 := range bzTg {
			if v == v1 {
				num++
			}
		}
	}

	for _, v := range dz[rg].([]string) {
		for _, v1 := range bzDz {
			if v == v1 {
				num++
			}
		}
	}
	return num
}

// TaoHua 桃花星
func (con XingModel) TaoHua(bz []string) int {
	rz := bz[5]
	rg := bz[4]
	num := 0
	bzDz := []string{
		bz[1], bz[3], bz[5], bz[7],
	}
	tg := map[string]string{
		"甲": "子",
		"乙": "巳",
		"丙": "卯",
		"丁": "申",
		"戊": "卯",
		"己": "申",
		"庚": "午",
		"辛": "亥",
		"壬": "酉",
		"癸": "寅",
	}
	dz := map[string]string{
		"子": "酉",
		"丑": "午",
		"寅": "卯",
		"卯": "子",
		"辰": "酉",
		"巳": "午",
		"午": "卯",
		"未": "子",
		"申": "酉",
		"酉": "午",
		"戌": "卯",
		"亥": "子",
	}

	// 沐浴
	for _, v := range bzDz {
		if v == tg[rg] {
			num++
		}
	}

	// 咸
	for _, v := range bzDz {
		if v == dz[rz] {
			num++
		}
	}

	return num
}

// GuiRen 贵人星
func (con XingModel) GuiRen(bz []string) int {
	yg := bz[0]
	rg := bz[4]
	bzDz := []string{
		bz[1], bz[3], bz[5], bz[7],
	}
	num := 0

	dz := map[string]interface{}{
		"甲": []string{"丑", "未"},
		"乙": []string{"申", "子"},
		"丙": []string{"酉", "亥"},
		"丁": []string{"酉", "亥"},
		"戊": []string{"丑", "未"},
		"己": []string{"申", "子"},
		"庚": []string{"丑", "未"},
		"辛": []string{"寅", "午"},
		"壬": []string{"卯", "巳"},
		"癸": []string{"卯", "巳"},
	}

	for _, v := range dz[yg].([]string) {
		for _, v1 := range bzDz {
			if v == v1 {
				num++
			}
		}
	}

	for _, v := range dz[rg].([]string) {
		for _, v1 := range bzDz {
			if v == v1 {
				num++
			}
		}
	}

	return num

}

// ZhengYuan 正缘星
func (con XingModel) ZhengYuan(bz []string) int {
	bzTg := []string{
		bz[0], bz[2], bz[4], bz[6],
	}
	bzDz := []string{
		bz[1], bz[3], bz[5], bz[7],
	}
	num := 0
	rg := bz[4]
	dz := map[string]interface{}{
		"甲": []string{"午", "未", "丑"},
		"乙": []string{"申", "巳"},
		"丙": []string{"丑", "酉", "戌"},
		"丁": []string{"亥", "申"},
		"戊": []string{"子", "丑", "辰"},
		"己": []string{"寅", "亥"},
		"庚": []string{"卯", "辰", "未"},
		"辛": []string{"寅", "巳"},
		"壬": []string{"午", "未", "戌"},
		"癸": []string{"寅", "辰", "戌"},
	}
	tg := map[string]string{
		"甲": "己",
		"乙": "庚",
		"丙": "辛",
		"丁": "壬",
		"戊": "癸",
		"己": "甲",
		"庚": "乙",
		"辛": "丙",
		"壬": "丁",
		"癸": "戊",
	}

	for _, v := range dz[rg].([]string) {
		for _, v1 := range bzDz {
			if v == v1 {
				num++
			}
		}
	}

	for _, v := range bzTg {
		if v == tg[rg] {
			num++
		}
	}

	return num
}
