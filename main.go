package main

import (
	"fmt"

	"github.com/goPractice/rest"
	"github.com/goPractice/service"
)

func main() {
	serviceContainer := service.NewServiceContainer()

	rest.StartServer(serviceContainer)
	fmt.Println("========== Rest Server Started ============")
	select {}
}
