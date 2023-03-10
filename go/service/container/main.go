package container

import (
	"service/service/app_service"
	print2 "service/service/print_service"
	"service/settings"
)

type Container struct {
	AppService   app_service.AppService
	PrintService print2.PrintService
}

func NewCombatContainer() Container {
	settings := settings.NewSettings()
	printService := print2.NewPrintService(settings)
	appService := app_service.NewAppService()

	container := Container{AppService: appService, PrintService: printService}
	return container
}
