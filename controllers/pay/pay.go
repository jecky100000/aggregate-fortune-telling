/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package pay

import (
	"aggregate-fortune-telling/ay"
	"aggregate-fortune-telling/controllers/api"
	"aggregate-fortune-telling/models"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/wechat"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type PayController struct {
}

// AliPay 支付宝
func (con PayController) AliPay(c *gin.Context) {

	log.Println(c.GetHeader("Referer"))
	var order models.Order
	ay.Db.First(&order, "oid = ?", c.Query("oid"))
	if order.Id == 0 {
		c.String(200, "订单错误")
		return
	}

	returnUrl := c.Query("return_url")
	if returnUrl != "" {
		returnUrl, _ = url.QueryUnescape(returnUrl)
		order.ReturnUrl = returnUrl
	} else {
		returnUrl = order.ReturnUrl
	}

	order.OutTradeNo = ay.MakeOrder(time.Now())
	if err := ay.Db.Save(&order).Error; err != nil {
		c.String(200, "请联系管理员")
		return
	}

	code, msg := con.Web(order.OutTradeNo, 1, order.Amount, returnUrl, api.GetRequestIP(c), order.Des)

	if code == 1 {
		models.AdvertLogModel{}.Add(1, order.Oid, c.GetHeader("Referer"), order.Amount, c.Query("request_id"), c.Query("ad_id"))
		c.Redirect(http.StatusMovedPermanently, msg)
		return
	} else {
		ay.Json{}.Msg(c, 400, msg, gin.H{})
		return
	}
}

type OpenId struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
}

// GetOpenid jsapi获取openid
func (con PayController) GetOpenid(c *gin.Context) {

	oid := c.Query("oid")
	code := c.Query("code")

	if oid == "" {
		c.String(200, "订单号不能为空")
		return
	}

	if code == "" {
		c.String(200, "code不能为空")
		return
	}
	var order models.Order
	ay.Db.First(&order, "oid = ?", oid)

	//config := models.ConfigModel{}.GetId(1)

	var pay models.Pay
	ay.Db.First(&pay, "id = ?", 6)

	oauth2Url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + pay.Appid + "&secret=" + pay.Secret + "&code=" + code + "&grant_type=authorization_code"

	resp, _ := http.Get(oauth2Url)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var j OpenId
	err := json.Unmarshal(body, &j)
	if err != nil {
		ay.Json{}.Msg(c, 400, "获取openid失败", gin.H{})
		return
	}

	openid := j.Openid
	log.Println(j)

	ctx, _ := context.WithCancel(context.Background())

	client := wechat.NewClient(pay.Appid, pay.MchId, pay.VKey, true)
	// 打开Debug开关，输出请求日志，默认关闭
	client.DebugSwitch = gopay.DebugOn
	client.SetCountry(wechat.China)

	bm := make(gopay.BodyMap)
	//v := strconv.FormatFloat(order.Amount*config.Rate, "g", -1, 64)
	//v := order.Amount * config.Rate
	bm.Set("nonce_str", util.RandomString(32)).
		Set("body", order.Des).
		Set("out_trade_no", order.OutTradeNo).
		Set("total_fee", order.Amount*100).
		Set("spbill_create_ip", api.GetRequestIP(c)).
		Set("notify_url", ay.Yaml.GetString("domain")+"/api/notify/wechat").
		Set("trade_type", "JSAPI").
		Set("sign_type", "MD5").
		Set("openid", openid)

	wxRsp, err := client.UnifiedOrder(ctx, bm)
	if err != nil {
		c.String(200, "请联系管理员")
		return
	}

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	packages := "prepay_id=" + wxRsp.PrepayId // 此处的 wxRsp.PrepayId ,统一下单成功后得到
	paySign := wechat.GetH5PaySign(pay.Appid, wxRsp.NonceStr, packages, wechat.SignType_MD5, timeStamp, pay.VKey)

	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, `
<script type="text/javascript">
function onBridgeReady(){
   WeixinJSBridge.invoke(
      "getBrandWCPayRequest", {
         "appId":"`+wxRsp.Appid+`",     //公众号名称，由商户传入     
         "timeStamp":"`+timeStamp+`",         //时间戳，自1970年以来的秒数     
         "nonceStr":"`+wxRsp.NonceStr+`", //随机串     
         "package":"`+packages+`",     
         "signType":"MD5",         //微信签名方式：     
         "paySign":"`+paySign+`" //微信签名 
      },
      function(res){
      if(res.err_msg == "get_brand_wcpay_request:ok" ){
      // 使用以上方式判断前端返回,微信团队郑重提示：
            //res.err_msg将在用户支付成功后返回ok，但并不保证它绝对可靠。
            location.href="`+order.ReturnUrl+`";
      } 
   }); 
}
if (typeof WeixinJSBridge == "undefined"){
   if( document.addEventListener ){
       document.addEventListener("WeixinJSBridgeReady", onBridgeReady, false);
   }else if (document.attachEvent){
       document.attachEvent("WeixinJSBridgeReady", onBridgeReady); 
       document.attachEvent("onWeixinJSBridgeReady", onBridgeReady);
   }
}else{
   onBridgeReady();
}
</script>
`)

}

// Wechat 微信支付
func (con PayController) Wechat(c *gin.Context) {

	oid := c.Query("oid")

	if oid == "" {
		c.String(200, "订单号不能为空")
		return
	}

	var order models.Order
	ay.Db.First(&order, "oid = ?", oid)
	if order.Id == 0 {
		c.String(200, "订单错误")
		return
	}

	if order.Status == 1 {
		ay.Json{}.Msg(c, 400, "该笔订单已支付过", gin.H{})
		return
	}

	returnUrl := c.Query("return_url")
	if returnUrl != "" {
		returnUrl, _ = url.QueryUnescape(returnUrl)
		order.ReturnUrl = returnUrl
	} else {
		returnUrl = order.ReturnUrl
	}

	order.OutTradeNo = ay.MakeOrder(time.Now())
	if err := ay.Db.Save(&order).Error; err != nil {
		c.String(200, "请联系管理员")
		return
	}

	var pay models.Pay
	ay.Db.First(&pay, "id = ?", 6)

	models.AdvertLogModel{}.Add(1, order.Oid, c.GetHeader("Referer"), order.Amount, c.Query("request_id"), c.Query("ad_id"))

	if con.IsWechat(c) {

		redirectUri := url.QueryEscape(ay.Yaml.GetString("domain") + "/pay/open?oid=" + order.Oid)
		urlX := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + pay.Appid + "&redirect_uri=" + redirectUri + "&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect"
		c.Redirect(http.StatusTemporaryRedirect, urlX)
		return
	} else {

		code, msg := con.Web(order.OutTradeNo, 2, order.Amount, returnUrl, api.GetRequestIP(c), order.Des)

		if code == 1 {
			log.Println(msg)
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(200, `<script type="text/javascript">window.location.href="`+msg+`"</script>`)
			return
		} else {
			ay.Json{}.Msg(c, 400, msg, gin.H{})
			return
		}
	}

}

// IsWechat 微信ua判断
func (con PayController) IsWechat(c *gin.Context) bool {
	ua := c.GetHeader("User-Agent")
	if strings.Contains(ua, "MicroMessenger") == false && strings.Contains(ua, "Windows Phone") == false {
		return false
	} else {
		return true
	}
}

// Web web支付 1支付宝 2微信
func (con PayController) Web(outTradeNo string, payType int, amount float64, returnUrl string, ip string, msg string) (int, string) {

	ctx, _ := context.WithCancel(context.Background())

	if payType == 1 {

		var pay models.Pay
		ay.Db.First(&pay, "id = ?", 7)

		client, err := alipay.NewClient(pay.Appid, pay.VKey, true)
		if err != nil {
			return 0, err.Error()
		}
		client.SetLocation(alipay.LocationShanghai).
			SetCharset(alipay.UTF8).                                         // 设置字符编码，不设置默认 utf-8
			SetSignType(alipay.RSA2).                                        // 设置签名类型，不设置默认 RSA2
			SetReturnUrl(returnUrl).                                         // 设置返回URL
			SetNotifyUrl(ay.Yaml.GetString("domain") + "/api/notify/alipay") // 设置异步通知URL

		bm := make(gopay.BodyMap)

		bm.Set("subject", msg).
			Set("product_code", "QUICK_WAP_PAY").
			Set("out_trade_no", outTradeNo).
			Set("total_amount", amount).
			Set("quit_url", returnUrl) // 中途退出

		aliRsp, err := client.TradeWapPay(ctx, bm)

		if err != nil {
			return 0, err.Error()
		}

		return 1, aliRsp
	} else if payType == 2 {
		// 微信支付 jsapi需要跳转页面获取openid
		var pay models.Pay
		ay.Db.First(&pay, "id = ?", 6)

		client := wechat.NewClient(pay.Appid, pay.MchId, pay.VKey, true)
		// 打开Debug开关，输出请求日志，默认关闭
		//client.DebugSwitch = gopay.DebugOn
		client.SetCountry(wechat.China)

		bm := make(gopay.BodyMap)
		bm.Set("nonce_str", util.RandomString(32)).
			Set("body", msg).
			Set("out_trade_no", outTradeNo).
			Set("total_fee", amount*100).
			Set("spbill_create_ip", ip).
			Set("notify_url", ay.Yaml.GetString("domain")+"/api/notify/wechat").
			Set("trade_type", "MWEB").
			Set("device_info", "WEB").
			Set("sign_type", "MD5").
			SetBodyMap("scene_info", func(bm gopay.BodyMap) {
				bm.SetBodyMap("h5_info", func(bm gopay.BodyMap) {
					bm.Set("type", "Wap")
					bm.Set("wap_url", returnUrl)
					bm.Set("wap_name", "H5测试支付")
				})
			}) /*.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")*/

		wxRsp, err := client.UnifiedOrder(ctx, bm)

		log.Println(bm)
		log.Println("returnUrl:" + returnUrl)

		if err != nil {
			return 0, err.Error()
		}
		return 1, wxRsp.MwebUrl
	} else {
		return 0, "支付类型不正确"
	}

	return 0, ""
}
