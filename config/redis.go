package config

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // redis地址
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
}
