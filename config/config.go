package config

type Config struct {
	Gins   Gins   `mapstructure:"gins" json:"gins" yaml:"gins"`       // gins全局配置
	Server Server `mapstructure:"server" json:"server" yaml:"server"` // 服务配置
}
