package config

type EmailConfig struct {
	Smtp struct {
		Host     string `json:",default=localhost"`
		Port     int    `json:",default=1025"`
		Username string
		Password string
	}
	From struct {
		Address string `json:",default=gohub@example.com"`
		Name    string `json:",default=Gohub"`
	}
}
