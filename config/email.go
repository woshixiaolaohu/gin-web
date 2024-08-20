package config

type Email struct {
	To       string `json:"to" yaml:"to" mapstructure:"to"`                      // 收件人：多个英文逗号隔开
	From     string `json:"from" yaml:"from" mapstructure:"form"`                // 发件人
	Host     string `json:"host" yaml:"host" mapstructure:"host"`                // 服务器地址 smtp查看
	Secret   string `json:"secret" yaml:"secret" mapstructure:"secret"`          // 密钥 邮箱smtp申请用于登录的密钥
	NickName string `json:"nick_name" yaml:"nick_name" mapstructure:"nick_name"` // 昵称
	Port     int    `json:"port" yaml:"port" mapstructure:"port"`                // 端口 smtp协议端口 大部分为465
	IsSSL    bool   `json:"is_ssl" yaml:"is_ssl" mapstructure:"is_ssl"`          //是否开启ssl
}
