//go:build pro

package config

var Config = config{
	DB: DBConfig{
		Dsn: "root:123456@tcp(127.0.0.1:3306)/test",
	},
	Redis: RedisConfig{
		Addr: "127.0.0.1:6379",
	},
}
