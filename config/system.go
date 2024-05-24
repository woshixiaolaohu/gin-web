package config

type System struct {
	DBType        string `json:"db_type" yaml:"db_type" mapStructure:"db_type"`    //数据库类型，默认mysql，可选(sqlite|sqlserver|postgresql)
	OssType       string `json:"oss-type" yaml:"oss-type" mapStructure:"oss-type"` // Oss类型
	RouterPrefix  string `json:"router_prefix" yaml:"router_prefix" mapStructure:"router_prefix"`
	Addr          int    `json:"addr" yaml:"addr" mapStructure:"addr"` //端口值
	LimitCountIP  int    `json:"limit_count_ip" yaml:"limit_count_ip" mapStructure:"limit_count_ip"`
	LimitTimeIP   int    `json:"limit_time_ip" yaml:"limit_time_ip" mapStructure:"limit_time_ip"`
	UseMultipoint bool   `json:"use_multipoint" yaml:"use_multipoint" mapStructure:"use_multipoint"` //多点登录拦截
	UseRedis      bool   `json:"use_redis" yaml:"use_redis" mapStructure:"use_redis"`                //使用redis
	UseMongo      bool   `json:"use_mongo" yaml:"use_mongo" mapStructure:"use_mongo"`                //使用mongo
}
