package config

type Datasource struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *Datasource) GetLogMode() string {
	return m.LogMode
}
