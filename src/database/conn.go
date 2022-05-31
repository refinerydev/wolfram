package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect() (*gorm.DB, error) {
	dbConfig := DbConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Database,
		dbConfig.Port,
		dbConfig.SslMode,
	)

	// dsn := "host=localhost user=postgres password= dbname=jersey_dev port=5432 sslmode=disable"

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Could not connect to database :%v", err)
	}

	return db, nil
}
