package customer

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/refinerydev/wolfram/src/helper"
)

func PostCustomer(c echo.Context) error {
	req := new(request)
	if err := c.Bind(req); err != nil {
		statusCode := http.StatusBadRequest
		response := helper.HttpResponseError(statusCode, err.Error())

		return c.JSON(statusCode, response)
	}

	if err := c.Validate(req); err != nil {
		errors := helper.ErrorResponseFormatter(err)

		statusCode := http.StatusBadRequest
		response := helper.HttpResponseError(statusCode, errors)

		return c.JSON(statusCode, response)
	}

	if err := appService.AddCustomer(req); err != nil {
		statusCode := http.StatusBadRequest
		response := helper.HttpResponseError(statusCode, err.Error())

		return c.JSON(statusCode, response)
	}

	response := helper.ResponseFormatter(
		http.StatusCreated,
		"success",
		"Add new customer successful",
		nil,
	)

	return c.JSON(http.StatusCreated, response)
}
