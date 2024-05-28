package config

type Mssql struct {
	GeneralDB `yaml:",inline" mapStructure:",squash"`
}

func (m *Mssql) Dsn() string {
	return "sqlserver://" + m.UserName + ":" + m.Password + "@" + m.Path + ":" + m.Port + "?database=" + m.DBName + "&encrypt=disable"
}

func (m *Mssql) GetLogMode() string {
	return m.LogMode
}
