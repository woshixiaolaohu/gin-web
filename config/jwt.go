package config

type JWT struct {
	SigningKey  string `json:"signing_key" yaml:"signing_key" mapStructure:"signing_key"`    //JWT签名
	ExpiresTime string `json:"expires_time" yaml:"expires_time" mapStructure:"expires_time"` //过期时间
	BufferTime  string `json:"buffer_time" yaml:"buffer_time" mapStructure:"buffer_time"`    //缓冲时间
	Issuer      string `json:"issuer" yaml:"issuer" mapStructure:"issuer"`                   //签发者
}
