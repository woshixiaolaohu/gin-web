package config

type Disk struct {
	MountPoint string `json:"mount_point" yaml:"mount_point" mapstructure:"mount_point"`
}

type DiskList struct {
	Disk `yaml:",inline" mapstructure:",squash"`
}
