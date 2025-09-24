package smssdk

import (
	unisms "github.com/apistd/uni-go-sdk/sms"
)

func (t SmsConf) NewUniClient() (client *unisms.UniSMSClient, err error) {
	client = unisms.NewClient(t.SecretId, t.SecretKey)
	return client, nil
}
