package config

import (
	"go.uber.org/zap/zapcore"
	"time"
)

type Zap struct {
	Level         string `json:"level" yaml:"level" mapstructure:"level"`                            //级别
	Prefix        string `json:"prefix" yaml:"prefix" mapstructure:"prefix"`                         //日志前缀
	Format        string `json:"format" yaml:"format" mapstructure:"format"`                         //输出
	Director      string `json:"director" yaml:"director" mapstructure:"director"`                   //日志文件夹
	EncodeLevel   string `json:"encode_level" yaml:"encode_level" mapstructure:"encode_level"`       //编码级
	StacktraceKey string `json:"stacktrace_key" yaml:"stacktrace_key" mapstructure:"stacktrace_key"` //栈名
	ShowLine      bool   `json:"show_line" yaml:"show_line" mapstructure:"show_line"`                //显示行
	LogInConsole  bool   `json:"log_in_console" yaml:"log_in_console" mapstructure:"log_in_console"` //输出控制台
}

// Levels 根据字符串转化为zapcore.Levels
func (z *Zap) Levels() []zapcore.Level {
	levels := make([]zapcore.Level, 0, 7)
	level, err := zapcore.ParseLevel(z.Level)
	if err != nil {
		level = zapcore.DebugLevel
	}
	for ; level <= zapcore.FatalLevel; level++ {
		levels = append(levels, level)
	}
	return levels
}

func (z *Zap) Encoder() zapcore.Encoder {
	config := zapcore.EncoderConfig{
		TimeKey:       "time",
		NameKey:       "name",
		LevelKey:      "level",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: z.StacktraceKey,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(z.Prefix + t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeLevel:    z.LevelEncoder(),
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	if z.Format == "json" {
		return zapcore.NewJSONEncoder(config)
	}
	return zapcore.NewConsoleEncoder(config)
}

// LevelEncoder  根据 EncodeLevel 返回 zapcore.LevelEncoder
func (z *Zap) LevelEncoder() zapcore.LevelEncoder {
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder": //小写编码器（默认）
		return zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder": //大写编码器
		return zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder": //大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}
