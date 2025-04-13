package main

import (
	"fmt"
	"hornet"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Initializing Hornet service...")

	service := hornet.Service{}
	service.New()

	go func() {
		err := service.ConnectDB()

		for err != nil {
			time.Sleep(5 * time.Second)

			err = service.ConnectDB()
		}

		fmt.Println("Database connection for Hornet is successful")

		_ = service.RunMigrations()
	}()

	e := echo.New()

	service.AddRoutes(e, &service)

	fmt.Println("Firing up the server...")

	err := e.Start(service.ServiceAddress + ":" + service.ServicePort)
	if err != nil {
		fmt.Println("Exiting server. Error: " + err.Error())
	}

	service.End()
}
