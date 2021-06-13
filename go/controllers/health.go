package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthCheck - healh check
func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Health Check: Pass")
}
