/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package common

type EightCharModel struct {
}

// Ctg 10天干
var Ctg = []string{"癸", "甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬"}

// Cdz 12地支
var Cdz = []string{"亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌"}

func (con EightCharModel) Getifx(bz []string) []int {
	b := []int{
		con.GetKey(bz[0], Ctg),
		con.GetKey(bz[1], Cdz),
		con.GetKey(bz[2], Ctg),
		con.GetKey(bz[3], Cdz),
		con.GetKey(bz[4], Ctg),
		con.GetKey(bz[5], Cdz),
		con.GetKey(bz[6], Ctg),
		con.GetKey(bz[7], Cdz),
	}
	return b
}

func (con EightCharModel) GetGanKey(bz string) int {
	return con.GetKey(bz, Ctg)
}

func (con EightCharModel) GetZhiKey(bz string) int {
	return con.GetKey(bz, Cdz)
}

func (con EightCharModel) GetKey(bz string, arr []string) int {

	for k, v := range arr {
		if v == bz {
			return k
		}
	}
	return -1

}
