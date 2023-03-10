package app_service

import (
	"fmt"
)

type AppService struct {
}

func NewAppService() AppService {
	appService := AppService{}
	return appService
}

func (appService *AppService) ProcessEvent() {
	fmt.Printf("Running\n")
}
