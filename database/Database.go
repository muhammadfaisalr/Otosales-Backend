package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func config() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
		return nil, err
	}

	println("Success to Connect DB.")
	return db, nil
}

func Connect() (*gorm.DB, error) {
	db, err := config()

	if err != nil {
		return nil, err
	}

	return db, err
}

const (
	host     = "localhost"
	port     = 1234
	user     = "admin"
	password = "admin"
	dbname   = "dev_otosales"
)

const (
	POST   = "POST"
	GET    = "GET"
	PUT    = "PUT"
	DELETE = "DELETE"
)
