package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Env          string        `mapstructure:"ENV"`
	Server       ServerConfig  `mapstructure:"SERVER"`
	DatabaseDsn  string        `mapstructure:"DATABASE_DSN"`
	TokenSalt    string        `mapstructure:"TOKEN_SALT"`
	PasswordSalt string        `mapstructure:"PASSWORD_SALT"`
	TokenSignKey string        `mapstructure:"TOKEN_SIGN_KEY"`
	TokenTTL     time.Duration `mapstructure:"TOKEN_TTL"`
	Redis        RedisConfig   `mapstructure:"REDIS"`
}

type ServerConfig struct {
	Host    string        `mapstructure:"HOST"`
	Port    int           `mapstructure:"PORT"`
	Timeout time.Duration `mapstructure:"TIMEOUT"`
}

type RedisConfig struct {
	Host string `mapstructure:"HOST"`
	Port int    `mapstructure:"PORT"`
}

func New(path string) *Config {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic("failed read config: " + err.Error())
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic("failed unmarshal config: " + err.Error())
	}

	return &config
}
