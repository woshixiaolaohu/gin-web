package config

type Email struct {
	To       string `json:"to" yaml:"to" mapstructure:"to"`                      // 收件人:多个以英文逗号分隔 例：a@qq.com b@qq.com 正式开发中请把此项目作为参数使用
	From     string `json:"from" yaml:"from" mapstructure:"from"`                // 发件人  你自己要发邮件的邮箱
	Host     string `json:"host" yaml:"host" mapstructure:"host"`                // 服务器地址 例如 smtp.qq.com  请前往QQ或者你要发邮件的邮箱查看其smtp协议
	Secret   string `json:"secret" yaml:"secret" mapstructure:"secret"`          // 密钥    用于登录的密钥 最好不要用邮箱密码 去邮箱smtp申请一个用于登录的密钥
	NickName string `json:"nick_name" yaml:"nick_name" mapstructure:"nick_name"` // 昵称    发件人昵称 通常为自己的邮箱
	Port     int    `json:"port" yaml:"port" mapstructure:"port"`                // 端口     请前往QQ或者你要发邮件的邮箱查看其smtp协议 大多为 465
	IsSSL    bool   `json:"is_ssl" yaml:"is_ssl" mapstructure:"is_ssl"`          // 是否SSL   是否开启SSL
}
