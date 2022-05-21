/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package common

import (
	"fmt"
	"gin/ay"
	"gin/models"
	"github.com/6tail/lunar-go/calendar"
	"strconv"
	"time"
)

type HaulModel struct {
}

func (con HaulModel) Detail(y, m, d, h, i, s int) (*calendar.EightChar, []string, string, map[string]interface{}) {
	solar := calendar.NewSolar(y, m, d, h, i, s)
	lunar := solar.GetLunar()
	BaZi := lunar.GetEightChar()
	BaZi.SetSect(1)

	// 获取八字
	bz := con.GetBZ(BaZi)

	// 命格
	var jiazi models.JiaZi
	ay.Db.First(&jiazi, "jiazi = ?", bz[0]+bz[1])

	// 事业 财富 爱情
	var rgnm models.Rgnm
	ay.Db.First(&rgnm, "rgz = ?", bz[4]+bz[5])

	hour := ""
	vh := h

	if h == -1 {
		hour = "未知"
		vh = 0
	} else {
		hour = LunarModel{}.HourChinese(h)
	}

	//节气
	str := con.GetYunShi(bz)

	return BaZi, bz, hour, map[string]interface{}{
		"bazi":   bz, // 八字
		"lunar":  fmt.Sprintf("%s", lunar),
		"hour":   hour,
		"siji":   con.GetSiJi(y, m, d, vh),
		"ming":   jiazi.NaYin,
		"rglm":   rgnm,
		"text":   str,
		"animal": lunar.GetYearShengXiaoByLiChun(),
	}
}

func (con HaulModel) GetYunShi(bz []string) string {
	lunar := calendar.NewLunarFromYmd(2022, 12, 7)
	//lunar := solar.GetLunar()
	jieQi := lunar.GetJieQiTable()

	var arr []string

	j := 0

	key := []string{}
	value := []string{}

	start := ""
	for i := lunar.GetJieQiList().Front(); i != nil; i = i.Next() {
		name := i.Value.(string)
		arr = append(arr, jieQi[name].ToYmdHms())

		if j > 24 {
			break
		}

		if j == 0 {
			start = jieQi[name].ToYmdHms()
			j++
			continue
		} else if j%2 == 0 {
			s, _ := time.Parse("2006-01-02 15:04:05", jieQi[name].ToYmdHms())
			solarL := calendar.NewLunarFromYmd(s.Year(), int(s.Month()), s.Day())
			ss := solarL.GetEightChar().GetMonth()
			key = append(key, start+" ~ "+jieQi[name].ToYmdHms())
			start = jieQi[name].ToYmdHms()
			value = append(value, ss)
		} else if j == 24 {
			s, _ := time.Parse("2006-01-02 15:04:05", jieQi[name].ToYmdHms())
			solarL := calendar.NewLunarFromYmd(s.Year(), int(s.Month()), s.Day())
			ss := solarL.GetEightChar().GetMonth()
			key = append(key, start+" ~ "+jieQi[name].ToYmdHms())
			value = append(value, ss)
		}

		j++
	}
	str := ""
	startStr := "<font color='#ffcc33'>"

	for k, v := range key {
		var liuyue models.LiuYue
		var tengods models.TenGods
		ay.Db.First(&liuyue, "tian_gan = ? and yue_gan = ?", bz[4], value[k])
		ay.Db.First(&tengods, "name = ? ", liuyue.TenGods)
		str += startStr + v + " （" + value[k] + "月）</font><br>" + tengods.Content + "</font><br>"
	}
	return str
}

func (con HaulModel) GetBZ(BaZi *calendar.EightChar) []string {
	year := []rune(BaZi.GetYear())
	month := []rune(BaZi.GetMonth())
	day := []rune(BaZi.GetDay())
	hour := []rune(BaZi.GetTime())

	return []string{
		string(year[:1]),
		string(year[1:2]),
		string(month[:1]),
		string(month[1:2]),
		string(day[:1]),
		string(day[1:2]),
		string(hour[:1]),
		string(hour[1:2]),
	}
}

func (con HaulModel) GetSiJi(y, m, d, h int) string {
	solar := calendar.NewSolarFromYmd(y, m, d)
	lunar := solar.GetLunar()
	jieQi := lunar.GetJieQiTable()

	ms := strconv.Itoa(m)
	if len(ms) == 1 {
		ms = "0" + ms
	}

	ds := strconv.Itoa(d)
	if len(ds) == 1 {
		ds = "0" + ds
	}

	birth := strconv.Itoa(y) + "-" + ms + "-" + ds + " " + strconv.Itoa(h) + ":00:00"

	stamp, _ := time.ParseInLocation("2006-01-02 15:04:05", birth, time.Local)

	last_jq := ""
	j := 0

	for i := lunar.GetJieQiList().Front(); i != nil; i = i.Next() {
		name := i.Value.(string)
		ti, _ := time.ParseInLocation("2006-01-02 15:04:05", jieQi[name].ToYmdHms(), time.Local)
		if j > 24 {
			break
		}
		j++
		if stamp.Unix() > ti.Unix() {

			switch name {
			case "DA_XUE":
				name = "大雪"
			case "DONG_ZHI":
				name = "冬至"
			case "XIAO_HAN":
				name = "小寒"
			case "DA_HAN":
				name = "大寒"
			case "LI_CHUN":
				name = "立春"
			case "YU_SHUI":
				name = "雨水"
			case "JING_ZHE":
				name = "惊蛰"
			}

			last_jq = name
			continue
		} else {
			break
		}

	}

	switch last_jq {
	case "立春", "雨水", "惊蛰", "春分", "清明", "谷雨":
		return "春"
	case "立夏", "小满", "芒种", "夏至", "小暑", "大暑":
		return "夏"
	case "立秋", "处暑", "白露", "秋分", "寒露", "霜降":
		return "秋"
	case "立冬", "小雪", "大雪", "冬至", "小寒", "大寒":
		return "冬"
	}
	return ""

}
