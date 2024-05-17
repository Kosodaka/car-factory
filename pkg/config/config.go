package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	PostgresDSN string
	Env         string
	HttpPort    string
	HttpHost    string
}

func (c *Config) GetHTTPPort() string {
	return c.HttpPort
}

func (c *Config) GetEnv() string {
	return c.Env
}

func LoadEnv(filenames ...string) error {
	const op = "pkg.config.LoadEnv"
	err := godotenv.Load(filenames...)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func LoadConfig() *Config {
	cfg := &Config{
		PostgresDSN: "",
		Env:         "local",
		HttpHost:    "0.0.0.0",
	}

	postgresDsn := os.Getenv("DSN")
	env := os.Getenv("ENV")
	httpPort := os.Getenv("HTTP_PORT")
	httpHost := os.Getenv("HTTP_HOST")

	if postgresDsn != "" {
		cfg.PostgresDSN = postgresDsn
	}
	if env != "" {
		cfg.Env = env
	}
	if httpPort != "" {
		cfg.HttpPort = httpPort
	}
	if httpHost != "" {
		cfg.HttpHost = httpHost
	}

	return cfg
}
