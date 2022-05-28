/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package api

import (
	"aggregate-fortune-telling/ay"
	"aggregate-fortune-telling/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yuchenfw/gocrypt"
	"github.com/yuchenfw/gocrypt/rsa"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
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

// Baidu 支付
func (con BaiDuController) Baidu(oid string) (bool, string, map[string]interface{}) {
	var order models.Order
	ay.Db.First(&order, "oid = ?", oid)
	if order.Id == 0 {
		return false, "订单错误", map[string]interface{}{}
	}

	order.OutTradeNo = ay.MakeOrder(time.Now())
	if err := ay.Db.Save(&order).Error; err != nil {
		return false, "请联系管理员", map[string]interface{}{}
	}

	var pay models.Pay
	ay.Db.First(&pay, "id = ?", Appid)

	signArr := gin.H{
		"appKey":      pay.PayKey,
		"dealId":      pay.PayDealId,
		"tpOrderId":   order.OutTradeNo,
		"totalAmount": order.Amount * 100,
	}
	str := ""
	for k, v := range signArr {
		if _, ok := v.(float64); ok {
			str += k + "=" + strconv.FormatFloat(v.(float64), 'g', -1, 64) + "&"
		} else {
			str += k + "=" + v.(string) + "&"
		}
	}

	secretInfo := rsa.RSASecret{
		PublicKey:          pay.PublicKey,
		PublicKeyDataType:  gocrypt.Base64,
		PrivateKey:         pay.PrivateKey,
		PrivateKeyType:     gocrypt.PKCS1,
		PrivateKeyDataType: gocrypt.Base64,
	}

	handleRSA := rsa.NewRSACrypt(secretInfo) //RSA
	sign, err := handleRSA.Encrypt(strings.TrimRight(str, "&"), gocrypt.HEX)
	if err != nil {
		return false, err.Error(), map[string]interface{}{}
	}

	orderInfo := map[string]interface{}{
		"dealId":          pay.PayDealId,
		"appKey":          pay.PayKey,
		"totalAmount":     order.Amount * 100,
		"tpOrderId":       order.OutTradeNo,
		"dealTitle":       order.Des,
		"signFieldsRange": "1",
		"rsaSign":         sign,
	}

	return true, "success", orderInfo
	//return true, "success", map[string]interface{}{}

}
