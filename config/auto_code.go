package config

type AutoCode struct {
	Server           string `json:"server" yaml:"server" mapstructure:"server"`
	ServerModel      string `json:"server_model" yaml:"server_model" mapstructure:"server_model"`
	ServerRouter     string `json:"server_router" yaml:"server_router" mapstructure:"server_router"`
	ServerApi        string `json:"server_api" yaml:"server_api" mapstructure:"server_api"`
	ServerPlug       string `json:"server_plug" yaml:"server_plug" mapstructure:"server_plug"`
	ServerInitialize string `json:"server_initialize" yaml:"server_initialize" mapstructure:"server_initialize"`
	Root             string `json:"root" yaml:"root" mapstructure:"root"`
	Web              string `json:"web" yaml:"web" mapstructure:"web"`
	WebTable         string `json:"web_table" yaml:"web_table" mapstructure:"web_table"`
	ServerService    string `json:"server_service" yaml:"server_service" mapstructure:"server_service"`
	ServerRequest    string `json:"server_request" yaml:"server_request" mapstructure:"server_request"`
	WebApi           string `json:"web_api" yaml:"web_api" mapstructure:"web_api"`
	WebForm          string `json:"web_form" yaml:"web_form" mapstructure:"web_form"`
}
