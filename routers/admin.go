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
	"gin/controllers/admin"
	"github.com/gin-gonic/gin"
)

func AdminRouters(r *gin.RouterGroup) {

	adminGroup := r.Group("/admin/")
	// 登入
	adminGroup.POST("user/login", admin.Controller{}.Login)

}
