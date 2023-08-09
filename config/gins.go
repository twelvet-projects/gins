package config

type Gins struct {
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"` // redis配置
	Zap   Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`       // 日志配置
}
