/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type BaiDuController struct {
}

type baidu struct {
	Openid           string `json:"openid"`
	SessionKey       string `json:"session_key"`
	Errno            int    `json:"errno"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (con BaiDuController) GetOpenid(code, key, secret string) (int, string, string) {
	url := "https://spapi.baidu.com/oauth/jscode2sessionkey?code=" + code + "&client_id=" + key + "&sk=" + secret
	log.Println(url)
	response, err := http.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		return 0, string(rune(http.StatusServiceUnavailable)), ""
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err.Error(), ""
	}
	defer response.Body.Close()
	var j baidu
	json.Unmarshal(body, &j)

	if j.Errno != 0 {
		return 0, j.ErrorDescription, ""
	}

	return 1, j.Openid, j.SessionKey
}
