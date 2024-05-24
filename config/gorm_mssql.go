package config

type Mssql struct {
	GeneralDB `yaml:",inline" mapStructure:",squash"`
}

func (m *Mssql) GetLogMode() string {
	return m.LogMode
}
