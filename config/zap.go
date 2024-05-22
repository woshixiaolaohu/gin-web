package config

type Zap struct {
	Level         string `json:"level" yaml:"level" mapStructure:"level"`                            //级别
	Prefix        string `json:"prefix" yaml:"prefix" mapStructure:"prefix"`                         //日志前缀
	Format        string `json:"format" yaml:"format" mapStructure:"format"`                         //输出
	Director      string `json:"director" yaml:"director" mapStructure:"director"`                   //日志文件夹
	EncodeLevel   string `json:"encode_level" yaml:"encode_level" mapStructure:"encode_level"`       //编码级
	StacktraceKey string `json:"stacktrace_key" yaml:"stacktrace_key" mapStructure:"stacktrace_key"` //栈名
	ShowLine      bool   `json:"show_line" yaml:"show_line" mapStructure:"show_line"`                //显示行
	LogInConsole  bool   `json:"log_in_console" yaml:"log_in_console" mapStructure:"log_in_console"` //输出控制台
}
