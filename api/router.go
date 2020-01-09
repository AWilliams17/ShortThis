package api

import (
	"github.com/awilliams17/ShortThis/api/handlers"
	"github.com/labstack/echo"
)


func CreateRouter() *echo.Echo {
	e := echo.New()

	e.GET("/api", handlers.IndexHandler)
	e.GET("/api/redirect/:uuid", handlers.GetURLFromUUID())
	e.POST("/api/create_uuid_key", handlers.CreateUUIDFromURL())

	return e
}
