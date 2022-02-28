/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package routers

import (
	"gin/controllers/pay"
	"github.com/gin-gonic/gin"
)

func PayRouters(r *gin.RouterGroup) {

	payGroup := r.Group("/pay/")

	payGroup.GET("alipay", pay.Controller{}.AliPay)
	payGroup.GET("wechat", pay.Controller{}.Wechat)
	payGroup.GET("open", pay.Controller{}.GetOpenid)

}
