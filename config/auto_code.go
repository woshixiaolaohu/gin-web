package config

type AutoCode struct {
	Server           string `json:"server" yaml:"server" mapStructure:"server"`
	ServerModel      string `json:"server_model" yaml:"server_model" mapStructure:"server_model"`
	ServerRouter     string `json:"server_router" yaml:"server_router" mapStructure:"server_router"`
	ServerApi        string `json:"server_api" yaml:"server_api" mapStructure:"server_api"`
	ServerPlug       string `json:"server_plug" yaml:"server_plug" mapStructure:"server_plug"`
	ServerInitialize string `json:"server_initialize" yaml:"server_initialize" mapStructure:"server_initialize"`
	Root             string `json:"root" yaml:"root" mapStructure:"root"`
	Web              string `json:"web" yaml:"web" mapStructure:"web"`
	WebTable         string `json:"web_table" yaml:"web_table" mapStructure:"web_table"`
	ServerService    string `json:"server_service" yaml:"server_service" mapStructure:"server_service"`
	ServerRequest    string `json:"server_request" yaml:"server_request" mapStructure:"server_request"`
	WebApi           string `json:"web_api" yaml:"web_api" mapStructure:"web_api"`
	WebForm          string `json:"web_form" yaml:"web_form" mapStructure:"web_form"`
}
