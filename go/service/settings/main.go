package settings

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Settings struct {
	ByCountLimit int `env:"by_count_limit"  envDefault:"100"`
}

func NewSettings() Settings {
	_ = godotenv.Load(".env")

	settings := Settings{}
	if err := env.Parse(&settings); err != nil {
		panic(err)
	}

	return settings
}
