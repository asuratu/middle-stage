package config

type JwtConfig struct {
	AccessSecret string `json:",default=ae0536f9-6450-4606-8e13-5a19ed505da0"`
	AccessExpire int64  `json:",default=3600"` // 过期时间, 单位秒
}
