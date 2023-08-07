package config

type Gins struct {
	// gorm
	DbType     string          `mapstructure:"db-type" json:"db-type" yaml:"db-type"`          // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	Datasource Datasource      `mapstructure:"datasource" json:"datasource" yaml:"datasource"` //数据库配置
	DBList     []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`

	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"` // redis配置
	Zap   Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`       // 日志配置
}
