package config

type HuaWeiObs struct {
	Path      string `json:"path" yaml:"path" mapstructure:"path"`
	Bucket    string `json:"bucket" yaml:"bucket" mapstructure:"bucket"`
	EndPoint  string `json:"end_point" yaml:"end_point" mapstructure:"end_point"`
	AccessKey string `json:"access_key" yaml:"access_key" mapstructure:"access_key"`
	SecretKey string `json:"secret_key" yaml:"secret_key" mapstructure:"secret_key"`
}
