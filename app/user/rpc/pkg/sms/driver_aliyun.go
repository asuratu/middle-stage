package sms

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"

	aliyunsmsclient "github.com/KenmyZhang/aliyun-communicate"
)

// Aliyun 实现 sms.Driver interface
type Aliyun struct{}

// Send 实现 sms.Driver interface 的 Send 方法
func (s *Aliyun) Send(phone string, message Message, config map[string]string) bool {

	smsClient := aliyunsmsclient.New("http://dysmsapi.aliyuncs.com/")

	templateParam, err := json.Marshal(message.Data)
	if err != nil {
		logx.Errorf("短信[阿里云] ：解析绑定错误, err:%v", err.Error())
		return false
	}

	logx.Debugf("短信[阿里云] ：配置信息, %+v", config)

	result, err := smsClient.Execute(
		config["access_key_id"],
		config["access_key_secret"],
		phone,
		config["sign_name"],
		message.Template,
		string(templateParam),
	)

	logx.Debugf("短信[阿里云] ：请求内容, err:%v", smsClient.Request)
	logx.Debugf("短信[阿里云] ：接口响应, err:%v", result)

	if err != nil {
		logx.Errorf("短信[阿里云] ：发信失败, err:%v", err.Error())
		return false
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		logx.Errorf("短信[阿里云] ：解析响应 JSON 错误, err:%v", err.Error())
		return false
	}

	if result.IsSuccessful() {
		logx.Debugf("短信[阿里云] ：发信成功")
		return true
	} else {
		logx.Errorf("短信[阿里云] ：服务商返回错误, err:%s", string(resultJSON))
		return false
	}
}
