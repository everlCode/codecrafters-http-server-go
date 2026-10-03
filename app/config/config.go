package config

import (
	"flag"
)

type Config struct {
	Dirname string
}

func New() *Config {
	cfg := &Config{}

	dir := flag.String("directory", "dir", "Директория хранения файлов")
	flag.Parse()

	cfg.Dirname = *dir

	return cfg
}

