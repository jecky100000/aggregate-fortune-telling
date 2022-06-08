/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type AdvertVivoModel struct {
}

type AdvertVivo struct {
	BaseModel
	AdvertiserId string `json:"advertiser_id"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	SrcId        string `json:"src_id"`
	Url          string `json:"url"`
}

func (AdvertVivo) TableName() string {
	return "sm_advert_vivo"
}
