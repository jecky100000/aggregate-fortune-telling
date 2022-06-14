/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

import (
	"aggregate-fortune-telling/ay"
	"log"
	"net/url"
)

type AdvertLogModel struct {
}

type AdvertLog struct {
	BaseModel
	Type      int     `json:"type"`
	Oid       string  `json:"oid"`
	Cid       int64   `json:"cid"`
	Amount    float64 `json:"amount"`
	RequestId string  `json:"request_id"`
	AdId      string  `json:"ad_id"`
	Ip        string  `json:"ip"`
	Status    int     `json:"status"`
}

func (AdvertLog) TableName() string {
	return "sm_advert_log"
}

func (AdvertLogModel) Add(vType int, cid int64, oid string, urlx string, amount float64, requestId, adId, ip string) {
	f, _ := url.Parse(urlx)
	log.Println(f.Hostname())

	if vType == 1 {
		var advVivo AdvertVivo
		//ay.Db.Where("url = ?", f.Hostname()).First(&advVivo)
		ay.Db.Where("id = ?", cid).First(&advVivo)
		cid = advVivo.Id
	}

	var adv AdvertLog
	ay.Db.Where("oid = ?", oid).First(&adv)
	if adv.Id == 0 && requestId != "" && adId != "" && vType == 1 {
		ay.Db.Create(&AdvertLog{
			Type:      vType,
			Oid:       oid,
			Cid:       cid,
			Amount:    amount,
			RequestId: requestId,
			AdId:      adId,
			Status:    0,
			Ip:        ip,
		})
	}

	var adv1 AdvertLog
	ay.Db.Where("request_id = ? and ad_id = ? and type = ?", requestId, adId, vType).First(&adv1)
	if adv1.Id != 0 {
		adv1.Oid = oid
		adv1.Amount = amount
		ay.Db.Save(&adv1)
	}
}
