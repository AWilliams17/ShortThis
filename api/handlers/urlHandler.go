package handlers

import (
	"github.com/labstack/echo"
	"github.com/lithammer/shortuuid"
	"net/http"
)

var json map[string]interface{} = map[string]interface{}{}

func CreateUUIDFromURL(c echo.Context) error {
	if err := c.Bind(&json); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// originalURL := json["url"]
	urlUUID := shortuuid.New()[:8]

	return c.String(http.StatusOK, urlUUID)
}

func GetURLFromUUID(c echo.Context) error {
	urlUUID := c.Param("uuid")
	return c.String(http.StatusOK, urlUUID)
}
