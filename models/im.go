/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

import (
	"fmt"
	"gin/ay"
	"gin/sdk/tencentyun"
)

type ImModel struct {
}

var (
	ImUrl   = "https://console.tim.qq.com"
	UserSig string
)

func S() {
	UserSig, _ = tencentyun.GenUserSig(ay.Yaml.GetInt("im.appid"), ay.Yaml.GetString("im.key"), ay.Yaml.GetString("im.root"), 3600*24*30*365)
	fmt.Println(UserSig)
}

func MakeAccount(account string) {

}
