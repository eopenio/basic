package config

// WorkwxConfig workwx 配置 接口
type WorkwxConfig interface {
	GetEnabled() bool
	GetAgentId() int
	GetCorpId() string
	GetCorpSecret() string
	GetQyAPIHostOverride() string
	GetTlsKeyLogFile() string
}

// defaultWorkwxConfig workwx 配置
type defaultWorkwxConfig struct {
	Enable            bool   `json:"enabled"`
	AgentId           int    `json:"agentid"`
	CorpId            string `json:"corpid"`
	CorpSecret        string `json:"corpsecret"`
	QyAPIHostOverride string `json:"qy_api_host_override"`
	TlsKeyLogFile     string `json:"tls_key_log_file"`
}

// Enabled 激活
func (m defaultWorkwxConfig) GetEnabled() bool {
	return m.Enable
}

// 企业 ID
func (m defaultWorkwxConfig) GetAgentId() int {
	return m.AgentId
}

// 凭证密钥
func (m defaultWorkwxConfig) GetCorpId() string {
	return m.CorpId
}

// 企业应用 ID
func (m defaultWorkwxConfig) GetCorpSecret() string {
	return m.CorpSecret
}

// 使用自定义 HOST 覆盖默认企业微信 API 地址
func (m defaultWorkwxConfig) GetQyAPIHostOverride() string {
	return m.QyAPIHostOverride
}

// HTTPS 会话所用密钥
func (m defaultWorkwxConfig) GetTlsKeyLogFile() string {
	return m.TlsKeyLogFile
}
