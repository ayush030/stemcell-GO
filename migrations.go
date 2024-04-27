package hornet

import (
	"context"
	"fmt"
	"hornet/models"
)

func (service *Service) RunMigrations() error {
	db := service.DB(context.Background())

	// create resource relation
	if err := db.Migrator().CreateTable(&models.Resource{}); err != nil {
		fmt.Println("RunMigrations: Error creating table. Error: " + err.Error())
		return err
	}

	fmt.Println("RunMigrations: Migrations ran successfully")
	return nil
}
