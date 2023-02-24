package Shared

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Settings struct {
	ByCountLimit int `env:"by_count_limit"  envDefault:"100"`
}

func NewSettings() Settings {
	err := godotenv.Load(".env")
	if err != nil {
		_ = 0
		// fmt.Println("Error loading .env file, ignoring it")
	}

	settings := Settings{}
	if err := env.Parse(&settings); err != nil {
		fmt.Printf("%+v\n", err)
	}

	return settings
}
