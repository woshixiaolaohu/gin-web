package config

type CORS struct {
	Mode      string          `json:"mode" yaml:"mode" mapstructure:"mode"`
	WhiteList []CORSWhiteList `json:"white_list" yaml:"white_list" mapstructure:"white_list"`
}
type CORSWhiteList struct {
	AllowOrigin      string `json:"allow_origin" yaml:"allow_origin" mapstructure:"allow_origin"`
	AllowMethods     string `json:"allow_methods" yaml:"allow_methods" mapstructure:"allow_methods"`
	AllowHeaders     string `json:"allow_headers" yaml:"allow_headers" mapstructure:"allow_headers"`
	ExposeHeaders    string `json:"expose_headers" yaml:"expose_headers" mapstructure:"expose_headers"`
	AllowCredentials bool   `json:"allow_credentials" yaml:"allow_credentials" mapstructure:"allow_credentials"`
}
