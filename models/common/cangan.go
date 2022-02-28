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
	"unicode/utf8"
)

type CanGanModel struct {
}

var (
	bz   []int
	desc map[string][]interface{}
	//desc  []interface{}
	value []string
	tdz   = []string{"壬甲", "癸", "癸己辛", "丙甲戊", "乙", "乙戊癸", "戊丙庚", "己丁", "丁己乙", "壬庚戊", "辛", "辛戊丁"}
	ctg   = []string{"癸", "甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬"}
)

func (con CanGanModel) Get(b []int) map[string][]interface{} {
	//var desc []interface{}
	bz = b
	desc = map[string][]interface{}{}
	value = []string{
		tdz[b[1]],
		tdz[b[3]],
		tdz[b[5]],
		tdz[b[7]],
	}

	desc["year"] = con.GetDesc("year", 0)
	desc["month"] = con.GetDesc("month", 1)
	desc["day"] = con.GetDesc("day", 2)
	desc["hour"] = con.GetDesc("hour", 3)
	return desc
}

func (con CanGanModel) GetOne(b []int, ganZhi int) []interface{} {
	bz = b
	desc = map[string][]interface{}{}
	value = []string{
		tdz[ganZhi],
	}
	return con.GetDesc("cangGan", 0)
}

func (con CanGanModel) GetDesc(field string, index int) []interface{} {

	arr := []interface{}{}

	vRune := []rune(value[index])
	length := utf8.RuneCountInString(value[index])
	//val := ""
	//wx := map[string]interface{}{}
	if length < 1 {
		return []interface{}{}
	}
	val := string(vRune[:1])
	wx := WuXingModel{}.Get(val)
	arr = append(arr, map[string]interface{}{
		"type":  wx["type"],
		"value": val,
		"attar": TenGodModel{}.Get(con.GetKey(val), bz[4], 0),
	})

	if length < 2 {
		goto End
	}
	val = string(vRune[1:2])
	wx = WuXingModel{}.Get(val)
	arr = append(arr, map[string]interface{}{
		"type":  wx["type"],
		"value": val,
		"attar": TenGodModel{}.Get(con.GetKey(val), bz[4], 0),
	})

	if length < 3 {
		goto End
	}
	val = string(vRune[2:3])

	wx = WuXingModel{}.Get(val)
	arr = append(arr, map[string]interface{}{
		"type":  wx["type"],
		"value": val,
		"attar": TenGodModel{}.Get(con.GetKey(val), bz[4], 0),
	})

End:

	//desc = append(desc, arr)
	return arr
}

func (con CanGanModel) GetKey(value string) int {
	for k, v := range ctg {
		if v == value {
			return k
		}
	}
	return 0
}
