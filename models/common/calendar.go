/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package common

import (
	"fmt"
	"github.com/6tail/lunar-go/calendar"
	"math"
	"strconv"
	"time"
)

type CalendarModel struct {
}

// YiJi 日宜 日忌
func (con CalendarModel) YiJi(lunar *calendar.Lunar) (string, string) {

	yi, ji := "", ""

	// 日宜
	for i := lunar.GetDayYi().Front(); i != nil; i = i.Next() {
		yi += i.Value.(string) + " "
	}

	// 日忌
	for i := lunar.GetDayJi().Front(); i != nil; i = i.Next() {
		ji += i.Value.(string) + " "
	}

	return yi, ji
}

func (con CalendarModel) Get(y, m, d int) map[string]interface{} {
	solar := calendar.NewSolarFromYmd(y, m, d)
	solarLunar := solar.GetLunar()

	//lunar := calendar.NewLunarFromYmd(y, m, d)
	lunar := solar.GetLunar()

	// 节日
	jieri := con.GetJieRi(solar)

	// 宜忌
	yi, ji := con.YiJi(solarLunar)

	// 八字
	bazi := con.GetBazi(y, m, d)

	nongliS := fmt.Sprintf("%s", solarLunar)
	nongliRune := []rune(nongliS)
	nongli := string(nongliRune[5:])

	jq := fmt.Sprintf("%s", solarLunar.GetJieQi())

	if jq == "" {
		next := solarLunar.GetNextJieQi()
		str := "距离下一节气" + next.GetName() + "，还有"
		loc, _ := time.LoadLocation("Local")

		tt := next.GetSolar().ToYmd()
		ttTime, _ := time.ParseInLocation("2006-01-02", tt, loc)
		vm := strconv.Itoa(m)
		vd := strconv.Itoa(d)
		if len(vm) == 1 {
			vm = "0" + vm
		}

		if len(vd) == 1 {
			vd = "0" + vd
		}

		now := strconv.Itoa(y) + "-" + vm + "-" + vd
		nowTime, _ := time.ParseInLocation("2006-01-02", now, loc)
		t := ttTime.Unix() - nowTime.Unix()
		jq = str + strconv.Itoa(int(math.Ceil(float64(t)/3600/24))) + "天"
	}

	return map[string]interface{}{
		"jieQi":   jq,
		"xingZuo": fmt.Sprintf("%s", solar.GetXingZuo()),
		"baZi":    bazi,
		"yi":      yi,
		"ji":      ji,
		"lunar":   nongli,
		"jieRi":   jieri,
		"week":    "星期" + lunar.GetWeekInChinese(),
		"dayLu":   lunar.GetDayLu(),
	}
}

func (con CalendarModel) GetBazi(y, m, d int) map[string]map[string]string {

	vh := time.Now().Format("15")
	h, _ := strconv.Atoi(vh)
	solar := calendar.NewSolar(y, m, d, h, 0, 0)
	lunar := solar.GetLunar()
	bazi := lunar.GetEightChar()

	arr := map[string]map[string]string{
		"year": {
			"bazi":      bazi.GetYear(),
			"shengxiao": lunar.GetYearShengXiao(),
			"nayin":     lunar.GetYearNaYin(),
		},
		"month": {
			"bazi":      bazi.GetMonth(),
			"shengxiao": lunar.GetMonthShengXiao(),
			"nayin":     lunar.GetYearNaYin(),
		},
		"day": {
			"bazi":      bazi.GetDay(),
			"shengxiao": lunar.GetDayShengXiao(),
			"nayin":     lunar.GetDayNaYin(),
		},
		"hour": {
			"bazi":      bazi.GetTime(),
			"shengxiao": lunar.GetTimeShengXiao(),
			"nayin":     lunar.GetTimeNaYin(),
		},
	}

	return arr
}

// GetJieRi 节日
func (con CalendarModel) GetJieRi(lunar *calendar.Solar) string {

	jieri := ""

	for i := lunar.GetFestivals().Front(); i != nil; i = i.Next() {
		jieri += i.Value.(string) + " "
	}

	for i := lunar.GetOtherFestivals().Front(); i != nil; i = i.Next() {
		jieri += i.Value.(string) + " "
	}
	return jieri
}
