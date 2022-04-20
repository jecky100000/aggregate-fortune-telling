/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package pay

import (
	"context"
	"encoding/json"
	"gin/ay"
	"gin/controllers/api"
	"gin/models"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
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

type Controller struct {
}

func (con Controller) AliPay(c *gin.Context) {
	var order models.Order
	ay.Db.First(&order, "oid = ?", c.Query("oid"))
	if order.Id == 0 {
		c.String(200, "订单错误")
		return
	}
	c.Redirect(http.StatusMovedPermanently, order.Json)
}

type OpenId struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
}

func (con Controller) GetOpenid(c *gin.Context) {

	oid := c.Query("oid")
	code := c.Query("code")

	var order models.Order
	ay.Db.First(&order, "oid = ?", oid)

	out_trade_no := ay.MakeOrder(time.Now())
	order.OutTradeNo = out_trade_no
	ay.Db.Save(&order)

	config := models.ConfigModel{}.GetId(1)

	var pay models.Pay
	ay.Db.First(&pay, "id = ?", 6)

	oauth2Url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + pay.Appid + "&secret=" + pay.Secret + "&code=" + code + "&grant_type=authorization_code"

	resp, _ := http.Get(oauth2Url)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var j OpenId
	err := json.Unmarshal(body, &j)
	if err != nil {
		api.Json.Msg(400, "获取openid失败", gin.H{})
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
	v := strconv.FormatFloat(order.Amount*config.Rate, 'g', -1, 64)
	bm.Set("nonce_str", util.GetRandomString(32)).
		Set("body", "充值"+v+"元").
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
      'getBrandWCPayRequest', {
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
       document.addEventListener('WeixinJSBridgeReady', onBridgeReady, false);
   }else if (document.attachEvent){
       document.attachEvent('WeixinJSBridgeReady', onBridgeReady); 
       document.attachEvent('onWeixinJSBridgeReady', onBridgeReady);
   }
}else{
   onBridgeReady();
}
</script>
`)

}

func (con Controller) Wechat(c *gin.Context) {

	var order models.Order
	ay.Db.First(&order, "oid = ?", c.Query("oid"))
	if order.Id == 0 {
		c.String(200, "订单错误")
		return
	}

	if order.Status == 1 {
		api.Json.Msg(400, "该笔订单已支付过", gin.H{})
		return
	}

	var pay models.Pay
	ay.Db.First(&pay, "id = ?", 6)

	if con.IsWechat(c) {
		redirect_uri := url.QueryEscape(ay.Yaml.GetString("domain") + "/pay/open?oid=" + c.Query("oid"))
		urlx := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + pay.Appid + "&redirect_uri=" + redirect_uri + "&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect"
		c.Redirect(http.StatusTemporaryRedirect, urlx)
		//c.Header("Content-Type", "text/html; charset=utf-8")
		//c.String(200, `<script type="text/javascript">window.location.href="`+urlx+`"</script>`)
		//c.String(200, urlx)
	} else {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, `<script type="text/javascript">window.location.href="`+order.Json+`"</script>`)
		//c.Redirect(http.StatusTemporaryRedirect, order.Json)
	}

}

func (con Controller) IsWechat(c *gin.Context) bool {
	ua := c.GetHeader("User-Agent")
	if strings.Contains(ua, "MicroMessenger") == false && strings.Contains(ua, "Windows Phone") == false {
		return false
	} else {
		return true
	}
}
