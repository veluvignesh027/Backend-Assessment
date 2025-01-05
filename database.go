package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbinc *gorm.DB

func InitDb() error {

	database := os.Getenv("BLUEPRINT_DB_DATABASE")
	password := os.Getenv("BLUEPRINT_DB_PASSWORD")
	username := os.Getenv("BLUEPRINT_DB_USERNAME")
	port := os.Getenv("BLUEPRINT_DB_PORT")
	host := os.Getenv("BLUEPRINT_DB_HOST")
	schema := os.Getenv("BLUEPRINT_DB_SCHEMA")

	log.Println("Connecting to the database with..")
	log.Println("Host:", host, " User:", username, " Port:", port)
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s", host, port, username, password, database, schema)
	dbinc, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	dbinc.Migrator().DropTable(&Customer{}, &Product{}, &Order{})

	err = dbinc.AutoMigrate(&Customer{}, &Product{}, &Order{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	return err
}
func GetDBInstance() (*gorm.DB, error) {
	var err error
	if dbinc == nil {
		err = InitDb()
	}
	return dbinc, err
}
