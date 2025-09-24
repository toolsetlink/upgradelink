// Package smssdk defines all the sdk integrated for simple admin mms service
package smssdk

import (
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	unisms "github.com/apistd/uni-go-sdk/sms"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type SmsConf struct {
	SecretId  string
	SecretKey string
	Provider  string
	Region    string
}

type SmsGroup struct {
	TencentSmsClient *sms.Client
	AliyunSmsClient  *dysmsapi20170525.Client
	UniSmsClient     *unisms.UniSMSClient
}
