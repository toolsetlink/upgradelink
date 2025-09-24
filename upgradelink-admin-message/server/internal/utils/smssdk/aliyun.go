package smssdk

import (
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
)

func (t SmsConf) NewAliyunClient() (client *dysmsapi20170525.Client, err error) {
	config := &openapi.Config{
		AccessKeyId:     &t.SecretId,
		AccessKeySecret: &t.SecretKey,
	}

	if strings.HasPrefix(t.Region, "cn") {
		config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	} else {
		config.Endpoint = tea.String("dysmsapi.ap-southeast-1.aliyuncs.com")
	}

	client = &dysmsapi20170525.Client{}
	client, err = dysmsapi20170525.NewClient(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}
