package config

type JWT struct {
	SigningKey  string `json:"signing_key" yaml:"signing_key" mapstructure:"signing_key"`    //JWT签名
	ExpiresTime string `json:"expires_time" yaml:"expires_time" mapstructure:"expires_time"` //过期时间
	BufferTime  string `json:"buffer_time" yaml:"buffer_time" mapstructure:"buffer_time"`    //缓冲时间
	Issuer      string `json:"issuer" yaml:"issuer" mapstructure:"issuer"`                   //签发者
}
