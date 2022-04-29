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
	"gin/sdk/tencentyun"
	"github.com/6tail/lunar-go/calendar"
	"github.com/gin-gonic/gin"
	"log"
	"path"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
}

type GetUserEditForm struct {
	NickName string `form:"nickname"`
	BirthDay string `form:"birthday"`
	AreaId   int    `form:"area_id"`
	Gender   int    `form:"gender"`
	Avatar   string `form:"avatar"`
}

// Edit 修改信息
func (con UserController) Edit(c *gin.Context) {
	var getForm GetUserEditForm
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

	if getForm.NickName != "" {
		user.NickName = getForm.NickName
	}

	if getForm.Avatar != "" {
		user.Avatar = getForm.Avatar
	}

	if getForm.AreaId != 0 {
		user.AreaId = getForm.AreaId
	}

	if getForm.Gender != 0 {
		user.Gender = getForm.Gender
	}

	if getForm.BirthDay != "" {
		user.BirthDay = getForm.BirthDay
	}

	if err := ay.Db.Save(&user).Error; err != nil {
		ay.Json{}.Msg(c, 400, "修改失败", gin.H{})
	} else {
		ay.Json{}.Msg(c, 200, "修改成功", gin.H{
			"avatar":   ay.Yaml.GetString("domain") + user.Avatar,
			"phone":    user.Phone,
			"nickname": user.NickName,
		})
	}

}

type GetUserCouponForm struct {
	Page int `form:"page" binding:"required" label:"页码"`
}

// Coupon 优惠卷
func (con UserController) Coupon(c *gin.Context) {
	var getForm GetUserCouponForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	page := getForm.Page - 1

	var user models.User
	ay.Db.First(&user, "id = ?", GetToken(Token))

	if user.Id == 0 {
		ay.Json{}.Msg(c, 401, "Token错误", gin.H{})
		return
	}

	var coupon []models.Coupon
	ay.Db.Where("uid = ?", user.Id).Limit(10).Offset(page * 10).Find(&coupon)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list": coupon,
	})
}

// Info 用户信息
func (con UserController) Info(c *gin.Context) {

	config := models.ConfigModel{}.GetId(1)

	var user models.User
	ay.Db.First(&user, "id = ?", GetToken(Token))

	if user.Id == 0 {
		ay.Json{}.Msg(c, 401, "Token错误", gin.H{})
		return
	}

	var frozenInviteAmount float64
	ay.Db.Table("sm_user_invite_consumption").Where("pid = ? AND status=0", user.Id).Pluck("SUM(amount)", &frozenInviteAmount)

	UserSig, _ := tencentyun.GenUserSig(ay.Yaml.GetInt("im.appid"), ay.Yaml.GetString("im.key"), user.Phone, 3600*24)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"nickname":             user.NickName,
		"avatar":               ay.Yaml.GetString("domain") + user.Avatar,
		"gender":               user.Gender,
		"area_id":              user.AreaId,
		"birthday":             user.BirthDay,
		"phone":                user.Phone,
		"amount":               user.Amount * config.Rate,
		"frozen_invite_amount": frozenInviteAmount,
		"invite_amount":        user.InviteAmount * config.Rate,
		"aff":                  user.Aff,
		"UserSig":              UserSig,
	})
}

// Upload 上传头像
func (con UserController) Upload(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		ay.Json{}.Msg(c, 400, "上传图片出错", gin.H{})
		return
	}
	//log.Println(file.Filename)

	var user models.User
	ay.Db.First(&user, "id = ?", GetToken(Token))

	if user.Id == 0 {
		ay.Json{}.Msg(c, 401, "Token错误", gin.H{})
		return
	}

	fileExt := strings.ToLower(path.Ext(c.PostForm("filename")))
	//fileExt := strings.ToLower(path.Ext(file.Filename))

	log.Println(fileExt)
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
		ay.Json{}.Msg(c, 400, "上传失败!只允许png,jpg,gif,jpeg文件", gin.H{})
		return
	}
	fileName := ay.MD5(fmt.Sprintf("%s%s", file.Filename, time.Now().String()))
	fildDir := fmt.Sprintf("static/upload/user/%d-%d/", time.Now().Year(), time.Now().Month())

	err = ay.CreateMutiDir(fildDir)
	if err != nil {
		log.Println(err)
	}

	filepath := fmt.Sprintf("%s%s%s", fildDir, fileName, fileExt)
	c.SaveUploadedFile(file, filepath)
	ay.Json{}.Msg(c, 200, "上传成功!", gin.H{
		"url": "/" + filepath,
	})
}

type GetUserCollectForm struct {
	Type string `form:"type" binding:"required"`
	Page int    `form:"page" binding:"required"`
}

type returnCollect struct {
	Id        int64    `json:"id"`
	Cid       int64    `json:"cid"`
	Title     string   `json:"title"`
	Cover     string   `json:"cover"`
	Name      string   `json:"name"`
	Type      int      `json:"type"`
	Rate      float64  `json:"rate"`
	Avatar    string   `json:"avatar"`
	Online    int      `json:"online"`
	Sign      string   `json:"sign"`
	TypeName  []string `json:"type_name"`
	Label     string   `json:"label"`
	Years     int      `json:"years"`
	Collect   int      `json:"collect"`
	CreatedAt string   `json:"created_at"`
	View      int64    `json:"view"`
}

// Collect 收藏
func (con UserController) Collect(c *gin.Context) {

	var getForm GetUserCollectForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	page := getForm.Page - 1

	var user models.User
	ay.Db.First(&user, "id = ?", GetToken(Token))

	if user.Id == 0 {
		ay.Json{}.Msg(c, 401, "Token错误", gin.H{})
		return
	}

	var collect []models.Collect
	ay.Db.Where("uid = ? and type = ?", user.Id, getForm.Type).Limit(10).Offset(page * 10).Find(&collect)

	var guji []returnCollect

	for _, v := range collect {
		if v.Type == 2 {
			var baike models.BaiKe
			ay.Db.First(&baike, "id = ?", v.Cid)
			guji = append(guji, returnCollect{
				Cid:       v.Cid,
				Id:        v.Id,
				Title:     baike.Title,
				Cover:     baike.Cover,
				Type:      v.Type,
				CreatedAt: ay.LastTime(int(baike.CreatedAt.Unix())),
				View:      baike.View,
				Collect:   1,
			})
		} else if v.Type == 3 {
			var g models.Ancient
			ay.Db.First(&g, "id = ?", v.Cid)
			guji = append(guji, returnCollect{
				Cid:       v.Cid,
				Id:        v.Id,
				Title:     g.Title,
				Cover:     g.Cover,
				Type:      v.Type,
				Collect:   1,
				CreatedAt: ay.LastTime(int(g.CreatedAt.Unix())),
				View:      g.View,
			})
		} else if v.Type == 1 {

			type cc struct {
				models.Master
				Avatar string `json:"avatar"`
			}

			var d cc

			ay.Db.Table("sm_user").
				Select("sm_master.name,sm_master.sign,sm_master.type,sm_master.years,sm_master.online,sm_user.avatar,sm_master.rate,sm_master.label").
				Joins("left join sm_master on sm_user.master_id=sm_master.id").
				Where("sm_user.id", v.Cid).
				First(&d)

			var type_name []string

			for _, v := range strings.Split(d.Type, ",") {
				var master_type models.MasterType
				ay.Db.First(&master_type, "id = ?", v)
				if master_type.Name != "" {
					type_name = append(type_name, master_type.Name)
				}

			}

			guji = append(guji, returnCollect{
				Cid:      v.Cid,
				Id:       v.Id,
				Name:     d.Name,
				Avatar:   ay.Yaml.GetString("domain") + d.Avatar,
				Type:     v.Type,
				Rate:     d.Rate,
				TypeName: type_name,
				Sign:     d.Sign,
				Label:    d.Label,
				Years:    d.Years,
				Collect:  1,
			})
		}
	}

	if guji == nil {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": []string{},
		})
	} else {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": guji,
		})
	}

}

type GetUserHistoryForm struct {
	Type int `form:"type" binding:"required" label:"类型"`
	Page int `form:"page" binding:"required" label:"页码"`
}

type ReturnHistory struct {
	Type      int     `json:"type"`
	CreatedAt string  `json:"created_at"`
	Op        int     `json:"op"`
	Amount    float64 `json:"amount"`
	Status    int     `json:"status"`
}

// History 历史订单
func (con UserController) History(c *gin.Context) {
	var getForm GetUserHistoryForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	page := getForm.Page - 1

	var user models.User
	ay.Db.First(&user, "id = ?", GetToken(Token))

	if user.Id == 0 {
		ay.Json{}.Msg(c, 401, "Token错误", gin.H{})
		return
	}

	var order []models.Order

	if getForm.Type == 1 {
		ay.Db.Where("uid = ? and type != 8 and type != 9", user.Id).Order("id desc").Limit(10).Offset(page * 10).Find(&order)
	} else if getForm.Type == 2 {
		ay.Db.Where("uid = ? and type = 9 and status = 1", user.Id).Order("id desc").Limit(10).Offset(page * 10).Find(&order)
	} else {
		ay.Db.Where("uid = ? and type = 8", user.Id).Order("id desc").Limit(10).Offset(page * 10).Find(&order)
	}

	var history []ReturnHistory

	config := models.ConfigModel{}.GetId(1)

	for _, v := range order {
		loc, _ := time.LoadLocation("Local")
		theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", v.CreatedAt.String()[:19], loc)
		history = append(history, ReturnHistory{
			Type:      v.Type,
			Op:        v.Op,
			Amount:    v.Amount * config.Rate,
			CreatedAt: time.Unix(theTime.Unix(), 0).Format("2006/1/2"),
			Status:    v.Status,
		})
	}

	if history == nil {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list":   []string{},
			"amount": user.Amount,
		})
	} else {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list":   history,
			"amount": user.Amount,
		})
	}

}

type GetUserControllerWithdrawal struct {
	Type    int     `form:"type" binding:"required" label:"类型"`
	Amount  float64 `form:"amount" binding:"required" label:"金额"`
	Account string  `form:"account" binding:"required" label:"账号"`
}

// Withdrawal 提现
func (con UserController) Withdrawal(c *gin.Context) {

	var getForm GetUserControllerWithdrawal
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

	if user.Amount < getForm.Amount {
		ay.Json{}.Msg(c, 400, "余额不足", gin.H{})
		return
	}

	user.Amount = user.Amount - getForm.Amount
	if err := ay.Db.Save(&user).Error; err != nil {
		ay.Json{}.Msg(c, 400, "提现失败", gin.H{})
		return
	}

	oid := ay.MakeOrder(time.Now())
	v := strconv.FormatFloat(getForm.Amount, 'g', -1, 64)
	order := &models.Order{
		Oid:        oid,
		Type:       8,
		Ip:         GetRequestIP(c),
		Des:        "提现" + v + "元",
		Amount:     getForm.Amount,
		Uid:        user.Id,
		Status:     0,
		Appid:      Appid,
		PayType:    getForm.Type,
		OutTradeNo: oid,
		Op:         2,
		OldAmount:  getForm.Amount,
		Json:       getForm.Account,
	}

	ay.Db.Create(order)

	if order.Id == 0 {
		ay.Json{}.Msg(c, 400, "提现失败", gin.H{})
	} else {
		ay.Json{}.Msg(c, 200, "等待确认", gin.H{})
	}

}

type GetUserControllerLog struct {
	Type int `form:"type" binding:"required" label:"类型"`
	Page int `form:"page" binding:"required" label:"页码"`
}

type Pp struct {
	Type      int    `json:"type"`
	Oid       string `json:"oid"`
	UserName  string `json:"username"`
	Gender    int    `json:"gender"`
	Y         int    `json:"y"`
	M         int    `json:"m"`
	D         int    `json:"d"`
	CreatedAt string `json:"created_at"`
	Animal    string `json:"animal"`
	XingZuo   string `json:"xingZuo"`
	Status    int    `json:"status"`
}

// Log 历史记录
func (con UserController) Log(c *gin.Context) {
	var getForm GetUserControllerLog
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

	page := getForm.Page - 1

	if getForm.Type == 1 {
		var row []map[string]interface{}
		var usermasterlog []models.UserMasterLog
		ay.Db.Order("id desc").Limit(10).Offset(page*10).Find(&usermasterlog, "uid = ?", user.Id)
		for _, v1 := range usermasterlog {
			type cc struct {
				models.Master
				Avatar string `json:"avatar"`
			}
			var res []cc

			ay.Db.Table("sm_master").
				Select("sm_master.*,sm_user.avatar").
				Joins("left join sm_user on sm_master.id=sm_user.master_id").
				Where("sm_master.id = ?", v1.MasterId).
				Find(&res)

			for _, v := range res {
				var type_name []string

				for _, v := range strings.Split(v.Type, ",") {
					var masterType models.MasterType
					ay.Db.First(&masterType, "id = ?", v)
					if masterType.Name != "" {
						type_name = append(type_name, masterType.Name)
					}
				}

				row = append(row, map[string]interface{}{
					"id":        v.Id,
					"name":      v.Name,
					"sign":      v.Sign,
					"years":     v.Years,
					"online":    v.Online,
					"avatar":    ay.Yaml.GetString("domain") + v.Avatar,
					"rate":      v.Rate,
					"type_name": type_name,
				})
			}
		}
		if row == nil {
			ay.Json{}.Msg(c, 200, "success", gin.H{
				"list": []string{},
			})
		} else {
			ay.Json{}.Msg(c, 200, "success", gin.H{
				"list": row,
			})
		}

	} else if getForm.Type == 2 {

		var order []models.Order
		ay.Db.Limit(10).Offset(page*10).Where("uid = ? and type = 1", user.Id).Order("id desc").Find(&order)

		var pp []Pp
		for _, v := range order {
			pp = append(pp, Pp{
				Type:      2,
				Oid:       v.Oid,
				UserName:  v.UserName,
				Gender:    v.Gender,
				Y:         v.Y,
				M:         v.M,
				D:         v.D,
				CreatedAt: v.CreatedAt.Format("2006-01-02 15:04"),
				Status:    v.Status,
			})
		}
		if pp == nil {
			ay.Json{}.Msg(c, 200, "success", gin.H{
				"list": []string{},
			})
		} else {
			ay.Json{}.Msg(c, 200, "success", gin.H{
				"list": pp,
			})
		}

	} else {

		var order []models.Order
		ay.Db.Limit(10).Offset(page*10).Where("uid = ? and type = 2", user.Id).Order("id desc").Find(&order)

		var pp []Pp
		for _, v := range order {
			solar := calendar.NewSolar(v.Y, v.M, v.D, v.H, v.I, 0)
			lunar := solar.GetLunar()
			pp = append(pp, Pp{
				Oid:       v.Oid,
				UserName:  v.UserName,
				Gender:    v.Gender,
				Y:         v.Y,
				M:         v.M,
				D:         v.D,
				Type:      3,
				CreatedAt: v.CreatedAt.Format("2006-01-02 15:04"),
				Animal:    lunar.GetYearShengXiaoByLiChun(),
				XingZuo:   fmt.Sprintf("%s", solar.GetXingZuo()),
			})
		}
		if pp == nil {
			ay.Json{}.Msg(c, 200, "success", gin.H{
				"list": []string{},
			})
		} else {
			ay.Json{}.Msg(c, 200, "success", gin.H{
				"list": pp,
			})
		}
	}

}

type GetUserAsk struct {
	Page int `form:"page"`
}

// Ask 提问记录
func (con UserController) Ask(c *gin.Context) {
	var getForm GetUserAsk
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

	var order []models.Order
	ay.Db.Where("type = 3 and uid = ?", user.Id).Limit(10).Offset((getForm.Page) * 10).Find(&order)

	var res []map[string]interface{}

	for _, v := range order {

		var count int64
		ay.Db.Model(&models.AskReply{}).Where("ask_id", v.Oid).Count(&count)
		res = append(res, map[string]interface{}{
			"ask_id":   v.Oid,
			"nickname": user.NickName,
			"avatar":   ay.Yaml.GetString("domain") + user.Avatar,
			"type":     v.Des,
			"content":  v.Json,
			"status":   v.Status,
			"reply":    count,
			"amount":   v.Amount,
		})
	}

	if res == nil {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": []string{},
		})
	} else {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"list": res,
		})
	}
}
