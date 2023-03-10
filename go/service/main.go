package main

import (
	"service/container"
)

func main() {
	container := container.NewCombatContainer()
	container.AppService.ProcessEvent()
}
