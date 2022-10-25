package config

type Server struct {
	Port int `mapstructure:"port" json:"port" yaml:"port"` // 服务端口
}
