package config

type Server struct {
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`                            // 服务端口
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"` // 全局路由前缀
}
