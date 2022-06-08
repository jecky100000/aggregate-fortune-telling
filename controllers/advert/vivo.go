/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package advert

import (
	"aggregate-fortune-telling/ay"
	"aggregate-fortune-telling/models"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Vivo struct {
}

var (
//url = "https://sandbox-marketing-api.vivo.com.cn"
//url = "https://marketing-api.vivo.com.cn"
)

type accessTokenRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		AccessToken      string `json:"access_token"`
		RefreshToken     string `json:"refresh_token"`
		TokenDate        int64  `json:"token_date"`
		RefreshTokenDate int64  `json:"refresh_token_date"`
	} `json:"data"`
}

func (con Vivo) GetCode(c *gin.Context) {
	type getCodeFrom struct {
		Code     string `form:"code"`
		ClientId int    `form:"clientId"`
		State    string `form:"state"`
	}

	var data getCodeFrom
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	log.Println(data)

	con.GetAccessToken(data.ClientId, data.Code)
}

func (con Vivo) GetAccessToken(clientId int, code string) {

	var vivo models.AdvertVivo
	ay.Db.Where("client_id = ?", clientId).First(&vivo)

	log.Println(clientId)
	log.Println(vivo)

	resp, err := http.Get("https://marketing-api.vivo.com.cn/openapi/v1/oauth2/token?client_id=" + vivo.ClientId + "&client_secret=" + vivo.ClientSecret + "&grant_type=code&code=" + code)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var rj accessTokenRes
	json.Unmarshal(body, &rj)
	ay.CreateMutiDir("log/advert/vivo")
	ioutil.WriteFile("log/advert/vivo/"+strconv.Itoa(clientId)+".txt", body, 0666)
	log.Println([]byte(body))

}

func (con Vivo) Up(cid int64, amount string, requestId, addId string) {
	var vivo models.AdvertVivo
	ay.Db.Where("id = ?", cid).First(&vivo)

	vUrl := "https://marketing-api.vivo.com.cn/openapi/v1/advertiser/behavior/upload"

	f, err := ioutil.ReadFile("log/advert/vivo/" + vivo.ClientId + ".txt")

	var r accessTokenRes
	json.Unmarshal(f, &r)
	nonce := ay.GetRandomString(10)
	vUrl += "?access_token=" + r.Data.AccessToken + "&timestamp=" + strconv.FormatInt(time.Now().Unix(), 10) + "000&nonce=" + nonce + "&advertiser_id=" + vivo.AdvertiserId
	jsonStr := `
{
    "srcType":"Web",
    "pageUrl":"` + vivo.Url + `",
    "srcId":"` + vivo.SrcId + `",
    "dataList":{
        "cvType":"SUBMIT",
        "cvTime":` + strconv.FormatInt(time.Now().Unix(), 10) + `000,
        "cvParam":"` + amount + `",
        "requestId":"` + requestId + `",
        "creativeId":"` + addId + `"
    }
}`
	req, err := http.NewRequest("POST", vUrl, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(string(body))
}
