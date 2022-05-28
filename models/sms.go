/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

import (
	"aggregate-fortune-telling/ay"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

type SmsModel struct {
}

type Sms struct {
	Id        int `gorm:"primaryKey" json:"id"`
	Phone     string
	Code      string
	Ip        string
	Ymd       string
	Status    int
	CreatedAt int64
}

func (Sms) TableName() string {
	return "sm_sms"
}

func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func (con SmsModel) Send(phone, code string) (_err error) {

	client, _err := CreateClient(tea.String(ay.Yaml.GetString("sms.ak")), tea.String(ay.Yaml.GetString("sms.sk")))
	if _err != nil {
		return _err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String(ay.Yaml.GetString("sms.sign")),
		TemplateCode:  tea.String(ay.Yaml.GetString("sms.template_code")),
		PhoneNumbers:  tea.String(phone),
		TemplateParam: tea.String(`{"time":"` + code + `"}`),
	}
	// 复制代码运行请自行打印 API 的返回值
	_, _err = client.SendSms(sendSmsRequest)

	if _err != nil {
		return _err
	}
	return _err
}
