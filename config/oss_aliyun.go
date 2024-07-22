package config

type AliyunOSS struct {
	EndPoint        string `json:"end_point" yaml:"end_point" mapstructure:"end_point"`
	AccessKeyID     string `json:"access_key_id" yaml:"access_key_id" mapstructure:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret" yaml:"access_key_secret" mapstructure:"access_key_secret"`
	BucketName      string `json:"bucket_name" yaml:"bucket_name" mapstructure:"bucket_name"`
	BucketUrl       string `json:"bucket_url" yaml:"bucket_url" mapstructure:"bucket_url"`
	BasePath        string `json:"base_path" yaml:"base_path" mapstructure:"base_path"`
}
