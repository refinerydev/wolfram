package customer

import (
	"database/sql"
	"fmt"

	"github.com/refinerydev/wolfram/src/database"
)

type repository interface {
	create(customer entity) error
	read() ([]entity, error)
	readById() (entity, error)
	update(customer entity) error
	delete(id uint) error
}

type repo struct {
	db *sql.DB
}

func Repository() *repo {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
	}
	return &repo{db}
}
