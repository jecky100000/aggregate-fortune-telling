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
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type AskController struct {
}

type Res struct {
	User   models.User
	Master models.Master
}

func (con AskController) Main(c *gin.Context) {
	var comment []models.UserComment
	ay.Db.Order("RAND()").Limit(10).Find(&comment)

	var com []map[string]interface{}
	for _, v := range comment {

		// 获取用户信息
		var user models.User
		ay.Db.First(&user, v.Uid)
		if user.Id == 0 {
			continue
		}

		type cc struct {
			models.Master
			Avatar   string `json:"avatar"`
			Nickname string `json:"nickname"`
			Phone    string `json:"phone"`
		}
		var re cc
		ay.Db.Table("sm_user").
			Select("sm_user.id,sm_user.phone,sm_user.nickname,sm_user.phone,sm_master.sign,sm_master.type,sm_master.years,sm_master.online,sm_user.avatar,sm_master.rate").
			Joins("left join sm_master on sm_user.master_id=sm_master.id").
			Where("sm_user.id = ?", v.MasterId).
			Order("RAND()").
			Limit(10).
			First(&re)

		if re.Id == 0 {
			continue
		}

		var typeName []string

		for _, v1 := range strings.Split(re.Type, ",") {
			var masterType models.MasterType
			ay.Db.First(&masterType, "id = ?", v1)
			if masterType.Name != "" {
				typeName = append(typeName, masterType.Name)
			}

		}

		com = append(com, map[string]interface{}{
			"user_avatar":   ay.Yaml.GetString("domain") + user.Avatar,
			"user_nickname": user.NickName,
			"created_at":    v.CreatedAt.Format("2006-01-02 15:04:05"),
			"content":       v.Content,

			"master_id":        re.Id,
			"master_name":      re.Nickname,
			"master_phone":     re.Phone,
			"master_sign":      re.Sign,
			"master_avatar":    ay.Yaml.GetString("domain") + re.Avatar,
			"master_type_name": typeName,
			"rate":             v.Rate,
		})
	}
	//fw := []string{
	//	"婚恋情感", "事业财运", "命运详批", "事业指点",
	//}
	var dynamic []string
	var phone string
	type mst struct {
		models.Master
		Nickname string `json:"nickname"`
	}

	//var res []mst
	//
	//ay.Db.Table("sm_user").
	//	Select("sm_user.nickname").
	//	Joins("left join sm_master on sm_user.master_id=sm_master.id").
	//	Where("sm_user.type = 1").
	//	First(&res)

	for i := 0; i < 10; i++ {
		phone = MakePhone()
		//name = res[rand.Intn(len(res))].Nickname
		//dynamic = append(dynamic, phone+" 购买了 "+name+" 的"+fw[rand.Intn(len(fw)-1)]+"服务")
		dynamic = append(dynamic, "尾号"+phone+"用户"+strconv.Itoa(rand.Intn(58)+1)+"分钟前发布"+strconv.Itoa(rand.Intn(200)+1)+"元悬赏找到心仪大师")
	}

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"comment": com,
		"dynamic": dynamic,
	})
}

type GetAskGetForm struct {
	Type string `form:"type" binding:"required" label:"类型"`
}

func (con AskController) Get(c *gin.Context) {
	var getForm GetAskGetForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var log []models.AskLog

	ay.Db.Order("RAND()").Find(&log, "type = ?", getForm.Type)

	var ss []string
	for _, v := range log {
		ss = append(ss, v.Content)
	}

	var masterCount int64
	ay.Db.Model(&models.Master{}).Where("online = 1").Count(&masterCount)

	online := 20 + masterCount
	number := 77 + rand.Intn(9) + rand.Intn(9)
	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list":   ss,
		"online": online,
		"number": number,
	})
}

type GetAskSubmitForm struct {
	UserName  string  `form:"username" binding:"required" label:"名称"`
	Gender    int     `form:"gender" binding:"required" label:"性别"`
	Birth     string  `form:"birth" binding:"required" label:"生日"`
	Type      int     `form:"type"`
	Content   string  `form:"content" binding:"required" label:"内容"`
	Amount    float64 `form:"amount"`
	AreaId    int     `form:"area_id" binding:"required" label:"地区id"`
	PayType   int     `form:"pay_type" binding:"required" label:"支付类型"`
	ReturnUrl string  `form:"return_url" binding:"required" label:"回调地址"`
}

func (con AskController) Submit(c *gin.Context) {
	var getForm GetAskSubmitForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var user models.User
	ay.Db.First(&user, "id = ?", GetToken(Token))

	if user.Id == 0 {
		ay.Json{}.Msg(c, 401, "Token错误", gin.H{})
		return
	}

	amount := getForm.Amount

	if amount != 0 {
		if amount < 5 {
			ay.Json{}.Msg(c, 400, "悬赏金额需要大于5元", gin.H{})
			return
		}

		if getForm.PayType == 3 {
			// 余额支付
			if user.Amount < amount {
				ay.Json{}.Msg(c, 406, "余额不足，请选择其他支付方式", gin.H{})
				return
			}

			user.Amount = user.Amount - amount
			if err := ay.Db.Save(&user).Error; err != nil {
				ay.Json{}.Msg(c, 400, "请联系管理员", gin.H{})
				return
			}
		}
	}

	oid := ay.MakeOrder(time.Now())

	des := strconv.Itoa(getForm.Type)

	order := &models.Order{
		Oid:        oid,
		Type:       3,
		Ip:         GetRequestIP(c),
		Des:        des,
		Amount:     getForm.Amount,
		Uid:        user.Id,
		Status:     -1,
		UserName:   getForm.UserName,
		Gender:     getForm.Gender,
		Appid:      Appid,
		PayType:    0,
		OutTradeNo: oid,
		Op:         2,
		OldAmount:  getForm.Amount,
		Birthday:   getForm.Birth,
		Content:    getForm.Content,
		AreaId:     getForm.AreaId,
		ReturnUrl:  getForm.ReturnUrl,
	}

	ay.Db.Create(order)

	if getForm.PayType != 3 {
		if (getForm.PayType == 1 || getForm.PayType == 2) && Appid == 1 {
			// 支付宝 微信
			v := strconv.FormatFloat(order.Amount, 'g', -1, 64)
			order.Des = "在线提问" + v + "元"
			ay.Db.Save(&order)

			if getForm.PayType == 1 {
				ay.Json{}.Msg(c, 200, "success", gin.H{
					"url": ay.Yaml.GetString("domain") + "/pay/alipay?oid=" + order.Oid,
				})
				return
			} else {
				ay.Json{}.Msg(c, 200, "success", gin.H{
					"url": ay.Yaml.GetString("domain") + "/pay/wechat?oid=" + order.Oid,
				})
				return
			}

		} else if Appid == 2 {
			// 百度支付
			is, msg, rj := BaiDuController{}.Baidu(order.Oid)

			if is {
				ay.Json{}.Msg(c, 200, "success", gin.H{
					"info": rj,
				})
				return
			} else {
				ay.Json{}.Msg(c, 400, msg, gin.H{})
				return
			}
		} else {
			ay.Json{}.Msg(c, 400, "支付类型错误", gin.H{})
			return
		}
	}

	if order.Id == 0 {
		ay.Json{}.Msg(c, 400, "数据错误，请联系管理员", gin.H{})
	} else {
		order.Status = 0
		ay.Db.Save(&order)
		// 上级消费
		models.UserInviteConsumptionModel{}.Set(user.Id, user.Pid, getForm.Amount, oid)

		ay.Json{}.Msg(c, 200, "success", gin.H{
			"oid": oid,
		})
	}

}

func MakePhone() string {
	//top := []string{
	//	"134", "135", "136", "137", "138", "139", "147", "150", "121", "152", "157", "158", "159", "165",
	//	"172", "178", "182", "183", "184", "187", "188", "198",
	//	"130", "131", "132", "145", "155", "156", "166", "171", "175", "176", "185", "186",
	//	"133", "149", "153", "173", "177", "180", "181", "189", "199",
	//}

	//topNum := top[rand.Intn(len(top)-1)]

	s := strconv.Itoa(rand.Intn(9)) + strconv.Itoa(rand.Intn(9)) + strconv.Itoa(rand.Intn(9)) + strconv.Itoa(rand.Intn(9))

	//return topNum + "****" + s
	return s
}

type GetAskDetailForm struct {
	AskId string `form:"ask_id" binding:"required" label:"问题id"`
}

func (con AskController) Detail(c *gin.Context) {
	var getForm GetAskDetailForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var user models.User
	ay.Db.First(&user, "id = ?", GetToken(Token))

	if user.Id == 0 {
		ay.Json{}.Msg(c, 401, "Token错误", gin.H{})
		return
	}

	var order models.Order
	ay.Db.First(&order, "oid = ?", getForm.AskId)

	type S struct {
		ReplyId   string        `json:"reply_id"`
		AskId     string        `json:"ask_id"`
		MasterId  int64         `json:"master_id"`
		Content   string        `json:"content"`
		Adopt     int           `json:"adopt"`
		Name      string        `json:"name"`
		Avatar    string        `json:"avatar"`
		CreatedAt models.MyTime `json:"created_at"`
	}
	var ss []S
	ay.Db.Table("sm_ask_reply").
		Select("sm_ask_reply.id as reply_id,sm_ask_reply.master_id,sm_user.avatar,sm_user.nickname as name,sm_ask_reply.content,sm_ask_reply.created_at,sm_ask_reply.adopt").
		Joins("left join sm_user on sm_user.id=sm_ask_reply.master_id").
		Joins("left join sm_master on sm_user.master_id=sm_master.id").
		Where("sm_ask_reply.ask_id", getForm.AskId).
		Order("sm_ask_reply.id desc").
		Order("sm_ask_reply.created_at desc").
		Find(&ss)

	var count int64
	ay.Db.Model(&models.AskReply{}).Where("ask_id = ?", getForm.AskId).Count(&count)

	var list []interface{}

	adopt := 0
	for _, v := range ss {
		if v.Adopt == 1 {
			adopt = 1
		}
		list = append(list, map[string]interface{}{
			"reply_id":  v.ReplyId,
			"master_id": v.MasterId,
			"name":      v.Name,
			"avatar":    ay.Yaml.GetString("domain") + v.Avatar,
			"adopt":     v.Adopt,
			"content":   v.Content,
			//"created_at": v.CreatedAt.Format("2006/01/02"),
			"created_at": ay.LastTime1(int(v.CreatedAt.Unix())),
		})
	}

	if list == nil {
		list = []interface{}{}
	}

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list":     list,
		"avatar":   ay.Yaml.GetString("domain") + user.Avatar,
		"nickname": user.NickName,
		"amount":   order.Amount,
		"reply":    count,
		"content":  order.Content,
		"type":     order.Des,
		//"created_at": order.CreatedAt.Format("2006/01/02"),
		"created_at": ay.LastTime1(int(order.CreatedAt.Unix())),
		"adopt":      adopt,
	})
}

type GetAskAdoptForm struct {
	ReplyId int64 `form:"reply_id" binding:"required" label:"回复id"`
}

func (con AskController) Adopt(c *gin.Context) {
	var getForm GetAskAdoptForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var user models.User
	ay.Db.First(&user, "id = ?", GetToken(Token))

	if user.Id == 0 {
		ay.Json{}.Msg(c, 401, "Token错误", gin.H{})
		return
	}

	var res models.AskReply
	ay.Db.First(&res, getForm.ReplyId)

	if res.Id == 0 {
		ay.Json{}.Msg(c, 400, "回复不存在", gin.H{})
		return
	}

	if res.Adopt == 1 {
		ay.Json{}.Msg(c, 400, "已采纳无需再次采纳", gin.H{})
		return
	}

	var count int64
	ay.Db.Model(&models.AskReply{}).Where("adopt = 1 AND ask_id = ?", res.AskId).Count(&count)

	if count > 0 {
		ay.Json{}.Msg(c, 400, "已采纳无需再次采纳", gin.H{})
		return
	}

	tx := ay.Db.Begin()

	res.Adopt = 1
	if err := tx.Save(&res).Error; err == nil {

		var order models.Order
		tx.Where("oid = ?", res.AskId).First(&order)
		order.Status = 1
		tx.Save(&order)

		var master models.User
		tx.Where("id = ?", res.MasterId).First(&master)

		master.Amount += order.Amount
		tx.Save(master)

		tx.Commit()
		ay.Json{}.Msg(c, 200, "采纳成功", gin.H{})
	} else {
		tx.Rollback()
		ay.Json{}.Msg(c, 400, "操作失败", gin.H{})

	}
}
