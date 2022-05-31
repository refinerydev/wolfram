package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/refinerydev/wolfram/src/database"
)

func main() {
	_, err := database.Connect()
	if err != nil {
		fmt.Println(err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
