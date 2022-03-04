/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package api

import (
	"gin/ay"
	"gin/models"
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

		var re models.Master
		ay.Db.Table("sm_user").
			Select("sm_user.id,sm_master.name,sm_master.sign,sm_master.type,sm_master.years,sm_master.online,sm_master.avatar,sm_master.rate").
			Joins("left join sm_master on sm_user.master_id=sm_master.id").
			Where("sm_user.id = ?", v.MasterId).
			Order("RAND()").
			Limit(10).
			First(&re)

		if re.Id == 0 {
			continue
		}

		var type_name []string

		for _, v1 := range strings.Split(re.Type, ",") {
			var master_type models.MasterType
			ay.Db.First(&master_type, "id = ?", v1)
			if master_type.Name != "" {
				type_name = append(type_name, master_type.Name)
			}

		}

		com = append(com, map[string]interface{}{
			"user_avatar":   ay.Domain + user.Avatar,
			"user_nickname": user.NickName,
			"created_at":    v.CreatedAt.Format("2006-01-02 15:04:05"),
			"content":       v.Content,

			"master_id":        re.Id,
			"master_name":      re.Name,
			"master_sign":      re.Sign,
			"master_avatar":    ay.Domain + re.Avatar,
			"master_type_name": type_name,
		})
	}
	fw := []string{
		"婚恋情感", "事业财运", "命运详批", "事业指点",
	}
	var dynamic []map[string]string
	var phone string
	var name string

	var res []models.Master

	ay.Db.Table("sm_user").
		Select("sm_master.name").
		Joins("left join sm_master on sm_user.master_id=sm_master.id").
		Where("sm_user.type = 1").
		First(&res)

	for i := 0; i < 10; i++ {
		phone = MakePhone()
		name = res[rand.Intn(len(res))].Name
		//dynamic = append(dynamic, phone+" 购买了 "+name+" 的"+fw[rand.Intn(len(fw)-1)]+"服务")
		dynamic = append(dynamic, map[string]string{
			"phone": phone,
			"name":  name,
			"info":  fw[rand.Intn(len(fw))],
		})
	}

	ay.Json{}.Msg(c, "200", "success", gin.H{
		"comment": com,
		"dynamic": dynamic,
	})
}

type GetAskGetForm struct {
	Type string `form:"type" binding:"required"`
}

func (con AskController) Get(c *gin.Context) {
	var getForm GetAskGetForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, "400", ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var log []models.AskLog

	ay.Db.Find(&log, "type = ?", getForm.Type)

	var ss []string
	for _, v := range log {
		ss = append(ss, v.Content)
	}

	var masterCount int64
	ay.Db.Model(&models.Master{}).Where("online = 1").Count(&masterCount)

	online := 20 + masterCount
	number := 77 + rand.Intn(9) + rand.Intn(9)
	ay.Json{}.Msg(c, "200", "success", gin.H{
		"list":   ss,
		"online": online,
		"number": number,
	})
}

type GetAskSubmitForm struct {
	UserName string  `form:"username" binding:"required"`
	Gender   int     `form:"gender" binding:"required"`
	Birth    string  `form:"birth" binding:"required"`
	Type     int     `form:"type"`
	Content  string  `form:"content" binding:"required"`
	Amount   float64 `form:"amount"`
	AreAId   int     `form:"area_id"`
}

func (con AskController) Submit(c *gin.Context) {
	var getForm GetAskSubmitForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, "400", ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var user models.User
	ay.Db.First(&user, "id = ?", GetToken(Token))

	if user.Id == 0 {
		ay.Json{}.Msg(c, "401", "Token错误", gin.H{})
		return
	}

	amount := getForm.Amount

	if amount != 0 {
		if amount < 5 {
			ay.Json{}.Msg(c, "400", "悬赏金额需要大于5元", gin.H{})
			return
		}

		if user.Amount < amount {
			ay.Json{}.Msg(c, "400", "余额不足", gin.H{})
			return
		}

		user.Amount = user.Amount - amount
		ay.Db.Save(&user)
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
		Status:     0,
		UserName:   getForm.UserName,
		Gender:     getForm.Gender,
		Appid:      Appid,
		PayType:    0,
		OutTradeNo: oid,
		Op:         2,
		OldAmount:  getForm.Amount,
		Birthday:   getForm.Birth,
		Json:       getForm.Content,
		AreaId:     getForm.AreAId,
	}

	ay.Db.Create(order)

	if order.Id == 0 {
		ay.Json{}.Msg(c, "400", "数据错误，请联系管理员", gin.H{})
	} else {
		ay.Json{}.Msg(c, "200", "success", gin.H{
			"oid": oid,
		})
	}

}

func MakePhone() string {
	top := []string{
		"134", "135", "136", "137", "138", "139", "147", "150", "121", "152", "157", "158", "159", "165",
		"172", "178", "182", "183", "184", "187", "188", "198",
		"130", "131", "132", "145", "155", "156", "166", "171", "175", "176", "185", "186",
		"133", "149", "153", "173", "177", "180", "181", "189", "199",
	}

	topNum := top[rand.Intn(len(top)-1)]

	s := strconv.Itoa(rand.Intn(9)) + strconv.Itoa(rand.Intn(9)) + strconv.Itoa(rand.Intn(9)) + strconv.Itoa(rand.Intn(9))

	return topNum + "****" + s
}

type GetAskDetailForm struct {
	AskId string `form:"ask_id" binding:"required"`
}

func (con AskController) Detail(c *gin.Context) {
	var getForm GetAskDetailForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, "400", ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var user models.User
	ay.Db.First(&user, "id = ?", GetToken(Token))

	if user.Id == 0 {
		ay.Json{}.Msg(c, "401", "Token错误", gin.H{})
		return
	}

	var order models.Order
	ay.Db.First(&order, "oid = ?", getForm.AskId)

	type S struct {
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
		Select("sm_ask_reply.master_id,sm_master.avatar,sm_master.name,sm_ask_reply.content,sm_ask_reply.created_at,sm_ask_reply.adopt").
		Joins("left join sm_user on sm_user.id=sm_ask_reply.master_id").
		Joins("left join sm_master on sm_user.master_id=sm_master.id").
		Where("sm_ask_reply.ask_id", getForm.AskId).
		Order("sm_ask_reply.id desc").
		Find(&ss)

	var count int64
	ay.Db.Model(&models.AskReply{}).Where("ask_id = ?", getForm.AskId).Count(&count)

	var list []interface{}

	for _, v := range ss {
		list = append(list, map[string]interface{}{
			"master_id":  v.MasterId,
			"name":       v.Name,
			"avatar":     ay.Domain + v.Avatar,
			"adopt":      v.Adopt,
			"content":    v.Content,
			"created_at": v.CreatedAt.Format("2006/01/02 15:04"),
		})
	}

	if list == nil {
		list = []interface{}{}
	}

	ay.Json{}.Msg(c, "200", "success", gin.H{
		"list":       list,
		"avatar":     ay.Domain + user.Avatar,
		"nickname":   user.NickName,
		"amount":     order.Amount,
		"reply":      count,
		"type":       order.Des,
		"created_at": order.CreatedAt.Format("2006/01/02"),
	})
}
