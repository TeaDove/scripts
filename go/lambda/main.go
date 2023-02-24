package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Str("status", "cold.start").Send()
}

func main() {
	lambda.Start(HandleRequest)
}
