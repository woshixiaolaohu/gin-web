package config

type Server struct {
	Zap    Zap    `json:"zap" yaml:"zap" mapstructure:"zap"`
	JWT    JWT    `json:"jwt" yaml:"jwt" mapstructure:"jwt"`
	System System `json:"system" yaml:"system" mapstructure:"system"`
	// auto
	AutoCode AutoCode `json:"auto_code" yaml:"auto_code" mapstructure:"auto_code"`
	// gorm
	Mysql  Mysql          `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
	Mssql  Mssql          `json:"mssql" yaml:"mssql" mapstructure:"mssql"`
	Pgsql  Pgsql          `json:"pgsql" yaml:"pgsql" mapstructure:"pgsql"`
	Oracle Oracle         `json:"oracle" yaml:"oracle" mapstructure:"oracle"`
	Sqlite Sqlite         `json:"sqlite" yaml:"sqlite" mapstructure:"sqlite"`
	DBList []SpecializeDB `json:"db_list" yaml:"db_list" mapstructure:"db_list"`
}
