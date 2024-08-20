package config

type TencentCOS struct {
	Bucket     string `json:"bucket" yaml:"bucket" mapstructure:"bucket"`
	Region     string `json:"region" yaml:"region" mapstructure:"region"`
	SecretID   string `json:"secret_id" yaml:"secret_id" mapstructure:"secret_id"`
	SecretKey  string `json:"secret_key" yaml:"secret_key" mapstructure:"secret_key"`
	BaseURL    string `json:"base_url" yaml:"base_url" mapstructure:"base_url"`
	PathPrefix string `json:"path_prefix" yaml:"path_prefix" mapstructure:"path_prefix"`
}
