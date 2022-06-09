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
	Status    int     `json:"status"`
}

func (AdvertLog) TableName() string {
	return "sm_advert_log"
}

func (AdvertLogModel) Add(vType int, oid string, urlx string, amount float64, requestId, adId string) {
	f, _ := url.Parse(urlx)
	log.Println(f.Hostname())
	var advVivo AdvertVivo
	//ay.Db.Where("url = ?", f.Hostname()).First(&advVivo)
	ay.Db.Where("id = ?", 1).First(&advVivo)

	var adv AdvertLog
	ay.Db.Where("oid = ?", oid).First(&adv)
	if adv.Id == 0 && requestId != "" && adId != "" {
		ay.Db.Create(&AdvertLog{
			Type:      vType,
			Oid:       oid,
			Cid:       advVivo.Id,
			Amount:    amount,
			RequestId: requestId,
			AdId:      adId,
			Status:    0,
		})
	}

}
