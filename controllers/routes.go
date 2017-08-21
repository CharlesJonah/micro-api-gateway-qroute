package controllers

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Setup() {

	environment := os.Getenv("GO_ENV")
	router := echo.New()
	router.Use(middleware.CORS())
	router.Use(middleware.Recover())

	coreResource := CoreResource()
	pulse := router.Group("/api/v1/qroute")
	{
		pulse.POST("/servicevendor", coreResource.CreateServiceVendor())
	}

	router.GET("/", func(c echo.Context) error {
		return c.JSON(200, echo.Map{
			"message": "Welcome to qroute's system's entry point",
		})
	})

	return router
}
