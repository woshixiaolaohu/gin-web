package config

type Oracle struct {
	GeneralDB `yaml:",inline" mapStructure:",squash"`
}

func (o *Oracle) GetLogMode() string {
	return o.LogMode
}
