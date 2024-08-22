//go:build !pro

package config

var Config = config{
	DB: DBConfig{
		Dsn: "root:CQGWiRshWb@tcp(192.168.112.24:3306)/test?charset=utf8&parseTime=True&loc=Local",
	},
	Redis: RedisConfig{
		Addr: "192.168.112.24:6379",
		DB:   4,
	},
}
