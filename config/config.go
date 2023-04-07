package config

import (
	"errors"
	"log"

	"github.com/ory/viper"
)

type Config struct {
	Port           string `json:"port"`
	ReadTimeOut    int    `json:"readTimeout"`
	WriteTimeOut   int    `json:"writeTimeout"`
	MaxHeaderBytes int    `json:"maxHeaderBytes"`
}

func New(filename string) (*Config, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	var c *Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return c, nil
}
