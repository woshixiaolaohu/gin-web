package config

type Local struct {
	Path      string `json:"path" yaml:"path" mapstructure:"path"`                   // 本地文件访问路径
	StorePath string `json:"store_path" yaml:"store_path" mapstructure:"store_path"` // 本地文件存储路径
}
