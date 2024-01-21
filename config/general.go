package config

type EnvConfigGeneral struct {
	AppPort   string `mapstructure:"app_port"`
	SecretKey string `mapstructure:"secret_key_app"`
}

var (
	GeneralConfig EnvConfigGeneral
)
