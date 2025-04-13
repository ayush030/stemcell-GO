package hornet

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	ServiceAddress     string   `json:"service_address"`
	ServicePort        string   `json:"service_port"`
	DBUrl              string   `json:"db_url"`
	DatabaseConnection *gorm.DB `json:"-"`
}

func (service *Service) New() {
	service.ServicePort = "8080"
	service.ServiceAddress = ""
	service.DBUrl = "postgres://postgres@db/hornet?sslmode=disable"
}

func (service *Service) DB(ctx context.Context) *gorm.DB {
	return service.DatabaseConnection.WithContext(ctx)
}

func (service *Service) ConnectDB() error {
	dsn := service.DBUrl // "host=db user=postgres password=postgres dbname=hornet port=5432 sslmode=disable TimeZone=UTC"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Unable to connect to DB. Error: " + err.Error())
		return err
	}

	service.DatabaseConnection = db

	return nil
}

func (service *Service) End() {
}
