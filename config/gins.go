package config

type Gins struct {
	Datasource Datasource `mapstructure:"datasource" json:"datasource" yaml:"datasource"` //数据库配置
	Redis      Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`                // redis配置
	Zap        Zap        `mapstructure:"zap" json:"zap" yaml:"zap"`                      // 日志配置
}
