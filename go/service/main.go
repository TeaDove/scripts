package main

import (
	"service/container"
)

func main() {
	combatContainer := container.NewCombatContainer()
	combatContainer.AppService.ProcessEvent()
}
