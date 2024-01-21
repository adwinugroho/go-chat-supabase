package config

type EnvConfigGeneral struct {
	AppPort string `mapstructure:"app_port"`
}

var (
	GeneralConfig EnvConfigGeneral
)
