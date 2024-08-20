package config

type AwsS3 struct {
	Bucket           string `json:"bucket" yaml:"bucket" mapstructure:"bucket"`
	Region           string `json:"region" yaml:"region" mapstructure:"region"`
	EndPoint         string `json:"end_point" yaml:"end_point" mapstructure:"end_point"`
	SecretID         string `json:"secret_id" yaml:"secret_id" mapstructure:"secret_id"`
	SecretKey        string `json:"secret_key" yaml:"secret_key" mapstructure:"secret_key"`
	BaseURL          string `json:"base_url" yaml:"base_url" mapstructure:"base_url"`
	PathPrefix       string `json:"path_prefix" yaml:"path_prefix" mapstructure:"path_prefix"`
	S3ForcePathStyle bool   `json:"s3_force_path_style" yaml:"s3_force_path_style" mapstructure:"s3_force_path_style"`
	DisableSSL       bool   `json:"disable_ssl" yaml:"disable_ssl" mapstructure:"disable_ssl"`
}
