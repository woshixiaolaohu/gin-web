package config

type Oracle struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (o *Oracle) Dsn() string {
	return "oracle://" + o.Username + ":" + o.Password + "@" + o.Path + ":" + o.Port + "/" + o.DBName + "?" + o.Config
}

func (o *Oracle) GetLogMode() string {
	return o.LogMode
}
