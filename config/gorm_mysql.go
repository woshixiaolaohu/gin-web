package config

type Mysql struct {
	GeneralDB `yaml:",inline" mapStructure:",squash"`
}

func (m *Mysql) Dsn() string {
	return m.UserName + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.DBName + "?" + m.Config
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
