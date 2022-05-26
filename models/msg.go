/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type MsgModel struct {
}

type Msg struct {
	BaseModel
	Type           int    `json:"type"`
	Account        string `json:"account"`
	ToAccount      string `json:"toAccount"`
	Content        string `json:"content"`
	MsgSeq         int    `json:"msgSeq"`
	MsgKey         string `json:"msgKey"`
	SendAt         MyTime `json:"sendAt"`
	Ip             string `json:"ip"`
	SendMsgResult  int    `json:"SendMsgResult"`
	OnlineOnlyFlag int    `json:"onlineOnlyFlag"`
}

func (Msg) TableName() string {
	return "sm_msg"
}
