/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package common

import (
	"strings"
)

type WuXingModel struct {
}

// 五行
func (con WuXingModel) Get(value string) map[string]interface{} {
	switch value {
	case "子", "壬":
		return map[string]interface{}{
			"type":  3, // 金木水火土 序号
			"attar": 2, // 1阴 2阳
			"value": "水",
		}
	case "癸", "亥":
		return map[string]interface{}{
			"type":  3,
			"attar": 1,
			"value": "水",
		}

	case "甲", "寅":
		return map[string]interface{}{
			"type":  2,
			"attar": 1,
			"value": "木",
		}
	case "乙", "卯":
		return map[string]interface{}{
			"type":  2,
			"attar": 2,
			"value": "木",
		}
	case "丙", "午":
		return map[string]interface{}{
			"type":  4,
			"attar": 1,
			"value": "火",
		}
	case "丁", "巳":
		return map[string]interface{}{
			"type":  4,
			"attar": 2,
			"value": "火",
		}
	case "庚", "申":
		return map[string]interface{}{
			"type":  1,
			"attar": 2,
			"value": "金",
		}
	case "辛", "酉":
		return map[string]interface{}{
			"type":  1,
			"attar": 1,
			"value": "金",
		}
	case "戊", "辰", "戌":
		return map[string]interface{}{
			"type":  5,
			"attar": 1,
			"value": "土",
		}
	case "己", "丑", "未":
		return map[string]interface{}{
			"type":  5,
			"attar": 2,
			"value": "土",
		}
	default:
		return map[string]interface{}{}
	}
}

// Calc 计算八字含五行个数
func (con WuXingModel) Calc(bz []string) map[string]int {
	wxArr := map[string]int{
		"金": 0, "木": 0, "水": 0, "火": 0, "土": 0,
	}

	for _, v := range bz {
		res := WuXingModel{}.Get(v)
		if value, ok := res["value"].(string); ok {
			wxArr[value] += 1
		}
	}
	return wxArr
}

func (con WuXingModel) Defect(bz []string) string {
	wxArr := con.Calc(bz)
	str := "偏枯缺"

	for k, v := range wxArr {
		if v == 0 {
			str += k + "、"
		}
	}

	str = strings.TrimRight(str, "、")

	if !strings.Contains(str, "、") {
		return "五行俱全"
	}
	return str
}

func (con WuXingModel) Wang(bz []string) string {
	wxArr := con.Calc(bz)
	str := ""

	for k, v := range wxArr {
		if v >= 3 {
			str += k + "、"
		}
	}

	str = strings.TrimRight(str, "、") + "旺"

	if str == "旺" {
		res := WuXingModel{}.Get(bz[4])
		if value, ok := res["value"].(string); ok {
			str = value + "旺"
		}
	}
	return str
}
