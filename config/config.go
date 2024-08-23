package config

type Server struct {
	Zap     Zap     `json:"zap" yaml:"zap" mapstructure:"zap"`
	JWT     JWT     `json:"jwt" yaml:"jwt" mapstructure:"jwt"`
	Redis   Redis   `json:"redis" yaml:"redis" mapstructure:"redis"`
	Mongo   Mongo   `json:"mongo" yaml:"mongo" mapstructure:"mongo"`
	System  System  `json:"system" yaml:"system" mapstructure:"system"`
	Captcha Captcha `json:"captcha" yaml:"captcha" mapstructure:"captcha"`
	// auto
	AutoCode AutoCode `json:"auto_code" yaml:"auto_code" mapstructure:"auto_code"`
	// gorm
	Mysql  Mysql          `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
	Mssql  Mssql          `json:"mssql" yaml:"mssql" mapstructure:"mssql"`
	Pgsql  Pgsql          `json:"pgsql" yaml:"pgsql" mapstructure:"pgsql"`
	Oracle Oracle         `json:"oracle" yaml:"oracle" mapstructure:"oracle"`
	Sqlite Sqlite         `json:"sqlite" yaml:"sqlite" mapstructure:"sqlite"`
	DBList []SpecializeDB `json:"db_list" yaml:"db_list" mapstructure:"db_list"`

	DiskList []DiskList `json:"disk_list" yaml:"disk_list" mapstructure:"disk_list"`
	// oss
	Local     Local     `json:"local" yaml:"local" mapstructure:"local"`
	AliyunOSS AliyunOSS `json:"aliyun_oss" yaml:"aliyun_oss" mapstructure:"aliyun_oss"`
}
