/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
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

	// 广告
	adv := models.AdvModel{}.GetType(1)

	for k, v := range adv {
		adv[k].Image = ay.Yaml.GetString("domain") + v.Image
	}

	// 热门咨询
	recommend := models.ConsultModel{}.GetType(1)
	hot := models.ConsultModel{}.GetType(2)

	var banner []models.Banner
	ay.Db.Order("sort asc").Find(&banner)

	for k, v := range banner {
		banner[k].Image = ay.Yaml.GetString("domain") + v.Image
	}

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"adv":   adv,
		"count": count,
		"consult": gin.H{
			"recommend": recommend,
			"hot":       hot,
		},
		"banner": banner,
	})
}

func (con HomeController) Config(c *gin.Context) {
	config := models.ConfigModel{}.GetId(1)
	// 广告
	adv := models.AdvModel{}.GetType(2)
	for k, v := range adv {
		adv[k].Image = ay.Yaml.GetString("domain") + v.Image
	}
	ay.Json{}.Msg(c, 200, "success", gin.H{
		"kf_link":     config.Kf,
		"master_link": config.MasterLink,
		"adv":         adv,
	})
}
