package config

// SmsConfig 配置文件中的短信配置
type SmsConfig struct {
	Aliyun struct {
		AccessKeyID     string
		AccessKeySecret string
		SignName        string
		TemplateCode    string
	}
}
