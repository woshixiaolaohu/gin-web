package config

type Server struct {
	Zap Zap `json:"zap" yaml:"zap" mapStructure:"zap"`
	JWT JWT `json:"jwt" yaml:"jwt" mapStructure:"jwt"`
	// auto
	AutoCode AutoCode `json:"auto_code" yaml:"auto_code" mapStructure:"auto_code"`
	// gorm
	Mysql Mysql `json:"mysql" yaml:"mysql" mapStructure:"mysql"`
}
