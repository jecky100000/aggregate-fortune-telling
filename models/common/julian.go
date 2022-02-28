/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package common

type JuLianModel struct {
}

func (con JuLianModel) Solar(jd float64) []int {
	var (
		y4h  float64
		init float64
	)
	if jd >= 2299160.5 { //1582年10月15日,此日起是儒略日历,之前是儒略历
		y4h = 146097
		init = 1721119.5
	} else {
		y4h = 146100
		init = 1721117.5
	}
	jdr := jd - init
	yh := y4h / 4
	cen := (jdr + 0.75) / yh
	d := jdr + 0.75 - cen*yh
	ywl := 1461.0 / 4
	jy := (d + 0.75) / ywl
	d = d + 0.75 - ywl*jy + 1
	ml := 153.0 / 5
	mp := (d - 0.5) / ml
	d = (d - 0.5) - 30.6*mp + 1
	y := (100 * cen) + jy
	m := (int(mp)+2)%12 + 1
	if m < 3 {
		y = y + 1
	}
	sd := int((jd+0.5-(jd+0.5))*24*60*60 + 0.00005)
	mt := int(sd / 60)
	ss := sd % 60
	hh := mt / 60
	mt = mt % 60
	yy := (y)
	mm := (m)
	dd := (d)

	return []int{int(yy), mm, int(dd), hh, mt, ss}
}
