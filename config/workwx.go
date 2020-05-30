package config

// WorkwxConfig wxclient 配置 接口
type WorkwxConfig interface {
	GetEnabled() bool
	GetAgentId() int64
	GetCorpId() string
	GetCorpSecret() string
	GetQyAPIHostOverride() string
	GetTlsKeyLogFile() string
	GetAdminUser() string
}

// defaultWorkwxConfig wxclient 配置
type defaultWorkwxConfig struct {
	Enable            bool   `json:"enabled"`
	AgentId           int64  `json:"agentid"`
	CorpId            string `json:"corpid"`
	CorpSecret        string `json:"corpsecret"`
	AdminUser         string `json:"admin_user"`
	QyAPIHostOverride string `json:"qy_api_host_override,omitempty"`
	TlsKeyLogFile     string `json:"tls_key_log_file,omitempty"`
}

// Enabled 激活
func (m defaultWorkwxConfig) GetEnabled() bool {
	return m.Enable
}

// 企业 ID
func (m defaultWorkwxConfig) GetAgentId() int64 {
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

// 管理用户
func (m defaultWorkwxConfig) GetAdminUser() string {
	return m.AdminUser
}

// 使用自定义 HOST 覆盖默认企业微信 API 地址
func (m defaultWorkwxConfig) GetQyAPIHostOverride() string {
	return m.QyAPIHostOverride
}

// HTTPS 会话所用密钥
func (m defaultWorkwxConfig) GetTlsKeyLogFile() string {
	return m.TlsKeyLogFile
}
