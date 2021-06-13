package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Check - healh check
func Check(c echo.Context) error {
	return c.String(http.StatusOK, "Health Check: Pass")
}
