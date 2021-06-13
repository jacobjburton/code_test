package main

import (
	"code_test/go/controllers"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())

	g := e.Group("/bracketizer")

	// endpoint to evaluate brackets - locally with docker and follower readme - endpoint bracketizer.test/bracketizer/check
	// body is a text string
	g.POST("/check", controllers.Bracketize)

	// endpoint to evaluate if service is running properly - endpoint bracketizer.test/bracketizer/health
	// if working returns "Health Check: Pass"
	g.GET("/health", controllers.Check)

	port := os.Getenv("APP_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}
