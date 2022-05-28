/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

import (
	"aggregate-fortune-telling/ay"
	"aggregate-fortune-telling/sdk/tencentyun"
	"bytes"
	"io/ioutil"
	"net/http"
)

type ImModel struct {
}

var (
	ImUrl   = "https://console.tim.qq.com"
	UserSig string
)

func (con ImModel) GetUserSig() string {
	UserSig, _ = tencentyun.GenUserSig(ay.Yaml.GetInt("im.appid"), ay.Yaml.GetString("im.key"), ay.Yaml.GetString("im.root"), 3600*24*30*365)
	return UserSig
}

func (con ImModel) HttpPost(url, phone string) (int, string) {
	post := "{\"UserID\":\"" + phone +
		"\",\"Nick\":\"" + phone +
		"\"}"

	var jsonStr = []byte(post)

	resp, err := http.Post(ImUrl+url+"?sdkappid="+ay.Yaml.GetString("im.appid")+"&identifier="+ay.Yaml.GetString("im.root")+"&usersig="+con.GetUserSig()+"&random=99999999&contenttype=json",
		"application/json",
		bytes.NewBuffer(jsonStr))
	if err != nil {
		return 0, err.Error()
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err.Error()
	}

	return 1, string(body)
}
