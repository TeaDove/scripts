package main

import (
	"service/Service"
)

func main() {
	// _ := Shared.NewSettings()
	appService := Service.AppService{}
	appService.ProcessEvent()
}
