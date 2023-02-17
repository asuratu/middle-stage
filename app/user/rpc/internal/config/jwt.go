package config

type JwtConfig struct {
	AccessSecret  string `json:",default=ae0536f9-6450-4606-8e13-5a19ed505da0"`
	AccessExpire  int64  `json:",default=3600"`   // 过期时间, 单位秒 (1小时)
	RefreshExpire int64  `json:",default=604800"` // 过期时间, 单位秒 (7天)
	DebugExpire   int64  `json:",default=86400"`  // 调试过期时间, 单位秒 (1天)
}
