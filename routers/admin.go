/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package routers

import (
	"aggregate-fortune-telling/controllers/admin"
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
	adminGroup.POST("account/upload", admin.AccountController{}.Upload)

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

	// 热门推荐
	adminGroup.GET("consult/list", admin.ConsultController{}.List)
	adminGroup.POST("consult/detail", admin.ConsultController{}.Detail)
	adminGroup.POST("consult/option", admin.ConsultController{}.Option)
	adminGroup.POST("consult/delete", admin.ConsultController{}.Delete)

	// 百科
	adminGroup.GET("baike/list", admin.BaikeController{}.List)
	adminGroup.POST("baike/detail", admin.BaikeController{}.Detail)
	adminGroup.POST("baike/option", admin.BaikeController{}.Option)
	adminGroup.POST("baike/delete", admin.BaikeController{}.Delete)
	adminGroup.POST("baike/upload", admin.BaikeController{}.Upload)

	// 古籍
	adminGroup.GET("ancient/list", admin.AncientController{}.List)
	adminGroup.POST("ancient/detail", admin.AncientController{}.Detail)
	adminGroup.POST("ancient/option", admin.AncientController{}.Option)
	adminGroup.POST("ancient/delete", admin.AncientController{}.Delete)
	adminGroup.POST("ancient/upload", admin.AncientController{}.Upload)

	adminGroup.GET("ancient_type/all", admin.AncientTypeController{}.All)

	// 古籍目录
	adminGroup.GET("ancient_catalogue/list", admin.AncientCatalogueController{}.List)
	adminGroup.POST("ancient_catalogue/detail", admin.AncientCatalogueController{}.Detail)
	adminGroup.POST("ancient_catalogue/option", admin.AncientCatalogueController{}.Option)
	adminGroup.POST("ancient_catalogue/delete", admin.AncientCatalogueController{}.Delete)
	adminGroup.POST("ancient_catalogue/upload", admin.AncientCatalogueController{}.Upload)

	// 古籍分类
	adminGroup.GET("ancient_type/list", admin.AncientTypeController{}.List)
	adminGroup.POST("ancient_type/detail", admin.AncientTypeController{}.Detail)
	adminGroup.POST("ancient_type/option", admin.AncientTypeController{}.Option)
	adminGroup.POST("ancient_type/delete", admin.AncientTypeController{}.Delete)

	// 优惠卷
	adminGroup.GET("coupon/list", admin.CouponController{}.List)
	adminGroup.POST("coupon/detail", admin.CouponController{}.Detail)
	adminGroup.POST("coupon/option", admin.CouponController{}.Option)
	adminGroup.POST("coupon/delete", admin.CouponController{}.Delete)

	// 周公解梦
	adminGroup.GET("dream/list", admin.DreamController{}.List)
	adminGroup.POST("dream/detail", admin.DreamController{}.Detail)
	adminGroup.POST("dream/option", admin.DreamController{}.Option)
	adminGroup.POST("dream/delete", admin.DreamController{}.Delete)

	// 问题示例
	adminGroup.GET("asklog/list", admin.AskLogController{}.List)
	adminGroup.POST("asklog/detail", admin.AskLogController{}.Detail)
	adminGroup.POST("asklog/option", admin.AskLogController{}.Option)
	adminGroup.POST("asklog/delete", admin.AskLogController{}.Delete)

	// 八字测算金额
	adminGroup.GET("haulAmount/list", admin.HaulAmountController{}.List)
	adminGroup.POST("haulAmount/detail", admin.HaulAmountController{}.Detail)
	adminGroup.POST("haulAmount/option", admin.HaulAmountController{}.Option)
	adminGroup.POST("haulAmount/delete", admin.HaulAmountController{}.Delete)

	// 八字测算案例
	adminGroup.GET("haulCases/list", admin.HaulCasesController{}.List)
	adminGroup.POST("haulCases/detail", admin.HaulCasesController{}.Detail)
	adminGroup.POST("haulCases/option", admin.HaulCasesController{}.Option)
	adminGroup.POST("haulCases/delete", admin.HaulCasesController{}.Delete)
	adminGroup.POST("haulCases/upload", admin.HaulCasesController{}.Upload)

	// 大师
	adminGroup.POST("master/detail", admin.MasterController{}.Detail)
	adminGroup.POST("master/option", admin.MasterController{}.Option)
	adminGroup.POST("master/upload", admin.MasterController{}.Upload)

	adminGroup.GET("master/type", admin.MasterController{}.AllType)
	adminGroup.GET("master/all", admin.MasterController{}.All)

	// 系统
	adminGroup.POST("config/detail", admin.ConfigController{}.Detail)
	adminGroup.POST("config/option", admin.ConfigController{}.Option)

	// 提问回复
	adminGroup.POST("ask_reply/detail", admin.AskReplyController{}.Detail)
	adminGroup.POST("ask_reply/option", admin.AskReplyController{}.Option)
	adminGroup.GET("ask_reply/list", admin.AskReplyController{}.List)
	adminGroup.POST("ask_reply/delete", admin.AskReplyController{}.Delete)

}
