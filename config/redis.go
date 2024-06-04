package config

type Redis struct {
	Addr         string   `json:"addr" yaml:"addr" mapstructure:"addr"`                            // 服务器地址：端口
	Password     string   `json:"password" yaml:"password" mapstructure:"password"`                // 密码
	DB           int      `json:"db" yaml:"db" mapstructure:"db"`                                  // 单实例模式下 redis 使用那个数据库
	UseCluster   bool     `json:"use_cluster" yaml:"use_cluster" mapstructure:"use_cluster"`       // 是否使用集群模式
	ClusterAddrs []string `json:"cluster_addrs" yaml:"cluster_addrs" mapstructure:"cluster_addrs"` // 集群模式下的节点地址列表
}
