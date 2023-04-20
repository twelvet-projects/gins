package config

type Redis struct {
	Enable   bool   `mapstructure:"enable" json:"enable" yaml:"enable"`       // 是否开启（默认关闭）
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // redis地址
	Port     string `mapstructure:"port" json:"port" yaml:"port"`             // 端口
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
}
