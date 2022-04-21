/*
 *  * 六爻
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
	"gin/models/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

type HexapodController struct {
	CommonController
}

type List struct {
	Image string `json:"image"`
	Name  string `json:"name"`
}

// Index 六爻首页
func (con HexapodController) Index(c *gin.Context) {

	list := make([]List, 0)
	//list1 := make([]List, 0)

	for _, v1 := range common.HexapodName {

		for k, v := range common.HexapodWuName {
			if v1 == v {
				list = append(list, List{Image: ay.Yaml.GetString("domain") + "/static/image/hexapod/desc/" + strconv.Itoa(k) + ".png", Name: v})
			}
		}
	}

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list": list,
	})

}

type GetForm struct {
	Data string `form:"data" binding:"required"`
}

// Get 六爻详情
func (con HexapodController) Get(c *gin.Context) {
	var getForm GetForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	_, name := common.HexapodModel{}.Hexapod(getForm.Data)
	image := ""
	for k, v := range common.HexapodName {
		if v == name {
			image = ay.Yaml.GetString("domain") + "/static/image/hexapod/simple/" + strconv.Itoa(k) + ".png"
		}
	}

	res := models.HexapodModel{}.GetContonent(name)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"name":    name,
		"image":   image,
		"content": res.Content + res.Handount,
	})

}
