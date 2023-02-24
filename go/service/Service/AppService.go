package Service

import (
	"fmt"
)

type AppService struct {
}

func (appService AppService) ProcessEvent() {
	fmt.Printf("Running\n")
}
