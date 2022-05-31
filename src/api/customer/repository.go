package customer

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/refinerydev/wolfram/src/database"
	"github.com/refinerydev/wolfram/src/database/entity"
)

type repository interface {
	create(customer entity.Customer) error
	// read() ([]entity.Customer, error)
	// readById() (entity.Customer, error)
	// update(customer entity.Customer) error
	// delete(id uint) error
}

type repo struct {
	db *gorm.DB
}

func customerRepository() *repo {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
	}
	return &repo{db}
}

func (r *repo) create(customer entity.Customer) error {
	err := r.db.Create(&customer).Error
	if err != nil {
		return err
	}

	return nil
}
