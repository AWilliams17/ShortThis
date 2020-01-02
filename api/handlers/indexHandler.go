package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

type Route struct {
	Route string `json:"Route"`
	Methods string `json:"Method(s)"`
	Parameters map[string]string
	Description string `json:"Description"`
}

func IndexHandler(c echo.Context) error {
	var redirectRoute = Route{
		Route: "/api/redirect:uuid",
		Methods: "GET",
		Description: "Returns the original URL for the given UUID, as well as when said UUID will expire.",
	}

	var createUUIDRoute = Route{
		Route: "/api/create_uuid_key",
		Methods: "POST",
		Parameters: map[string]string{"url": "The url to shorten"},
		Description: "Creates and returns a UUID for a url.",
	}

	var endpoints = map[string][]Route{
		"Endpoints": {redirectRoute, createUUIDRoute},
	}

	return c.JSON(http.StatusOK, endpoints)
}
