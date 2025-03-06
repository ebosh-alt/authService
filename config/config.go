package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log/slog"
)

func NewConfig() (*Config, error) {
	var cfg Config
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigName("config")
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		slog.Error("fail to read config", err)
		return &cfg, err
	}
	err = v.Unmarshal(&cfg)
	if err != nil {
		slog.Error("", fmt.Errorf("unable to decode config into struct, %w", err))
		return &cfg, err
	}
	if err := godotenv.Load(); err != nil {
		slog.Error("", fmt.Errorf("unable to get env, %w", err))
		return nil, errors.New("unable to get env")
	}
	return &cfg, nil
}
