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

	// 文章
	adminGroup.GET("notice/list", admin.NoticeController{}.List)
	adminGroup.POST("notice/detail", admin.NoticeController{}.Detail)
	adminGroup.POST("notice/option", admin.NoticeController{}.Option)
	adminGroup.POST("notice/delete", admin.NoticeController{}.Delete)
	adminGroup.POST("notice/upload", admin.NoticeController{}.Upload)

	// Banner
	adminGroup.GET("banner/list", admin.BannerController{}.List)
	adminGroup.POST("banner/detail", admin.BannerController{}.Detail)
	adminGroup.POST("banner/option", admin.BannerController{}.Option)
	adminGroup.POST("banner/delete", admin.BannerController{}.Delete)
	adminGroup.POST("banner/upload", admin.BannerController{}.Upload)

	// adv
	adminGroup.GET("adv/list", admin.AdvController{}.List)
	adminGroup.POST("adv/detail", admin.AdvController{}.Detail)
	adminGroup.POST("adv/option", admin.AdvController{}.Option)
	adminGroup.POST("adv/delete", admin.AdvController{}.Delete)
	adminGroup.POST("adv/upload", admin.AdvController{}.Upload)

	// 文章分类
	adminGroup.GET("notice_type/all", admin.NoticeTypeController{}.All)
	adminGroup.GET("notice_type/list", admin.NoticeTypeController{}.List)
	adminGroup.POST("notice_type/detail", admin.NoticeTypeController{}.Detail)
	adminGroup.POST("notice_type/option", admin.NoticeTypeController{}.Option)
	adminGroup.POST("notice_type/delete", admin.NoticeTypeController{}.Delete)

}
