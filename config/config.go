package config

type Server struct {
	// auto
	AutoCode AutoCode `json:"auto_code" yaml:"auto_code" mapStructure:"auto_code"`
	// gorm
	Mysql Mysql `json:"mysql" yaml:"mysql" mapStructure:"mysql"`
}
