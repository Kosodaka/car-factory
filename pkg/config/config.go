package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Env        string
	HttpPort   string
	HttpHost   string
	SqlitePath string
}

func (c *Config) GetHTTPPort() string {
	return c.HttpPort
}

func (c *Config) GetEnv() string {
	return c.Env
}

func LoadEnv(filenames ...string) error {
	const op = "pkg.Config.LoadEnv"
	err := godotenv.Load(filenames...)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func GetConfig() *Config {
	cfg := &Config{
		Env:      "local",
		HttpHost: "localhost",
	}

	sqlitePath := os.Getenv("SQL_PATH")
	env := os.Getenv("ENV")
	httpPort := os.Getenv("HTTP_PORT")
	httpHost := os.Getenv("HTTP_HOST")

	if env != "" {
		cfg.Env = env
	}
	if httpPort != "" {
		cfg.HttpPort = httpPort
	}
	if httpHost != "" {
		cfg.HttpHost = httpHost
	}
	if sqlitePath != "" {
		cfg.SqlitePath = sqlitePath
	}

	return cfg
}
