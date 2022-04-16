/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package api

import (
	"fmt"
	"gin/ay"
	"gin/models"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

type LoginController struct {
}

type GetLoginForm struct {
	Code   string `form:"code" binding:"required"`
	Phone  string `form:"phone"`
	Openid string `form:"openid"`
}

// Login 登入
func (con LoginController) Login(c *gin.Context) {
	var getForm GetLoginForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, "400", ay.Validator{}.Translate(err), gin.H{})
		return
	}

	res := 0
	token, session_key := "", ""

	if Appid == 1 {
		res, token = con.web(getForm.Phone, getForm.Code)
	} else {
		res, token, session_key = con.xcx(getForm.Code, Appid)
		//res, token, session_key = 1, "", ""
		if res == 0 {
			token = session_key
		}
	}

	if res == 0 {
		ay.Json{}.Msg(c, "400", token, gin.H{})
	} else if res == 201 {
		ay.Json{}.Msg(c, "201", "请绑定手机", gin.H{
			"session_key": session_key,
			"openid":      token,
		})
		return
	} else {
		var user models.User
		ay.Db.First(&user, "phone = ?", getForm.Phone)

		ay.Json{}.Msg(c, "200", "success", gin.H{
			"token":    token,
			"avatar":   ay.Domain + user.Avatar,
			"phone":    user.Phone,
			"nickname": user.NickName,
		})
	}

}

// AuthCode 验证码
func (con LoginController) AuthCode(phone, code string) (int, string) {
	var sms models.Sms
	ay.Db.Order("id desc").First(&sms, "phone = ? and code = ?", phone, code)

	if sms.Id == 0 {
		return 0, "验证码不存在"
	} else if sms.Status == 1 || time.Now().Unix() > (sms.CreatedAt+60*30) {
		return 0, "验证码已过期"
	} else {
		sms.Status = 1
		ay.Db.Save(sms)
		return 1, "success"
	}
}

// 手机号登入
func (con LoginController) web(phone, code string) (int, string) {

	if res, msg := con.AuthCode(phone, code); res != 1 {
		return 0, msg
	}

	log.Println(phone)
	var user models.User
	ay.Db.First(&user, "phone = ? and type = 0", phone)

	uid := user.Id
	if uid == 0 {
		ss := models.User{
			Phone:    phone,
			Appid:    0,
			Openid:   "",
			NickName: "用户" + strconv.Itoa(rand.Intn(90000)),
			Avatar:   "/static/user/default.png",
			Type:     0,
		}
		ay.Db.Create(&ss)
		// 用户注册 发放优惠卷
		//if vtype == 0 {
		con.SetCoupon(ss.Id)
		//}

		uid = ss.Id
	}
	token := ay.AuthCode(strconv.Itoa(int(uid)), "ENCODE", "", 0)
	return 1, token

}

// SetCoupon 注册发放优惠卷
func (con LoginController) SetCoupon(uid int64) {

	config := models.ConfigModel{}.GetId(1)

	if config.UserRegCoupon == 1 {
		ss := models.Coupon{
			Uid:    uid,
			Status: 0,
			//CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			Name:        config.CouponName,
			SubName:     config.CouponSubName,
			Product:     config.CouponProduct,
			Prohibit:    config.CouponProhibit,
			Effective:   config.CouponEffective,
			EffectiveAt: config.CouponEffectiveAt,
			AmountThan:  config.CouponAmountThan,
			Amount:      config.CouponAmount,
			UsedAt:      "",
		}
		ay.Db.Create(&ss)
	}

}

// 小程序登入
func (con LoginController) xcx(code string, appid int) (int, string, string) {
	var pay models.Pay
	ay.Db.First(&pay, "id = ?", appid)

	res := 0
	openid, session_key := "", ""

	switch pay.Type {
	case 1:

	case 2:
		res, openid, session_key = BaiDuController{}.GetOpenid(code, pay.VKey, pay.Secret)
	}

	if res == 0 {
		return 0, "", openid
	}

	var user models.User
	ay.Db.First(&user, "openid = ? and appid = ?", openid, appid)
	uid := user.Id
	// 用户不存在
	if uid == 0 {
		return 201, openid, session_key
	}

	token := ay.AuthCode(strconv.Itoa(int(uid)), "ENCODE", "", 0)
	return 1, token, ""

}

type GetBindForm struct {
	Appid         int    `form:"appid" binding:"required"`
	SessionKey    string `form:"session_key" binding:"required"`
	Iv            string `form:"iv" binding:"required"`
	Openid        string `form:"openid" binding:"required"`
	EncryptedData string `form:"encryptedData"`
}

// Bind 百度小程序绑定
func (con LoginController) Bind(c *gin.Context) {
	var getForm GetBindForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, "400", ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var pay models.Pay
	ay.Db.First(&pay, "id = ?", getForm.Appid)

	if pay.Id == 0 {
		ay.Json{}.Msg(c, "400", "appid错误", gin.H{})
		return
	}

}

type GetSendForm struct {
	Phone string `form:"phone" binding:"required,len=11"`
}

// Send 发送短信
func (con LoginController) Send(c *gin.Context) {

	var getForm GetSendForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, "400", ay.Validator{}.Translate(err), gin.H{})
		return
	}

	phone := getForm.Phone

	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	if !rgx.MatchString(phone) {
		ay.Json{}.Msg(c, "400", "手机号错误", gin.H{})
		return
	}

	ip := c.ClientIP()

	var count int64
	ay.Db.Model(&models.Sms{}).Where("ip = ? and phone = ? and ymd = ?", ip, phone, time.Now().Format("20060102")).Count(&count)

	if count > 3 {
		ay.Json{}.Msg(c, "400", "短信发送上限", gin.H{})
		return
	}

	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	err := models.SmsModel{}.Send(phone, code)
	if err != nil {
		log.Println(err)
	}
	smss := models.Sms{
		Code:      code,
		Phone:     phone,
		Ip:        ip,
		Ymd:       time.Now().Format("20060102"),
		Status:    0,
		CreatedAt: time.Now().Unix() + 3600,
	}
	ay.Db.Create(&smss)
	if smss.Id != 0 {
		ay.Json{}.Msg(c, "200", "短信发送成功", gin.H{})
	} else {
		ay.Json{}.Msg(c, "400", "短信发送失败", gin.H{})
	}

}
