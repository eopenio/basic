package config

// SmsConfig sms 配置 接口
type SmsConfig interface {
	GetEnabled() bool
	GetURL() string
}

// defaultSmsConfig sms 配置
type defaultSmsConfig struct {
	Enable bool   `json:"enabled"`
	URL    string `json:"url"`
}

// Enabled 激活
func (m defaultSmsConfig) GetEnabled() bool {
	return m.Enable
}

// SMS URL 连接
func (m defaultSmsConfig) GetURL() string {
	return m.URL
}
