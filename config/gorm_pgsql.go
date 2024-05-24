package config

type Pgsql struct {
	GeneralDB `yaml:",inline" mapStructure:",squash"`
}

func (p *Pgsql) GetLogMode() string {
	return p.LogMode
}
