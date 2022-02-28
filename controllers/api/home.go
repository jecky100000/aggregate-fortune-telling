/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package api

import (
	"gin/ay"
	"gin/models"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
}

func (con HomeController) Home(c *gin.Context) {

	// 测算量
	var count int64
	var order models.Order
	ay.Db.Model(&order).Where("type = 1 OR type = 2").Count(&count)

	// 图片
	var image []models.Home

	ay.Db.Order("sort asc").Find(&image)

	for k, v := range image {
		image[k].Image = ay.Domain + v.Image
	}

	var recommend []models.Consult
	ay.Db.Order("sort asc").Find(&recommend, "type = 1")

	var hot []models.Consult
	ay.Db.Order("sort asc").Find(&hot, "type = 2")

	var banner []models.Banner
	ay.Db.Order("sort asc").Find(&banner)

	for k, v := range banner {
		banner[k].Image = ay.Domain + v.Image
	}

	ay.Json{}.Msg(c, "200", "success", gin.H{
		"image": image,
		"count": count,
		"consult": gin.H{
			"recommend": recommend,
			"hot":       hot,
		},
		"banner": banner,
	})
}

func (con HomeController) Config(c *gin.Context) {
	var config models.Config
	ay.Db.First(&config, 1)

	ay.Json{}.Msg(c, "200", "success", gin.H{
		"kf_link":     config.Kf,
		"master_link": config.MasterLink,
	})
}
