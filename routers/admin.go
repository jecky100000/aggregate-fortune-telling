/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package routers

import (
	"gin/controllers/admin"
	"github.com/gin-gonic/gin"
)

func AdminRouters(r *gin.RouterGroup) {

	adminGroup := r.Group("/admin/")
	// 登入
	adminGroup.POST("user/login", admin.Controller{}.Login)

	// 用户
	adminGroup.GET("account/list", admin.AccountController{}.List)
	adminGroup.POST("account/detail", admin.AccountController{}.Detail)
	adminGroup.POST("account/option", admin.AccountController{}.Option)
	adminGroup.POST("account/delete", admin.AccountController{}.Delete)

	// 订单
	adminGroup.GET("order/list", admin.OrderController{}.List)
	adminGroup.POST("order/detail", admin.OrderController{}.Detail)
	adminGroup.POST("order/option", admin.OrderController{}.Option)
	adminGroup.POST("order/delete", admin.OrderController{}.Delete)

}
