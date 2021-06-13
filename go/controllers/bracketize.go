package controllers

import (
	"code_test/go/bracketizer"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type requestBody struct {
	Body string `json:"we_might_have_brackets,omitempty"`
}

// Bracketize - controller for check brackets endpoint
func Bracketize(c echo.Context) error {

	defer c.Request().Body.Close()

	var payload requestBody

	// decode request payload
	err := json.NewDecoder(c.Request().Body).Decode(&payload)
	if err != nil {
		log.Println("failed reading the request payload", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// check if payload is clean or dirty
	clean := bracketizer.SweepForBrackets(payload.Body)

	return c.JSON(http.StatusOK, clean)
}
