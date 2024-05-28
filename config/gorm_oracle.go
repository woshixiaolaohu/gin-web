package config

type Oracle struct {
	GeneralDB `yaml:",inline" mapStructure:",squash"`
}

func (o *Oracle) Dsn() string {
	return "oracle://" + o.UserName + ":" + o.Password + "@" + o.Path + ":" + o.Port + "/" + o.DBName + "?" + o.Config
}

func (o *Oracle) GetLogMode() string {
	return o.LogMode
}
