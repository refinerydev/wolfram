package customer

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/refinerydev/wolfram/src/database"
	"github.com/refinerydev/wolfram/src/database/entity"
)

type repository interface {
	create(customer entity.Customer) error
	read() ([]entity.Customer, error)
	// readById() (entity.Customer, error)
	update(customer entity.Customer) error
	delete(customer entity.Customer) error
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

func (r *repo) read() ([]entity.Customer, error) {
	customers := []entity.Customer{}

	err := r.db.Find(&customers).Error
	if err != nil {
		return customers, err
	}

	return customers, nil
}

func (r *repo) update(customer entity.Customer) error {
	err := r.db.Model(&customer).Updates(customer).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) delete(customer entity.Customer) error {
	err := r.db.Unscoped().Delete(&customer).Error
	if err != nil {
		return err
	}

	return nil
}
