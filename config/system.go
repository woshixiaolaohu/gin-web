package config

type System struct {
	DBType        string `json:"dbType" yaml:"db_type" mapstructure:"db_type"`     //数据库类型，默认mysql，可选(sqlite|sqlserver|postgresql)
	OssType       string `json:"oss-type" yaml:"oss-type" mapstructure:"oss-type"` // Oss类型
	RouterPrefix  string `json:"router_prefix" yaml:"router_prefix" mapstructure:"router_prefix"`
	Addr          int    `json:"addr" yaml:"addr" mapstructure:"addr"` //端口值
	LimitCountIP  int    `json:"limit_count_ip" yaml:"limit_count_ip" mapstructure:"limit_count_ip"`
	LimitTimeIP   int    `json:"limit_time_ip" yaml:"limit_time_ip" mapstructure:"limit_time_ip"`
	UseMultipoint bool   `json:"use_multipoint" yaml:"use_multipoint" mapstructure:"use_multipoint"` //多点登录拦截
	UseRedis      bool   `json:"use_redis" yaml:"use_redis" mapstructure:"use_redis"`                //使用redis
	UseMongo      bool   `json:"use_mongo" yaml:"use_mongo" mapstructure:"use_mongo"`                //使用mongo
}
