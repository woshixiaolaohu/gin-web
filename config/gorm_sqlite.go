package config

type Sqlite struct {
	GeneralDB `yaml:",inline" mapStructure:",squash"`
}

func (s *Sqlite) GetLogMode() string {
	return s.LogMode
}
