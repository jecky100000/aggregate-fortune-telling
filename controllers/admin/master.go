/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package admin

import (
	"gin/ay"
	"gin/models"
	"github.com/gin-gonic/gin"
)

type MasterController struct {
}

// Detail 用户详情
func (con MasterController) Detail(c *gin.Context) {
	var data orderDetailForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var user models.User
	ay.Db.First(&user, data.Id)

	var master models.Master

	ay.Db.First(&master, user.MasterId)

	if master.Id == 0 {
		masterData := &models.Master{
			Type:        "",
			Sign:        "",
			Label:       "",
			Introduce:   "",
			Years:       0,
			Online:      0,
			Rate:        0,
			AskNum:      0,
			BackImage:   "",
			IsRecommend: 0,
		}
		if err := ay.Db.Create(masterData).Error; err == nil {
			user.MasterId = masterData.Id
			ay.Db.Save(&user)
		} else {
			ay.Json{}.Msg(c, 400, "创建大师资料失败", gin.H{})
			return
		}
	}

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"info": master,
	})
}

type masterOptionForm struct {
	Id          int     `form:"id"`
	Type        string  `form:"type"`
	Sign        string  `form:"sign"`
	Label       string  `form:"label"`
	Years       int     `form:"years"`
	AskNum      int     `form:"ask_num"`
	Rate        float64 `form:"rate"`
	IsRecommend int     `form:"is_recommend"`
	BackImage   string  `form:"back_image"`
	Introduce   string  `form:"introduce"`
}

// Option 添加 编辑
func (con MasterController) Option(c *gin.Context) {
	var data masterOptionForm
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var res models.Master
	ay.Db.First(&res, data.Id)

	if data.Id != 0 {

		res.Type = data.Type
		res.Sign = data.Sign
		res.Label = data.Label
		res.Years = data.Years
		res.AskNum = data.AskNum
		res.Rate = data.Rate
		res.IsRecommend = data.IsRecommend
		res.BackImage = data.BackImage
		res.Introduce = data.Introduce

		if err := ay.Db.Save(&res).Error; err != nil {
			ay.Json{}.Msg(c, 400, "修改失败", gin.H{})
		} else {
			ay.Json{}.Msg(c, 200, "修改成功", gin.H{})
		}
	} else {
		ay.Db.Create(&models.Master{
			Type:        data.Type,
			Sign:        data.Sign,
			Label:       data.Label,
			Introduce:   data.Introduce,
			Years:       data.Years,
			Rate:        data.Rate,
			AskNum:      data.AskNum,
			BackImage:   data.BackImage,
			IsRecommend: data.IsRecommend,
		})

		ay.Json{}.Msg(c, 200, "创建成功", gin.H{})

	}

}

func (con MasterController) Upload(c *gin.Context) {

	code, msg := Upload(c, "master")

	if code != 200 {
		ay.Json{}.Msg(c, 400, msg, gin.H{})
	} else {
		ay.Json{}.Msg(c, 200, msg, gin.H{})
	}
}

func (con MasterController) AllType(c *gin.Context) {
	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}
	type list struct {
		Label string `gorm:"column:name" json:"label"`
		Value string `gorm:"column:id" json:"value"`
	}
	var l []list
	ay.Db.Table("sm_master_type").Find(&l)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"list": l,
	})
}
