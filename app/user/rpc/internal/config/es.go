package config

type EsConfig struct {
	Address  string `json:",default=http://localhost:9200"`
	Username string // 用户名
	Password string // 密码
}
