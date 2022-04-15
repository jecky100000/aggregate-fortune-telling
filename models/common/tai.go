/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package common

import (
	"log"
	"strconv"
	"time"
)

type TaiModel struct {
}

var TaiArray = map[string]map[string]map[string]string{
	"01": {
		"01": {"op": "2", "time": "3.9"},
		"02": {"op": "2", "time": "3.38"},
		"03": {"op": "2", "time": "4.6"},
		"04": {"op": "2", "time": "4.33"},
		"05": {"op": "2", "time": "5.1"},
		"06": {"op": "2", "time": "5.27"},
		"07": {"op": "2", "time": "5.54"},
		"08": {"op": "2", "time": "6.20"},
		"09": {"op": "2", "time": "6.45"},
		"10": {"op": "2", "time": "7.10"},
		"11": {"op": "2", "time": "7.35"},
		"12": {"op": "2", "time": "7.59"},
		"13": {"op": "2", "time": "8.22"},
		"14": {"op": "2", "time": "8.45"},
		"15": {"op": "2", "time": "9.7"},
		"16": {"op": "2", "time": "9.28"},
		"17": {"op": "2", "time": "9.49"},
		"18": {"op": "2", "time": "10.9"},
		"19": {"op": "2", "time": "10.28"},
		"20": {"op": "2", "time": "10.47"},
		"21": {"op": "2", "time": "11.5"},
		"22": {"op": "2", "time": "11.22"},
		"23": {"op": "2", "time": "11.38"},
		"24": {"op": "2", "time": "11.54"},
		"25": {"op": "2", "time": "12.8"},
		"26": {"op": "2", "time": "12.22"},
		"27": {"op": "2", "time": "12.35"},
		"28": {"op": "2", "time": "12.59"},
		"29": {"op": "2", "time": "13.10"},
		"30": {"op": "2", "time": "13.19"},
		"31": {"op": "2", "time": "13.37"},
	},
	"02": {
		"01": {"op": "2", "time": "13.44"},
		"02": {"op": "2", "time": "13.50"},
		"03": {"op": "2", "time": "13.56"},
		"04": {"op": "2", "time": "14.1"},
		"05": {"op": "2", "time": "14.5"},
		"06": {"op": "2", "time": "14.9"},
		"07": {"op": "2", "time": "14.11"},
		"08": {"op": "2", "time": "14.13"},
		"09": {"op": "2", "time": "14.14"},
		"10": {"op": "2", "time": "14.15"},
		"11": {"op": "2", "time": "14.14"},
		"12": {"op": "2", "time": "14.13"},
		"13": {"op": "2", "time": "14.11"},
		"14": {"op": "2", "time": "14.8"},
		"15": {"op": "2", "time": "14.5"},
		"16": {"op": "2", "time": "14.1"},
		"17": {"op": "2", "time": "13.56"},
		"18": {"op": "2", "time": "13.51"},
		"19": {"op": "2", "time": "13.44"},
		"20": {"op": "2", "time": "13.38"},
		"21": {"op": "2", "time": "13.30"},
		"22": {"op": "2", "time": "13.22"},
		"23": {"op": "2", "time": "13.13"},
		"24": {"op": "2", "time": "11.4"},
		"25": {"op": "2", "time": "12.54"},
		"26": {"op": "2", "time": "12.43"},
		"27": {"op": "2", "time": "12.32"},
		"28": {"op": "2", "time": "12.21"},
		"29": {"op": "2", "time": "12.8"},
	},
	"03": {
		"01": {"op": "2", "time": "11.56"},
		"02": {"op": "2", "time": "11.43"},
		"03": {"op": "2", "time": "11.29"},
		"04": {"op": "2", "time": "11.15"},
		"05": {"op": "2", "time": "11.1"},
		"06": {"op": "2", "time": "10.47"},
		"07": {"op": "2", "time": "10.32"},
		"08": {"op": "2", "time": "10.16"},
		"09": {"op": "2", "time": "10.1"},
		"10": {"op": "2", "time": "9.45"},
		"11": {"op": "2", "time": "9.28"},
		"12": {"op": "2", "time": "9.12"},
		"13": {"op": "2", "time": "8.55"},
		"14": {"op": "2", "time": "8.38"},
		"15": {"op": "2", "time": "8.21"},
		"16": {"op": "2", "time": "8.4"},
		"17": {"op": "2", "time": "7.46"},
		"18": {"op": "2", "time": "7.29"},
		"19": {"op": "2", "time": "7.11"},
		"20": {"op": "2", "time": "6.53"},
		"21": {"op": "2", "time": "6.35"},
		"22": {"op": "2", "time": "6.17"},
		"23": {"op": "2", "time": "5.58"},
		"24": {"op": "2", "time": "5.40"},
		"25": {"op": "2", "time": "5.22"},
		"26": {"op": "2", "time": "5.4"},
		"27": {"op": "2", "time": "4.45"},
		"28": {"op": "2", "time": "4.27"},
		"29": {"op": "2", "time": "4.9"},
		"30": {"op": "2", "time": "3.51"},
		"31": {"op": "2", "time": "3.33"},
	},
	"04": {
		"01": {"op": "2", "time": "3.16"},
		"02": {"op": "2", "time": "2.58"},
		"03": {"op": "2", "time": "2.41"},
		"04": {"op": "2", "time": "2.24"},
		"05": {"op": "2", "time": "2.7"},
		"06": {"op": "2", "time": "1.50"},
		"07": {"op": "2", "time": "1.33"},
		"08": {"op": "2", "time": "1.17"},
		"09": {"op": "2", "time": "1.1"},
		"10": {"op": "1", "time": "0.46"},
		"11": {"op": "1", "time": "0.30"},
		"12": {"op": "1", "time": "0.16"},
		"13": {"op": "1", "time": "0.1"},
		"14": {"op": "1", "time": "0.13"},
		"15": {"op": "1", "time": "0.27"},
		"16": {"op": "1", "time": "0.41"},
		"17": {"op": "1", "time": "0.54"},
		"18": {"op": "1", "time": "1.6"},
		"19": {"op": "1", "time": "1.19"},
		"20": {"op": "1", "time": "1.31"},
		"21": {"op": "1", "time": "1.42"},
		"22": {"op": "1", "time": "1.53"},
		"23": {"op": "1", "time": "2.4"},
		"24": {"op": "1", "time": "2.14"},
		"25": {"op": "1", "time": "2.23"},
		"26": {"op": "1", "time": "2.33"},
		"27": {"op": "1", "time": "2.41"},
		"28": {"op": "1", "time": "2.49"},
		"29": {"op": "1", "time": "2.57"},
		"30": {"op": "1", "time": "3.4"},
	},
	"05": {
		"01": {"op": "1", "time": "1.10"},
		"02": {"op": "1", "time": "3.16"},
		"03": {"op": "1", "time": "3.21"},
		"04": {"op": "1", "time": "3.26"},
		"05": {"op": "1", "time": "3.30"},
		"06": {"op": "1", "time": "3.37"},
		"07": {"op": "1", "time": "3.36"},
		"08": {"op": "1", "time": "3.39"},
		"09": {"op": "1", "time": "3.40"},
		"10": {"op": "1", "time": "3.42"},
		"11": {"op": "1", "time": "3.42"},
		"12": {"op": "1", "time": "3.42"},
		"13": {"op": "1", "time": "3.42"},
		"14": {"op": "1", "time": "3.41"},
		"15": {"op": "1", "time": "3.39"},
		"16": {"op": "1", "time": "3.37"},
		"17": {"op": "1", "time": "3.34"},
		"18": {"op": "1", "time": "3.31"},
		"19": {"op": "1", "time": "3.27"},
		"20": {"op": "1", "time": "3.23"},
		"21": {"op": "1", "time": "3.18"},
		"22": {"op": "1", "time": "3.13"},
		"23": {"op": "1", "time": "3.7"},
		"24": {"op": "1", "time": "3.1"},
		"25": {"op": "1", "time": "2.54"},
		"26": {"op": "1", "time": "2.47"},
		"27": {"op": "1", "time": "2.39"},
		"28": {"op": "1", "time": "2.31"},
		"29": {"op": "1", "time": "2.22"},
		"30": {"op": "1", "time": "2.13"},
		"31": {"op": "1", "time": "2.4"},
	},
	"06": {
		"01": {"op": "1", "time": "1.54"},
		"02": {"op": "1", "time": "1.44"},
		"03": {"op": "1", "time": "1.34"},
		"04": {"op": "1", "time": "1.23"},
		"05": {"op": "1", "time": "1.12"},
		"06": {"op": "1", "time": "1.0"},
		"07": {"op": "1", "time": "0.48"},
		"08": {"op": "1", "time": "0.36"},
		"09": {"op": "1", "time": "0.24"},
		"10": {"op": "1", "time": "0.12"},
		"11": {"op": "1", "time": "0.1"},
		"12": {"op": "1", "time": "0.14"},
		"13": {"op": "1", "time": "0.39"},
		"14": {"op": "1", "time": "0.52"},
		"15": {"op": "2", "time": "1.5"},
		"16": {"op": "2", "time": "1.18"},
		"17": {"op": "2", "time": "1.31"},
		"18": {"op": "2", "time": "1.45"},
		"19": {"op": "2", "time": "1.57"},
		"20": {"op": "2", "time": "2.10"},
		"21": {"op": "2", "time": "2.23"},
		"22": {"op": "2", "time": "2.36"},
		"23": {"op": "2", "time": "2.48"},
		"24": {"op": "2", "time": "3.1"},
		"25": {"op": "2", "time": "3.13"},
		"26": {"op": "2", "time": "3.25"},
		"27": {"op": "2", "time": "3.37"},
		"28": {"op": "2", "time": "3.49"},
		"29": {"op": "2", "time": "4.0"},
		"30": {"op": "2", "time": "4.11"},
	},
	"07": {
		"01": {"op": "2", "time": "4.22"},
		"02": {"op": "2", "time": "4.33"},
		"03": {"op": "2", "time": "4.43"},
		"04": {"op": "2", "time": "4.53"},
		"05": {"op": "2", "time": "5.2"},
		"06": {"op": "2", "time": "5.11"},
		"07": {"op": "2", "time": "5.20"},
		"08": {"op": "2", "time": "5.28"},
		"09": {"op": "2", "time": "5.36"},
		"10": {"op": "2", "time": "5.43"},
		"11": {"op": "2", "time": "5.50"},
		"12": {"op": "2", "time": "5.56"},
		"13": {"op": "2", "time": "6.2"},
		"14": {"op": "2", "time": "6.8"},
		"15": {"op": "2", "time": "6.12"},
		"16": {"op": "2", "time": "6.16"},
		"17": {"op": "2", "time": "6.20"},
		"18": {"op": "2", "time": "6.23"},
		"19": {"op": "2", "time": "6.25"},
		"20": {"op": "2", "time": "6.27"},
		"21": {"op": "2", "time": "6.29"},
		"22": {"op": "2", "time": "6.29"},
		"23": {"op": "2", "time": "6.29"},
		"24": {"op": "2", "time": "6.29"},
		"25": {"op": "2", "time": "6.28"},
		"26": {"op": "2", "time": "6.26"},
		"27": {"op": "2", "time": "6.24"},
		"28": {"op": "2", "time": "6.21"},
		"29": {"op": "2", "time": "6.17"},
		"30": {"op": "2", "time": "6.13"},
		"31": {"op": "2", "time": "6.8"},
	},
	"08": {
		"01": {"op": "2", "time": "6.3"},
		"02": {"op": "2", "time": "5.57"},
		"03": {"op": "2", "time": "5.51"},
		"04": {"op": "2", "time": "5.44"},
		"05": {"op": "2", "time": "5.36"},
		"06": {"op": "2", "time": "5.28"},
		"07": {"op": "2", "time": "5.19"},
		"08": {"op": "2", "time": "5.10"},
		"09": {"op": "2", "time": "5.0"},
		"10": {"op": "2", "time": "4.50"},
		"11": {"op": "2", "time": "4.39"},
		"12": {"op": "2", "time": "4.27"},
		"13": {"op": "2", "time": "4.15"},
		"14": {"op": "2", "time": "4.2"},
		"15": {"op": "2", "time": "3.49"},
		"16": {"op": "2", "time": "3.36"},
		"17": {"op": "2", "time": "3.21"},
		"18": {"op": "2", "time": "3.7"},
		"19": {"op": "2", "time": "2.51"},
		"20": {"op": "2", "time": "2.36"},
		"21": {"op": "2", "time": "2.20"},
		"22": {"op": "2", "time": "2.3"},
		"23": {"op": "2", "time": "1.47"},
		"24": {"op": "2", "time": "1.29"},
		"25": {"op": "2", "time": "1.12"},
		"26": {"op": "1", "time": "0.54"},
		"27": {"op": "1", "time": "0.35"},
		"28": {"op": "1", "time": "0.17"},
		"29": {"op": "1", "time": "0.2"},
		"30": {"op": "1", "time": "0.21"},
		"31": {"op": "1", "time": "0.41"},
	},
	"09": {
		"01": {"op": "1", "time": "1.0"},
		"02": {"op": "1", "time": "1.20"},
		"03": {"op": "1", "time": "1.40"},
		"04": {"op": "1", "time": "2.1"},
		"05": {"op": "1", "time": "2.21"},
		"06": {"op": "1", "time": "2.42"},
		"07": {"op": "1", "time": "3.3"},
		"08": {"op": "1", "time": "3.3"},
		"09": {"op": "1", "time": "3.24"},
		"10": {"op": "1", "time": "3.45"},
		"11": {"op": "1", "time": "4.6"},
		"12": {"op": "1", "time": "4.27"},
		"13": {"op": "1", "time": "4.48"},
		"14": {"op": "1", "time": "5.10"},
		"15": {"op": "1", "time": "5.31"},
		"16": {"op": "1", "time": "5.53"},
		"17": {"op": "1", "time": "6.14"},
		"18": {"op": "1", "time": "6.35"},
		"19": {"op": "1", "time": "6.57"},
		"20": {"op": "1", "time": "7.18"},
		"21": {"op": "1", "time": "7.39"},
		"22": {"op": "1", "time": "8.0"},
		"23": {"op": "1", "time": "8.21"},
		"24": {"op": "1", "time": "8.42"},
		"25": {"op": "1", "time": "9.2"},
		"26": {"op": "1", "time": "9.22"},
		"27": {"op": "1", "time": "9.42"},
		"28": {"op": "1", "time": "10.2"},
		"29": {"op": "1", "time": "10.21"},
		"30": {"op": "1", "time": "10.40"},
	},
	"10": {
		"01": {"op": "1", "time": "10.59"},
		"02": {"op": "1", "time": "11.18"},
		"03": {"op": "1", "time": "11.36"},
		"04": {"op": "1", "time": "11.36"},
		"05": {"op": "1", "time": "11.53"},
		"06": {"op": "1", "time": "12.11"},
		"07": {"op": "1", "time": "12.28"},
		"08": {"op": "1", "time": "12.44"},
		"09": {"op": "1", "time": "12.60"},
		"10": {"op": "1", "time": "13.16"},
		"11": {"op": "1", "time": "13.16"},
		"12": {"op": "1", "time": "13.31"},
		"13": {"op": "1", "time": "13.45"},
		"14": {"op": "1", "time": "13.59"},
		"15": {"op": "1", "time": "14.13"},
		"16": {"op": "1", "time": "14.26"},
		"17": {"op": "1", "time": "14.38"},
		"18": {"op": "1", "time": "14.50"},
		"19": {"op": "1", "time": "15.1"},
		"20": {"op": "1", "time": "15.12"},
		"21": {"op": "1", "time": "11.21"},
		"22": {"op": "1", "time": "15.31"},
		"23": {"op": "1", "time": "15.40"},
		"24": {"op": "1", "time": "15.48"},
		"25": {"op": "1", "time": "15.55"},
		"26": {"op": "1", "time": "16.1"},
		"27": {"op": "1", "time": "16.7"},
		"28": {"op": "1", "time": "16.12"},
		"29": {"op": "1", "time": "16.16"},
		"30": {"op": "1", "time": "16.20"},
		"31": {"op": "1", "time": "16.22"},
	},
	"11": {
		"01": {"op": "1", "time": "16.24"},
		"02": {"op": "1", "time": "16.25"},
		"03": {"op": "1", "time": "16.25"},
		"04": {"op": "1", "time": "16.24"},
		"05": {"op": "1", "time": "16.23"},
		"06": {"op": "1", "time": "16.21"},
		"07": {"op": "1", "time": "16.17"},
		"08": {"op": "1", "time": "16.13"},
		"09": {"op": "1", "time": "16.9"},
		"10": {"op": "1", "time": "16.3"},
		"11": {"op": "1", "time": "15.56"},
		"12": {"op": "1", "time": "15.49"},
		"13": {"op": "1", "time": "15.41"},
		"14": {"op": "1", "time": "15.32"},
		"15": {"op": "1", "time": "15.22"},
		"16": {"op": "1", "time": "15.11"},
		"17": {"op": "1", "time": "14.60"},
		"18": {"op": "1", "time": "14.47"},
		"19": {"op": "1", "time": "14.34"},
		"20": {"op": "1", "time": "14.20"},
		"21": {"op": "1", "time": "14.6"},
		"22": {"op": "1", "time": "13.50"},
		"23": {"op": "1", "time": "13.34"},
		"24": {"op": "1", "time": "13.17"},
		"25": {"op": "1", "time": "12.59"},
		"26": {"op": "1", "time": "12.40"},
		"27": {"op": "1", "time": "12.21"},
		"28": {"op": "1", "time": "12.1"},
		"29": {"op": "1", "time": "11.40"},
		"30": {"op": "1", "time": "11.18"},
	},
	"12": {
		"01": {"op": "1", "time": "10.56"},
		"02": {"op": "1", "time": "10.33"},
		"03": {"op": "1", "time": "10.9"},
		"04": {"op": "1", "time": "9.45"},
		"05": {"op": "1", "time": "9.21"},
		"06": {"op": "1", "time": "8.55"},
		"07": {"op": "1", "time": "8.29"},
		"08": {"op": "1", "time": "8.3"},
		"09": {"op": "1", "time": "7.36"},
		"10": {"op": "1", "time": "7.9"},
		"11": {"op": "1", "time": "6.42"},
		"12": {"op": "1", "time": "6.14"},
		"13": {"op": "1", "time": "5.46"},
		"14": {"op": "1", "time": "5.17"},
		"15": {"op": "1", "time": "4.48"},
		"16": {"op": "1", "time": "4.19"},
		"17": {"op": "1", "time": "3.50"},
		"18": {"op": "1", "time": "3.21"},
		"19": {"op": "1", "time": "2.51"},
		"20": {"op": "1", "time": "2.22"},
		"21": {"op": "1", "time": "1.52"},
		"22": {"op": "1", "time": "1.22"},
		"23": {"op": "1", "time": "0.52"},
		"24": {"op": "1", "time": "0.23"},
		"25": {"op": "1", "time": "0.7"},
		"26": {"op": "1", "time": "0.37"},
		"27": {"op": "2", "time": "1.6"},
		"28": {"op": "2", "time": "1.36"},
		"29": {"op": "2", "time": "2.5"},
		"30": {"op": "2", "time": "2.34"},
		"31": {"op": "2", "time": "3.3"},
	},
}

func (con TaiModel) Pin(lng float64, t string) string {

	minutes := int64((lng-120)*4) * 60
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	ttt := time.Unix(tt.Unix()+minutes, 0)
	return ttt.Format("2006-01-02 15:04:05")
}

func (con TaiModel) Really(pin string) string {
	month := time.Now().Format("01")
	day := time.Now().Format("02")

	date := TaiArray[month][day]
	minutes, _ := strconv.ParseFloat(date["time"], 64)

	t, _ := time.ParseInLocation("2006-01-02 15:04:05", pin, time.Local)

	unix1 := strconv.FormatInt(t.Unix(), 10)
	unix, _ := strconv.ParseFloat(unix1, 64)

	var tt float64

	if date["op"] == "-" {
		tt = unix - (minutes * 60)
	} else {
		tt = unix + (minutes * 60)
	}

	ttt := time.Unix(int64(tt), 0)
	log.Println(ttt)
	return ttt.Format("2006-01-02 15:04:05")
}
