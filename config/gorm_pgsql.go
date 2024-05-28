package config

type Pgsql struct {
	GeneralDB `yaml:",inline" mapStructure:",squash"`
}

// Dsn 基于配置文件获取dsn
func (p *Pgsql) Dsn() string {
	return "host=" + p.Path + "user=" + p.UserName + "password=" + p.Password + "dbname=" + p.DBName + "port=" + p.Port + " " + p.Config
}

// LinkDsn 根据 dbname 生成 dsn
func (p *Pgsql) LinkDsn(dbname string) string {
	return "host=" + p.Path + "user=" + p.UserName + "password=" + p.Password + "dbname=" + dbname + "port=" + p.Port + " " + p.Config
}

func (p *Pgsql) GetLogMode() string {
	return p.LogMode
}
