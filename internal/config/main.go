package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var pathYaml string
var pathENV string

type Config struct {
	Debug  bool `mapstructure:"debug"`
	Server struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"database"`
	JWT struct {
		// Время жизни токена в минутах
		TokenLifetime uint `mapstructure:"token_lifetime"`
		// Время жизни токена в часах
		RefreshTokenLifeTime uint `mapstructure:"refresh_token_lifetime"`
	} `mapstructure:"jwt"`
	Environment environment `mapstructure:"environment"`
}

type environment struct {
	TokenKey        string `env:"TOKEN_KEY,required"`
	RefreshTokenKey string `env:"REFRESH_TOKEN_KEY,required"`
}

func New(pYaml, pENV string) *Config {
	pathYaml = pYaml
	pathENV = pENV
	config := Config{}
	if err := config.getYamlConfig(); err != nil {
		log.Fatalf("failed to read yaml config: %s", err)
	}

	env, err := getENVConfig()
	if err != nil {
		log.Fatalf("failed to read env config: %s", err)
	}

	config.Environment = env

	return &config
}

func (conf *Config) getYamlConfig() error {
	viper.AddConfigPath(pathYaml)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&conf); err != nil {
		return err
	}

	return nil
}

func getENVConfig() (environment, error) {
	envi := environment{}
	err := godotenv.Load(pathENV)
	if err != nil {
		return envi, err
	}

	err = env.Parse(&envi)
	if err != nil {
		return envi, err
	}

	return envi, err
}
