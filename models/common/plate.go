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
	"fmt"
	"github.com/6tail/lunar-go/calendar"
	"log"
	"strconv"
	"time"
	"unicode/utf8"
)

type PlateModel struct {
}

func (con PlateModel) Detail(y, m, d, h, i, s, gender int) map[string]interface{} {
	solar := calendar.NewSolar(y, m, d, h, i, s)
	lunar := solar.GetLunar()
	BaZi := lunar.GetEightChar()
	BaZi.SetSect(1)

	// 获取八字
	bzArr := HaulModel{}.GetBZ(BaZi)
	//log.Println(bzArr)

	b := EightCharModel{}.Getifx(bzArr)
	//log.Println(b)

	cangan := CanGanModel{}.Get(b)
	log.Println(cangan)

	bzTypeArr := []map[string]interface{}{
		WuXingModel{}.Get(bzArr[0]),
		WuXingModel{}.Get(bzArr[1]),
		WuXingModel{}.Get(bzArr[2]),
		WuXingModel{}.Get(bzArr[3]),
		WuXingModel{}.Get(bzArr[4]),
		WuXingModel{}.Get(bzArr[5]),
		WuXingModel{}.Get(bzArr[6]),
		WuXingModel{}.Get(bzArr[7]),
	}
	//log.Println(bzTypeArr)
	xingYun := []string{
		AstrologyModel{}.Get(bzArr[4], bzArr[1]),
		AstrologyModel{}.Get(bzArr[4], bzArr[3]),
		AstrologyModel{}.Get(bzArr[4], bzArr[5]),
		AstrologyModel{}.Get(bzArr[4], bzArr[7]),
	}
	//log.Println(xingYun)

	ziZuo := []string{
		AstrologyModel{}.Get(bzArr[0], bzArr[1]),
		AstrologyModel{}.Get(bzArr[2], bzArr[3]),
		AstrologyModel{}.Get(bzArr[4], bzArr[5]),
		AstrologyModel{}.Get(bzArr[6], bzArr[7]),
	}
	//log.Println(ziZuo)

	shiShen := []string{

		BaZi.GetYearShiShenGan(),
		BaZi.GetMonthShiShenGan(),
		BaZi.GetDayShiShenGan(),
		BaZi.GetTimeShiShenGan(),
	}

	//log.Println(shiShen)

	nayin := []string{
		BaZi.GetYearNaYin(),
		BaZi.GetMonthNaYin(),
		BaZi.GetDayNaYin(),
		BaZi.GetTimeNaYin(),
	}
	//log.Println(nayin)

	xunkong := []string{
		lunar.GetYearXunKong(),
		lunar.GetMonthXunKong(),
		lunar.GetDayXunKong(),
		lunar.GetTimeXunKong(),
	}
	//log.Println(xunkong)

	myKey := b[4]

	yunStr := ""

	var yun *calendar.Yun

	if gender == 1 {
		yun = BaZi.GetYun(1)
	} else {
		yun = BaZi.GetYun(0)
	}

	yunStr = fmt.Sprintf("出生%d年%d个月%d天后起运", yun.GetStartYear(), yun.GetStartMonth(), yun.GetStartDay())

	//log.Println(yunStr)

	daYunArr := yun.GetDaYun()

	//age, _, _ := GetTimeFromStrDate(y, m, d)
	xz := 0
	daYunNum := 0

	nowY, _ := strconv.Atoi(time.Now().Format("2006"))
	var daYunA []interface{}
	for _, daYun := range daYunArr {
		check := 0

		if nowY-daYun.GetStartYear() < 10 && y != daYun.GetStartYear() && xz == 0 {
			check = 1
			xz++
			daYunNum = daYun.GetIndex()
		} else {
			check = 0
		}

		gz := daYun.GetGanZhi()
		length := utf8.RuneCountInString(gz)
		vRune := []rune(gz)

		gan, zhi := "", ""
		if daYun.GetIndex() == 0 {
			gan, zhi = "", ""
		} else {
			gan = string(vRune[:1])
			zhi = string(vRune[1:2])
		}

		ganKey := EightCharModel{}.GetGanKey(gan)
		zhiKey := EightCharModel{}.GetZhiKey(zhi)

		var cg []interface{}
		if length > 1 {
			cg = CanGanModel{}.GetOne(b, zhiKey)
		}

		sv := TenGodModel{}.Zhi(zhi, bzArr[4], 1)
		if zhi == "" {
			sv = ""
		}
		daYunA = append(daYunA, map[string]interface{}{
			"year":    daYun.GetStartYear(),
			"age":     daYun.GetStartAge(),
			"check":   check,
			"xingYun": AstrologyModel{}.Get(bzArr[4], zhi),
			"ziZuo":   AstrologyModel{}.Get(gan, zhi),
			"xunKong": XunKongModel{}.Get(gan + zhi),
			"naYin":   NaYinModel{}.Get(gan + zhi),
			"cangGan": cg,
			"gan": []interface{}{
				gan,
				WuXingModel{}.Get(gan),
				TenGodModel{}.Get(ganKey, myKey, 1),
			},
			"zhi": []interface{}{
				zhi,
				WuXingModel{}.Get(zhi),
				sv,
			},
			"shiShen": TenGodModel{}.Get(ganKey, myKey, 0),
		})
	}

	liuNianArr := daYunArr[daYunNum].GetLiuNian()
	xz = 0
	liuNianNum := 0
	var liuNianA []interface{}
	for _, liuNian := range liuNianArr {
		check := 0

		if nowY == liuNian.GetYear() && xz == 0 {
			check = 1
			xz++
			liuNianNum = liuNian.GetIndex()
		} else {
			check = 0
			//liuNianNum = 0
		}

		gz := liuNian.GetGanZhi()
		length := utf8.RuneCountInString(gz)
		vRune := []rune(gz)

		gan := string(vRune[:1])
		zhi := string(vRune[1:2])

		ganKey := EightCharModel{}.GetGanKey(gan)
		zhiKey := EightCharModel{}.GetZhiKey(zhi)

		var cg []interface{}
		if length > 1 {
			cg = CanGanModel{}.GetOne(b, zhiKey)
		}

		sv := TenGodModel{}.Zhi(zhi, bzArr[4], 1)
		if zhi == "" {
			sv = ""
		}
		liuNianA = append(liuNianA, map[string]interface{}{
			"year":    liuNian.GetYear(),
			"age":     liuNian.GetAge(),
			"check":   check,
			"xingYun": AstrologyModel{}.Get(bzArr[4], zhi),
			"ziZuo":   AstrologyModel{}.Get(gan, zhi),
			"xunKong": XunKongModel{}.Get(gan + zhi),
			"naYin":   NaYinModel{}.Get(gan + zhi),
			"cangGan": cg,
			"gan": []interface{}{
				gan,
				WuXingModel{}.Get(gan),
				TenGodModel{}.Get(ganKey, myKey, 1),
			},
			"zhi": []interface{}{
				zhi,
				WuXingModel{}.Get(zhi),
				sv,
			},
			"shiShen": TenGodModel{}.Get(ganKey, myKey, 0),
		})
	}

	jq := []string{
		"立春",
		"惊蛰",
		"清明",
		"立夏",
		"芒种",
		"小暑",
		"立秋",
		"白露",
		"寒露",
		"立冬",
		"大雪",
		"小寒",
	}

	liuYueArr := liuNianArr[0].GetLiuYue()

	solar1 := calendar.NewSolarFromYmd(liuNianArr[liuNianNum].GetYear(), 2, 4)
	lunar1 := solar1.GetLunar()
	jieQi := lunar1.GetJieQiTable()

	var liuYueA []interface{}
	for _, liuYue := range liuYueArr {

		gz := liuYue.GetGanZhi()
		length := utf8.RuneCountInString(gz)
		vRune := []rune(gz)

		gan := string(vRune[:1])
		zhi := string(vRune[1:2])

		ganKey := EightCharModel{}.GetGanKey(gan)
		zhiKey := EightCharModel{}.GetZhiKey(zhi)

		sv := TenGodModel{}.Zhi(zhi, bzArr[4], 1)
		if zhi == "" {
			sv = ""
		}

		var cg []interface{}
		if length > 1 {
			cg = CanGanModel{}.GetOne(b, zhiKey)
		}
		//name := ""
		//
		ti := ""
		for ii := lunar1.GetJieQiList().Front(); ii != nil; ii = ii.Next() {
			if jq[liuYue.GetIndex()] == ii.Value.(string) {
				ti = jieQi[ii.Value.(string)].ToYmdHms()
			}
		}
		loc, _ := time.LoadLocation("Local")
		theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", ti, loc)
		ti = time.Unix(theTime.Unix(), 0).Format("01/02")

		liuYueA = append(liuYueA, map[string]interface{}{
			"month":   jq[liuYue.GetIndex()],
			"time":    ti,
			"xingYun": AstrologyModel{}.Get(bzArr[4], zhi),
			"ziZuo":   AstrologyModel{}.Get(gan, zhi),
			"xunKong": XunKongModel{}.Get(gan + zhi),
			"naYin":   NaYinModel{}.Get(gan + zhi),
			"cangGan": cg,
			"gan": []interface{}{
				gan,
				WuXingModel{}.Get(gan),
				TenGodModel{}.Get(ganKey, myKey, 1),
			},
			"zhi": []interface{}{
				zhi,
				WuXingModel{}.Get(zhi),
				sv,
			},
			"shiShen": TenGodModel{}.Get(ganKey, myKey, 0),
		})

		//fmt.Printf("流月[%d] = %s月 %s\n", liuYue.GetIndex(), liuYue.GetMonthInChinese(), liuYue.GetGanZhi())
	}

	return map[string]interface{}{
		"xingZuo":  fmt.Sprintf("%s", solar.GetXingZuo()),
		"baZi":     bzArr,     // 八字
		"baZiType": bzTypeArr, // 八字类型
		"cangGan":  cangan,
		"xingYun":  xingYun,
		"ziZuo":    ziZuo,
		"lunar":    fmt.Sprintf("%s", lunar) + " " + LunarModel{}.HourChinese(h) + "时",
		"shiShen":  shiShen,
		"naYin":    nayin,
		"xunKong":  xunkong,
		"qiYun":    yunStr,
		"daYun":    daYunA,
		"liuNian":  liuNianA,
		"liuYue":   liuYueA,
		"animal":   lunar.GetYearShengXiaoByLiChun(),
	}

	//log.Println(liuNianA)
}

func (con PlateModel) GetLiuNian(y, m, d, h, i, s, gender, key int) []interface{} {
	solar := calendar.NewSolar(y, m, d, h, i, s)
	lunar := solar.GetLunar()
	BaZi := lunar.GetEightChar()
	BaZi.SetSect(1)

	// 获取八字
	bzArr := HaulModel{}.GetBZ(BaZi)

	b := EightCharModel{}.Getifx(bzArr)

	var yun *calendar.Yun

	if gender == 1 {
		yun = BaZi.GetYun(1)
	} else {
		yun = BaZi.GetYun(0)
	}

	daYunArr := yun.GetDaYun()

	nowY, _ := strconv.Atoi(time.Now().Format("2006"))

	liuNianArr := daYunArr[key].GetLiuNian()
	xz := 0
	//liuNianNum := 0
	var liuNianA []interface{}
	for _, liuNian := range liuNianArr {
		check := 0

		if nowY == liuNian.GetYear() && xz == 0 {
			check = 1
			xz++
			//liuNianNum = liuNian.GetIndex()
		} else {
			check = 0
			//liuNianNum = 0
		}

		gz := liuNian.GetGanZhi()
		length := utf8.RuneCountInString(gz)
		vRune := []rune(gz)

		gan := string(vRune[:1])
		zhi := string(vRune[1:2])

		ganKey := EightCharModel{}.GetGanKey(gan)
		zhiKey := EightCharModel{}.GetZhiKey(zhi)

		var cg []interface{}
		if length > 1 {
			cg = CanGanModel{}.GetOne(b, zhiKey)
		}

		sv := TenGodModel{}.Zhi(zhi, bzArr[4], 1)
		if zhi == "" {
			sv = ""
		}
		liuNianA = append(liuNianA, map[string]interface{}{
			"year":    liuNian.GetYear(),
			"age":     liuNian.GetAge(),
			"check":   check,
			"xingYun": AstrologyModel{}.Get(bzArr[4], zhi),
			"ziZuo":   AstrologyModel{}.Get(gan, zhi),
			"xunKong": XunKongModel{}.Get(gan + zhi),
			"naYin":   NaYinModel{}.Get(gan + zhi),
			"cangGan": cg,
			"gan": []interface{}{
				gan,
				WuXingModel{}.Get(gan),
				TenGodModel{}.Get(ganKey, b[4], 1),
			},
			"zhi": []interface{}{
				zhi,
				WuXingModel{}.Get(zhi),
				sv,
			},
			"shiShen": TenGodModel{}.Get(ganKey, b[4], 0),
		})
	}

	return liuNianA

}

func (con PlateModel) GetLiuYue(y, m, d, h, i, s, gender, key, index int) []interface{} {
	solar := calendar.NewSolar(y, m, d, h, i, s)
	lunar := solar.GetLunar()
	BaZi := lunar.GetEightChar()
	BaZi.SetSect(1)

	// 获取八字
	bzArr := HaulModel{}.GetBZ(BaZi)

	b := EightCharModel{}.Getifx(bzArr)

	var yun *calendar.Yun

	if gender == 1 {
		yun = BaZi.GetYun(1)
	} else {
		yun = BaZi.GetYun(0)
	}

	daYunArr := yun.GetDaYun()

	//nowY, _ := strconv.Atoi(time.Now().Format("2006"))

	jq := []string{
		"立春",
		"惊蛰",
		"清明",
		"立夏",
		"芒种",
		"小暑",
		"立秋",
		"白露",
		"寒露",
		"立冬",
		"大雪",
		"小寒",
	}

	// log.Println(jq[0])

	liuNianArr := daYunArr[index].GetLiuNian()
	liuYueArr := liuNianArr[key].GetLiuYue()

	solar1 := calendar.NewSolarFromYmd(liuNianArr[key].GetYear(), 2, 4)
	lunar1 := solar1.GetLunar()
	jieQi := lunar1.GetJieQiTable()

	var liuYueA []interface{}
	for _, liuYue := range liuYueArr {

		gz := liuYue.GetGanZhi()
		length := utf8.RuneCountInString(gz)
		vRune := []rune(gz)

		gan := string(vRune[:1])
		zhi := string(vRune[1:2])

		ganKey := EightCharModel{}.GetGanKey(gan)
		zhiKey := EightCharModel{}.GetZhiKey(zhi)

		sv := TenGodModel{}.Zhi(zhi, bzArr[4], 1)
		if zhi == "" {
			sv = ""
		}

		var cg []interface{}
		if length > 1 {
			cg = CanGanModel{}.GetOne(b, zhiKey)
		}
		//name := ""
		//
		ti := ""
		for ii := lunar1.GetJieQiList().Front(); ii != nil; ii = ii.Next() {
			if jq[liuYue.GetIndex()] == ii.Value.(string) {
				ti = jieQi[ii.Value.(string)].ToYmdHms()
			}
		}
		loc, _ := time.LoadLocation("Local")
		theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", ti, loc)
		ti = time.Unix(theTime.Unix(), 0).Format("01/02")

		liuYueA = append(liuYueA, map[string]interface{}{
			"month":   jq[liuYue.GetIndex()],
			"time":    ti,
			"xingYun": AstrologyModel{}.Get(bzArr[4], zhi),
			"ziZuo":   AstrologyModel{}.Get(gan, zhi),
			"xunKong": XunKongModel{}.Get(gan + zhi),
			"naYin":   NaYinModel{}.Get(gan + zhi),
			"cangGan": cg,
			"gan": []interface{}{
				gan,
				WuXingModel{}.Get(gan),
				TenGodModel{}.Get(ganKey, b[4], 1),
			},
			"zhi": []interface{}{
				zhi,
				WuXingModel{}.Get(zhi),
				sv,
			},
			"shiShen": TenGodModel{}.Get(ganKey, b[4], 0),
		})

	}

	return liuYueA

}
