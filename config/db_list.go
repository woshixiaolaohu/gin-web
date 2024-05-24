package config

type DsnProvider interface {
	Dsn() string
}

// GeneralDB 也被 pgsql 和 Mysql 原样使用
type GeneralDB struct {
	Prefix       string `json:"prefix" yaml:"prefix" mapStructure:"prefix"`
	Port         string `json:"port" yaml:"port" mapStructure:"port"`
	Config       string `json:"config" yaml:"config" mapStructure:"config"`          //高级配置
	DBName       string `json:"db_name" yaml:"db_name" mapStructure:"db_name"`       //数据库名称
	UserName     string `json:"user_name" yaml:"user_name" mapStructure:"user_name"` //数据库用户名
	Password     string `json:"password" yaml:"password" mapStructure:"password"`    //数据库密码
	Path         string `json:"path" yaml:"path" mapStructure:"path"`
	Engine       string `json:"engine" yaml:"engine" mapStructure:"engine" default:"InnoDB"`        //数据库引擎,默认InnoDB
	LogMode      string `json:"log_mode" yaml:"log_mode" mapStructure:"log_mode"`                   //是否打开Gorm全局日志
	MaxIdleConns int    `json:"max_idle_conns" yaml:"max_idle_conns" mapStructure:"max_idle_conns"` //空闲中的最大连接数
	MaxOpenConns int    `json:"max_open_conns" yaml:"max_open_conns" mapStructure:"max_open_conns"` //打开到数据库的最大连接数
	Singular     bool   `json:"singular" yaml:"singular" mapStructure:"singular"`                   //是否开始全局禁用复数，true表示开启
	LogZap       bool   `json:"log_zap" yaml:"log_zap" mapStructure:"log_zap"`                      //是否通过zap写入文件
}

type SpecializeDB struct {
	Type      string `json:"type" yaml:"type" mapStructure:"type"`
	AliasName string `json:"alias_name" yaml:"alias_name" mapStructure:"alias_name"`
	GeneralDB `yaml:",inline" mapStructure:",squash"`
	Disable   bool `json:"disable" yaml:"disable" mapStructure:"disable"`
}
