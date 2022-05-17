/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package api

import (
	"encoding/json"
	"gin/ay"
	"gin/models"
	"gin/models/common"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)

type HaulController struct {
}

type GetHaulSubmitForm struct {
	UserName string `form:"username" binding:"required" label:"名称"`
	Gender   int    `form:"gender" binding:"required" label:"性别"`
	Y        int    `form:"y" binding:"required" label:"年份"`
	M        int    `form:"m" binding:"required" label:"月份"`
	D        int    `form:"d" binding:"required" label:"日"`
	H        int    `form:"h"`
}

// Submit 生成订单
func (con HaulController) Submit(c *gin.Context) {
	var getForm GetHaulSubmitForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if getForm.M > 12 || getForm.M < 1 {
		ay.Json{}.Msg(c, 400, "请输入正确的月份", gin.H{})
		return
	}
	if getForm.D > 31 || getForm.D < 1 {
		ay.Json{}.Msg(c, 400, "请输入正确的天数", gin.H{})
		return
	}

	if getForm.H < -1 || getForm.H > 24 {
		ay.Json{}.Msg(c, 400, "请输入正确的时间", gin.H{})
		return
	}

	line := common.LineModel{}.Line()

	oid := ay.MakeOrder(time.Now())

	// 获取价格
	config := models.ConfigModel{}.GetId(1)

	var user models.User
	ay.Db.First(&user, "id = ?", GetToken(Token))

	if user.Id == 0 {
		ay.Json{}.Msg(c, 401, "Token错误", gin.H{})
		return
	}

	// 返回支付优惠
	coupon := ay.MakeCoupon(config.HaulDiscount)

	des := getForm.UserName + "的2022虎年运程"

	order := &models.Order{
		Oid:        oid,
		Type:       1,
		Ip:         GetRequestIP(c),
		Des:        des,
		Amount:     config.HaulAmount - coupon,
		Uid:        user.Id,
		Status:     0,
		UserName:   getForm.UserName,
		Gender:     getForm.Gender,
		Appid:      Appid,
		PayType:    0,
		OutTradeNo: oid,
		Y:          getForm.Y,
		M:          getForm.M,
		D:          getForm.D,
		H:          getForm.H,
		Discount:   coupon,
		Line:       line,
		Op:         2,
		OldAmount:  config.HaulAmount,
	}

	ay.Db.Create(order)

	if order.Id == 0 {
		ay.Json{}.Msg(c, 400, "数据错误，请联系管理员", gin.H{})
	} else {
		ay.Json{}.Msg(c, 200, "success", gin.H{
			"oid": oid,
		})
	}

}

type GetHaulDetailForm struct {
	Oid string `form:"oid" binding:"required" label:"订单号"`
}

// Detail 详情
func (con HaulController) Detail(c *gin.Context) {
	var getForm GetHaulDetailForm
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

	ay.Db.First(&order, "oid = ? and type = 1", getForm.Oid)

	if order.Id == 0 {
		ay.Json{}.Msg(c, 400, "订单不存在", gin.H{})
		return
	}

	BaZi, bz, hour, res := common.HaulModel{}.Detail(order.Y, order.M, order.D, order.H, order.I, order.S)

	res["info"] = map[string]string{
		"username": order.UserName,
		"y":        strconv.Itoa(order.Y),
		"m":        strconv.Itoa(order.M),
		"d":        strconv.Itoa(order.D),
		"h":        strconv.Itoa(order.H),
		"hour":     hour,
		"gender":   strconv.Itoa(order.Gender),
	}

	// 五行
	res["wuXingWang"] = common.WuXingModel{}.Wang(bz)
	res["wuXingDefect"] = common.WuXingModel{}.Defect(bz)

	// 星
	res["zhengYuanXing"] = common.XingModel{}.ZhengYuan(bz)
	res["taoHuaXing"] = common.XingModel{}.TaoHua(bz)
	res["caiXing"] = common.XingModel{}.Cai(bz)
	res["guiRenXing"] = common.XingModel{}.GuiRen(bz)

	res["pattern"] = common.PatternModel{}.Get(bz[4], bz[3])

	rz := "元女"

	if order.Gender == 1 {
		rz = "元男"
	}
	// 十神
	res["tenGods"] = []string{
		BaZi.GetYearShiShenGan(),
		BaZi.GetMonthShiShenGan(),
		rz,
		BaZi.GetTimeShiShenGan(),
	}

	var line []common.Line
	_ = json.Unmarshal([]byte(order.Line), &line)
	res["line"] = line
	res["coupon"] = order.Discount

	// 获取价格
	config := models.ConfigModel{}.GetId(1)

	var amount []models.HaulAmount
	ay.Db.Order("sort asc").Find(&amount)
	for k, v := range amount {
		amount[k].Amount = v.Amount * config.Rate
	}

	var count int64
	ay.Db.Model(&models.Order{}).Count(&count)
	res["amount"] = amount
	res["return_amount"] = config.HaulAmount
	res["num"] = 10000 + count

	if order.Status != 1 {
		res["isPay"] = false
		ay.Json{}.Msg(c, 200, "未支付", res)
	} else {
		res["isPay"] = true
		ay.Json{}.Msg(c, 200, "success", res)
	}

}

type GetHaulCouponForm struct {
	Id string `form:"id" binding:"required"`
}

// Coupon 优惠卷
func (con HaulController) Coupon(c *gin.Context) {
	var getForm GetHaulCouponForm
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

	var amount models.HaulAmount
	ay.Db.First(&amount, "id = ?", getForm.Id)

	if amount.Id == 0 {
		ay.Json{}.Msg(c, 400, "金额错误", gin.H{})
		return
	}

	var coupon []models.Coupon
	ay.Db.Where("uid = ? and FIND_IN_SET(1,product) and status=0 and amount_than <= ? and effective_at > ?", user.Id, amount.Amount, time.Now().Format("2006-01-02 15:04:05")).Order("id desc").Find(&coupon)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list": coupon,
	})

}

// Main 初始化
func (con HaulController) Main(c *gin.Context) {
	config := models.ConfigModel{}.GetId(1)

	var order []models.Order
	ay.Db.Where("type = 1 and status = 1 and amount > 5").Limit(20).Order("RAND()").Find(&order)

	var notice []string
	for _, v := range order {
		amount := strconv.FormatFloat(v.Amount*config.Rate, 'g', -1, 64)
		minutes := strconv.Itoa(rand.Intn(30))
		str := ""
		if v.Gender == 1 {
			str = string([]rune(v.UserName)[:1]) + "先生" + minutes + "分钟前支付" + amount + "元领取精心打造的命书"
		} else {
			str = string([]rune(v.UserName)[:1]) + "女士" + minutes + "分钟前支付" + amount + "元领取精心打造的命书"
		}
		notice = append(notice, str)
	}

	introduce := models.HaulCasesModel{}.GetType(2)
	cases := models.HaulCasesModel{}.GetType(1)

	for k, v := range introduce {
		introduce[k].Link = ay.Yaml.GetString("domain") + v.Link
		introduce[k].Cover = ay.Yaml.GetString("domain") + v.Cover
	}

	for k, v := range cases {
		cases[k].Link = ay.Yaml.GetString("domain") + v.Link
	}

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"introduce": introduce,
		"cases":     cases,
		"notice":    notice,
	})

}

func (con HaulController) Notice(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	// 获取价格
	config := models.ConfigModel{}.GetId(1)

	var order []models.Order
	ay.Db.Where("type = 1 and status = 1 and amount > 5").Limit(20).Order("RAND()").Find(&order)

	//var notice []N

	var notice []string
	for _, v := range order {
		amount := strconv.FormatFloat(v.Amount*config.Rate, 'g', -1, 64)
		minutes := strconv.Itoa(rand.Intn(30))
		str := ""
		if v.Gender == 1 {
			str = string([]rune(v.UserName)[:1]) + "先生" + minutes + "分钟前支付" + amount + "铜币领取精心打造的命书"
		} else {
			str = string([]rune(v.UserName)[:1]) + "女士" + minutes + "分钟前支付" + amount + "铜币领取精心打造的命书"
		}
		notice = append(notice, str)
	}

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list": &notice,
	})
}
