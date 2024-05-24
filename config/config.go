package config

type Server struct {
	Zap    Zap    `json:"zap" yaml:"zap" mapStructure:"zap"`
	JWT    JWT    `json:"jwt" yaml:"jwt" mapStructure:"jwt"`
	System System `json:"system" yaml:"system" mapStructure:"system"`
	// auto
	AutoCode AutoCode `json:"auto_code" yaml:"auto_code" mapStructure:"auto_code"`
	// gorm
	Mysql  Mysql          `json:"mysql" yaml:"mysql" mapStructure:"mysql"`
	Mssql  Mssql          `json:"mssql" yaml:"mssql" mapStructure:"mssql"`
	Pgsql  Pgsql          `json:"pgsql" yaml:"pgsql" mapStructure:"pgsql"`
	Oracle Oracle         `json:"oracle" yaml:"oracle" mapStructure:"oracle"`
	Sqlite Sqlite         `json:"sqlite" yaml:"sqlite" mapStructure:"sqlite"`
	DBList []SpecializeDB `json:"db_list" yaml:"db_list" mapStructure:"db_list"`
}
