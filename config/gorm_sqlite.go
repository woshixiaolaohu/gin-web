package config

import "path/filepath"

type Sqlite struct {
	GeneralDB `yaml:",inline" mapStructure:",squash"`
}

func (s *Sqlite) Dsn() string {
	return filepath.Join(s.Path, s.DBName+".db")
}

func (s *Sqlite) GetLogMode() string {
	return s.LogMode
}
