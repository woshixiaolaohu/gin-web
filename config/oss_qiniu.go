package config

type QiNiu struct {
	Zone          string `json:"zone" yaml:"zone" mapstructure:"zone"`                                  // 存储区域
	Bucket        string `json:"bucket" yaml:"bucket" mapstructure:"bucket"`                            // 空间名称
	ImgPath       string `json:"img_path" yaml:"img_path" mapstructure:"img_path"`                      // CDN加速域名
	AccessKey     string `json:"access_key" yaml:"access_key" mapstructure:"access_key"`                // 密钥AK
	SecretKey     string `json:"secret_key" yaml:"secret_key" mapstructure:"secret_key"`                // 密钥SK
	UseHttps      bool   `json:"use_https" yaml:"use_https" mapstructure:"use_https"`                   // 是否使用https
	UseCdnDomains bool   `json:"use_cdn_domains" yaml:"use_cdn_domains" mapstructure:"use_cdn_domains"` // 上传是否使用CDN加速
}
