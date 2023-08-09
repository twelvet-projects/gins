package config

type Config struct {
	DbType string // 数据库类型

	Datasource Datasource //数据库配置

	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"` // 多数据源
}
