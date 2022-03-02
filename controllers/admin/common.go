package admin

import (
	"gin/ay"
	"gin/controllers/api"
	"gin/models"
	"log"
)

type CommonController struct {
}

func Auth() bool {
	token := ay.AuthCode(api.Token, "DECODE", "", 0)
	if token == "" {
		log.Println("没有token")
	}
	var user models.Admin
	ay.Db.First(&user, "account = ?", token)
	if user.Id == 0 {
		return false
	} else {
		return true
	}
}
