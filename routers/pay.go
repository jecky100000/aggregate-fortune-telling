/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package routers

import (
	"aggregate-fortune-telling/controllers/pay"
	"github.com/gin-gonic/gin"
)

func PayRouters(r *gin.RouterGroup) {

	payGroup := r.Group("/pay/")

	payGroup.GET("alipay", pay.PayController{}.AliPay)
	payGroup.GET("wechat", pay.PayController{}.Wechat)
	payGroup.GET("open", pay.PayController{}.GetOpenid)

}
