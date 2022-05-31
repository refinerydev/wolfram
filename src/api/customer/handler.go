package customer

import (
	"net/http"
	"strconv"

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

	if err := appService.AddNew(req); err != nil {
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

func GetCustomers(c echo.Context) error {
	customers, err := appService.GetAll()
	if err != nil {

		statusCode := http.StatusInternalServerError
		response := helper.HttpResponseError(statusCode, err.Error())

		return c.JSON(statusCode, response)
	}

	data := customerListFormatter(customers)

	response := helper.ResponseFormatter(http.StatusOK, "success", "Get uploaded files successful", data)

	return c.JSON(http.StatusOK, response)
}

func PatchCustomer(c echo.Context) error {
	id := c.Param("id")
	customerId, _ := strconv.Atoi(id)

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

	if err := appService.Edit(req, uint(customerId)); err != nil {
		statusCode := http.StatusBadRequest
		response := helper.HttpResponseError(statusCode, err.Error())

		return c.JSON(statusCode, response)
	}

	response := helper.ResponseFormatter(
		http.StatusOK,
		"success",
		"Data updated",
		nil,
	)

	return c.JSON(http.StatusCreated, response)
}

func DeleteCustomer(c echo.Context) error {
	id := c.Param("id")
	customerId, _ := strconv.Atoi(id)

	if err := appService.Remove(uint(customerId)); err != nil {
		statusCode := http.StatusBadRequest
		response := helper.HttpResponseError(statusCode, err.Error())

		return c.JSON(statusCode, response)
	}

	response := helper.ResponseFormatter(
		http.StatusOK,
		"success",
		"Customer removed",
		nil,
	)

	return c.JSON(http.StatusCreated, response)
}
