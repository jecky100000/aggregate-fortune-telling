/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package api

import (
	"encoding/json"
	"fmt"
	"gin/ay"
	"gin/models"
	"gin/models/login"
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
	Aff    string `form:"aff"`
}

// Login 登入
func (con LoginController) Login(c *gin.Context) {
	var getForm GetLoginForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	res := 0
	token, session_key := "", ""

	if Appid == 1 {
		res, token = con.web(getForm.Phone, getForm.Code, getForm.Aff)
	} else {
		res, token, session_key = con.xcx(getForm.Code, Appid)
		//res, token, session_key = 1, "", ""
		if res == 0 {
			token = session_key
		}
	}

	if res == 0 {
		ay.Json{}.Msg(c, 400, token, gin.H{})
	} else if res == 201 {
		ay.Json{}.Msg(c, 201, "请绑定手机", gin.H{
			"session_key": session_key,
			"openid":      token,
		})
		return
	} else {
		var user models.User
		ay.Db.First(&user, "phone = ?", getForm.Phone)

		ay.Json{}.Msg(c, 200, "success", gin.H{
			"token":    token,
			"avatar":   ay.Yaml.GetString("domain") + user.Avatar,
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
func (con LoginController) web(phone, code, aff string) (int, string) {

	if res, msg := con.AuthCode(phone, code); res != 1 {
		return 0, msg
	}

	var pid int64 = 0
	var pUser models.User

	if aff == "" {
		pid = 0
	} else {
		ay.Db.First(&pUser, "aff = ? and type = 0", aff)
		if pUser.Id != 0 {
			pid = pUser.Id
		} else {
			pid = 0
		}
	}

	var user models.User
	ay.Db.First(&user, "phone = ? and type = 0", phone)

	uid := user.Id
	if uid == 0 {

		ss := models.User{
			Phone:    phone,
			NickName: "用户" + strconv.Itoa(rand.Intn(90000)),
			Avatar:   "/static/user/default.png",
			Type:     0,
			Pid:      pid,
			Aff:      con.GetCode(),
		}
		if err := ay.Db.Create(&ss).Error; err != nil {
			return 0, "注册失败"
		}
		//
		if pid != 0 {
			ay.Db.Create(&models.UserInvite{
				Pid: pid,
				Uid: ss.Id,
			})
		}

		// 用户注册 发放优惠卷
		//if vtype == 0 {
		con.SetCoupon(ss.Id)
		//}
		// 生成im账号
		r, rs := con.MakeImAccount(phone)
		if r != 1 {
			return 0, rs
		}

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
func (con LoginController) xcx(code string, appid int64) (int, string, string) {
	var pay models.Pay
	ay.Db.First(&pay, "id = ?", appid)

	res := 0
	openid, session_key := "", ""

	switch pay.Type {
	case 1:

	case 3:
		res, openid, session_key = BaiDuController{}.GetOpenid(code, pay.VKey, pay.Secret)
		//res, openid, session_key = 1, "456", ""
	}

	if res == 0 {
		return 0, "", openid
	}

	var user models.UserOpenid
	ay.Db.First(&user, "openid = ? AND appid = ?", openid, Appid)
	uid := user.Id
	// 用户不存在
	if uid == 0 {
		return 201, openid, session_key
	}

	token := ay.AuthCode(strconv.Itoa(int(uid)), "ENCODE", "", 0)
	return 1, token, ""

}

type GetBindForm struct {
	SessionKey    string `form:"session_key" binding:"required"`
	Iv            string `form:"iv" binding:"required"`
	Openid        string `form:"openid" binding:"required"`
	EncryptedData string `form:"encryptedData"`
}

// BaiduBind 百度小程序绑定
func (con LoginController) BaiduBind(c *gin.Context) {
	var getForm GetBindForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var pay models.Pay
	ay.Db.First(&pay, "id = ?", Appid)

	if pay.Id == 0 {
		ay.Json{}.Msg(c, 400, "appid错误", gin.H{})
		return
	}

	// 解密
	crypt := login.BdBizDataCrypt{
		AppID:      pay.Appid,
		SessionKey: getForm.SessionKey,
	}
	res, err := crypt.Decrypt(getForm.EncryptedData, getForm.Iv, true)

	if err != nil {
		ay.Json{}.Msg(c, 400, "数据错误", gin.H{})
		return
	}
	type jxm struct {
		Mobile string `json:"mobile"`
	}
	var j jxm
	json.Unmarshal([]byte(res.(string)), &j)

	// 查询之前绑定的openid
	userOpenid := models.UserOpenidModel{}.Get(Appid, getForm.Openid)

	var id int64

	if userOpenid.Id == 0 {
		// 快捷登入不存在
		user := models.UserModel{}.GetPhone(j.Mobile)

		if user.Id == 0 {
			// 用户不存在
			ss := models.User{
				NickName: "用户" + strconv.Itoa(rand.Intn(90000)),
				Avatar:   "/static/user/default.png",
				Type:     0,
				Aff:      con.GetCode(),
				Phone:    j.Mobile,
			}
			if err := ay.Db.Create(&ss).Error; err != nil {
				ay.Json{}.Msg(c, 400, "注册失败", gin.H{})
			}
			sv := &models.UserOpenid{
				Appid:  Appid,
				Uid:    ss.Id,
				Openid: getForm.Openid,
			}
			if err := ay.Db.Create(&sv).Error; err != nil {
				ay.Json{}.Msg(c, 400, "注册失败", gin.H{})
			}
			// 用户注册 发放优惠卷
			con.SetCoupon(ss.Id)
			// 生成im账号
			r, rs := con.MakeImAccount(j.Mobile)
			if r != 1 {
				ay.Json{}.Msg(c, 400, rs, gin.H{})
			}
			id = ss.Id
		} else {
			sv := &models.UserOpenid{
				Appid:  Appid,
				Uid:    user.Id,
				Openid: getForm.Openid,
			}
			if err := ay.Db.Create(&sv).Error; err != nil {
				ay.Json{}.Msg(c, 400, "注册失败", gin.H{})
			}
			id = user.Id
		}
	} else {
		id = userOpenid.Uid
	}

	token := ay.AuthCode(strconv.Itoa(int(id)), "ENCODE", "", 0)
	ay.Json{}.Msg(c, 200, "success", gin.H{
		"token": token,
	})

}

func (con LoginController) GetCode() string {
	makeCode := ay.GetRandomString(6)
	for {
		var affUser models.User
		ay.Db.First(&affUser, "aff = ? and type = 0", makeCode)
		if affUser.Id == 0 {
			break
		} else {
			makeCode = ay.GetRandomString(6)
		}
	}
	return makeCode
}

type GetSendForm struct {
	Phone string `form:"phone" binding:"required,len=11"`
}

// Send 发送短信
func (con LoginController) Send(c *gin.Context) {

	var getForm GetSendForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	phone := getForm.Phone

	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	if !rgx.MatchString(phone) {
		ay.Json{}.Msg(c, 400, "手机号错误", gin.H{})
		return
	}

	ip := c.ClientIP()

	var count int64
	ay.Db.Model(&models.Sms{}).Where("ip = ? and phone = ? and ymd = ?", ip, phone, time.Now().Format("20060102")).Count(&count)

	if count > ay.Yaml.GetInt64("sms.count") {
		ay.Json{}.Msg(c, 400, "短信发送上限", gin.H{})
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
		ay.Json{}.Msg(c, 200, "短信发送成功", gin.H{})
	} else {
		ay.Json{}.Msg(c, 400, "短信发送失败", gin.H{})
	}

}

func (con LoginController) MakeImAccount(phone string) (int, string) {
	imCode, imMsg := models.ImModel{}.HttpPost("/v4/im_open_login_svc/account_import", phone)

	if imCode != 1 {
		return 0, imMsg
	}
	type rj struct {
		ActionStatus string `json:"ActionStatus"`
		ErrorCode    int    `json:"ErrorCode"`
		ErrorInfo    string `json:"ErrorInfo"`
	}
	var r rj
	json.Unmarshal([]byte(imMsg), &r)
	if r.ErrorCode != 0 {
		return 0, r.ErrorInfo
	} else {
		return 1, "success"
	}
}
