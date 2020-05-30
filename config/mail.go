package config

// MailConfig mail 配置 接口
type MailConfig interface {
	GetEnabled() bool
	GetHost() string
	GetPort() int
	GetSender() string
	GetUsername() string
	GetPassword() string
}

// defaultMailConfig mail 配置
type defaultMailConfig struct {
	Enable   bool   `json:"enabled"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Sender   string `json:"sender"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Enabled 激活
func (m defaultMailConfig) GetEnabled() bool {
	return m.Enable
}

// 邮件服务地址
func (m defaultMailConfig) GetHost() string {
	return m.Host
}

// 邮件服务端口
func (m defaultMailConfig) GetPort() int {
	return m.Port
}

// 发送方地址
func (m defaultMailConfig) GetSender() string {
	return m.Sender
}

// 登陆账号
func (m defaultMailConfig) GetUsername() string {
	return m.Username
}

// 登陆密码
func (m defaultMailConfig) GetPassword() string {
	return m.Password
}
