/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package common

import (
	"encoding/json"
	"math/rand"
	"sort"
	"strconv"
)

type LineModel struct {
}

type Line struct {
	Key   int    `json:"key"`
	Value string `json:"vaule"`
	Text  string `json:"text"`
}

type Vline []Line

func (s Vline) Len() int           { return len(s) }
func (s Vline) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Vline) Less(i, j int) bool { return s[i].Value < s[j].Value }

func (con LineModel) Line() string {

	top := []string{
		"贵人相助", "财运爆发", "收获爱情", "晋升提拔", "小人退散",
	}
	bottom := []string{
		"破财", "伤灾", "横祸", "血光", "官灾",
	}

	var c Vline
	var d Vline

	for i := 0; i < 12; i++ {
		x := rand.Intn(5) + 5
		c = append(c, Line{
			Key:   i,
			Value: strconv.Itoa(x),
			Text:  "",
		})
		d = append(d, Line{
			Key:   i,
			Value: strconv.Itoa(x),
			Text:  "",
		})
	}

	sort.Stable(c)
	for k, v := range d {
		if v.Key == c[0].Key {
			d[k].Text = bottom[rand.Intn(len(bottom))]
		}

		if v.Key == c[1].Key {
			d[k].Text = bottom[rand.Intn(len(bottom))]
		}

		if d[k].Key == c[10].Key {
			d[k].Text = top[rand.Intn(len(top))]
		}

		if d[k].Key == c[11].Key {
			d[k].Text = top[rand.Intn(len(top))]
		}
	}
	b, _ := json.Marshal(d)
	return string(b)

}
