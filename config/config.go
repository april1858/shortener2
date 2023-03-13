package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	BaseURL         string `env:"BASE_URL"`
	ServerAdres     string `env:"SERVER_ADDRESS"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	DatabaseDsn     string `env:"DATABASE_DSN"`
}

var A = flag.String("a", "", "SERVER_ADDRESS is string")
var B = flag.String("b", "", "BASE_URL is string")
var F = flag.String("f", "", "FILE_STORAGE_PATH is string")
var D = flag.String("d", "", "DATABASE_DSN is string")

var cfg = Config{
	ServerAdres:     "localhost:8080",
	BaseURL:  "/",
	FileStoragePath: "",
	DatabaseDsn:     "",
}

func NewConfig() Config {
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
