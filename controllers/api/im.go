/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package api

import (
	"encoding/json"
	"gin/ay"
	"gin/models"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

type ImController struct {
}

func (con ImController) Notify(c *gin.Context) {
	cmd := c.Query("CallbackCommand")
	log.Println(cmd)
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	log.Println(string(bodyBytes))

	if cmd == "State.StateChange" {
		// 在线
		//con.Online(bodyBytes)
	}

}

func (con ImController) Online(str []byte) {
	type status struct {
		CallbackCommand string `json:"CallbackCommand"`
		EventTime       int64  `json:"EventTime"`
		Info            struct {
			ToAccount string `json:"To_Account"`
			Action    string `json:"Action"`
			Reason    string `json:"Reason"`
		} `json:"Info"`
	}

	var data status
	json.Unmarshal(str, &data)

	online := 0
	if data.Info.Action == "Login" {
		online = 1
	}

	var user models.User
	ay.Db.Table("sm_user").Where("phone = ?", data.Info.ToAccount).First(&user)
	ay.Db.Model(models.Master{}).Where("id = ?", user.MasterId).UpdateColumn("online", online)

}

func (con ImController) Msg(str []byte) {

	type msg struct {
		MsgBody []struct {
			MsgType    string `json:"MsgType"`
			MsgContent struct {
				Text string `json:"Text"`
			} `json:"MsgContent"`
		} `json:"MsgBody"`
		CallbackCommand string `json:"CallbackCommand"`
		FromAccount     string `json:"From_Account"`
		ToAccount       string `json:"To_Account"`
		MsgRandom       int    `json:"MsgRandom"`
		MsgSeq          int    `json:"MsgSeq"`
		MsgTime         int    `json:"MsgTime"`
		MsgKey          string `json:"MsgKey"`
		OnlineOnlyFlag  int    `json:"OnlineOnlyFlag"`
		SendMsgResult   int    `json:"SendMsgResult"`
		ErrorInfo       string `json:"ErrorInfo"`
		UnreadMsgNum    int    `json:"UnreadMsgNum"`
	}

	var data msg
	json.Unmarshal(str, &data)

}
