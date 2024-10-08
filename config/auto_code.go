package config

type AutoCode struct {
	SModel          string `mapstructure:"server_model" json:"server_model" yaml:"server_model"`
	SRouter         string `mapstructure:"server_router" json:"server_router" yaml:"server_router"`
	Server          string `mapstructure:"server" json:"server" yaml:"server"`
	SApi            string `mapstructure:"server_api" json:"server_api" yaml:"server_api"`
	SPlug           string `mapstructure:"server_plug" json:"server_plug" yaml:"server_plug"`
	SInitialize     string `mapstructure:"server_initialize" json:"server_initialize" yaml:"server_initialize"`
	Root            string `mapstructure:"root" json:"root" yaml:"root"`
	WTable          string `mapstructure:"web_table" json:"web_table" yaml:"web_table"`
	Web             string `mapstructure:"web" json:"web" yaml:"web"`
	SService        string `mapstructure:"server_service" json:"server_service" yaml:"server_service"`
	SRequest        string `mapstructure:"server_request" json:"server_request"  yaml:"server_request"`
	WApi            string `mapstructure:"web_api" json:"web_api" yaml:"web_api"`
	WForm           string `mapstructure:"web_form" json:"web_form" yaml:"web_form"`
	TransferRestart bool   `mapstructure:"transfer_restart" json:"transfer_restart" yaml:"transfer_restart"`
}
