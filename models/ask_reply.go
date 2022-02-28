/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package models

type AskReplyModel struct {
}

type AskReply struct {
	BaseModel
	AskId    string `json:"ask_id"`
	MasterId int64  `json:"master_id"`
	Content  string `json:"content"`
	Adopt    int    `json:"adopt"`
}

func (AskReply) TableName() string {
	return "sm_ask_reply"
}
