package print_service

import (
	"fmt"
	"service/settings"
)

type PrintService struct {
	settings settings.Settings
}

func NewPrintService(settings settings.Settings) PrintService {
	printService := PrintService{settings}

	fmt.Println(settings.ByCountLimit)
	return printService
}
