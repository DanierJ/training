package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

/*
	USER = myUser
	PASS = myPass
	HOST = myHost
	PORT = myPort
	DBNAME = myDBNAME

*/
const (
	USER   = "postgres"
	PASS   = "admin123"
	HOST   = "localhost"
	PORT   = 5432
	DBNAME = "tododb"
)

// Connect connects to the db
func Connect() *gorm.DB {
	URL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", HOST, PORT, USER, DBNAME, PASS)

	db, err := gorm.Open("postgres", URL)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}
