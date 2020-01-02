package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func CreateUUIDFromURL(c echo.Context) error {
	return c.String(http.StatusOK, "h")
}

func GetURLFromUUID(c echo.Context) error {
	uuid := c.Param("uuid")
	return c.String(http.StatusOK, uuid)
}
