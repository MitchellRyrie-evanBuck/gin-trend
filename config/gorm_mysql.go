package config

type Mysql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *Mysql) Dsn() string {
	println(m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config)
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
