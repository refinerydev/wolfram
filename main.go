package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/refinerydev/wolfram/src/api/customer"
	"github.com/refinerydev/wolfram/src/database"
	"github.com/refinerydev/wolfram/src/database/entity"
	"github.com/refinerydev/wolfram/src/utils"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&entity.Customer{})

	e := echo.New()

	e.Validator = utils.SetValidator()

	e.File("/public/images/default.png", "src/public/images/default.png")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/api/customers", customer.PostCustomer)
	e.GET("/api/customers", customer.GetCustomers)
	e.PATCH("/api/customers/:id", customer.PatchCustomer)
	e.DELETE("/api/customers/:id", customer.DeleteCustomer)
	e.Logger.Fatal(e.Start(":1323"))
}
