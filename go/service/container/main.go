package container

import (
	"service/service"
	"service/settings"
)

type Container struct {
	AppService   service.AppService
	PrintService service.PrintService
}

func NewCombatContainer() Container {
	newSettings := settings.NewSettings()
	printService := service.NewPrintService(newSettings)
	appService := service.NewAppService()

	container := Container{AppService: appService, PrintService: printService}
	return container
}
