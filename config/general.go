package config

type EnvConfigGeneral struct {
	RedisHost     string `mapstructure:"redis_host"`
	RedisPort     string `mapstructure:"redis_port"`
	RedisPassword string `mapstructure:"redis_password"`
}

var (
	GeneralConfig EnvConfigGeneral
)
