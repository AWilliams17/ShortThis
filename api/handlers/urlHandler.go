package handlers

import (
	"github.com/awilliams17/ShortThis/datastore"
	"github.com/labstack/echo"
	"net/http"
	"net/url"
)

type ErrorResponse struct {
	Error string `json:"Error"`
}

type UrlBinding struct {
	Url string `json:"URL"`
}

type UUIDBinding struct {
	UUID string `json:"UUID"`
}

func CreateUUIDFromURL() echo.HandlerFunc {
	return func(c echo.Context) error {
		urlBinding := new(UrlBinding)

		if err := c.Bind(urlBinding); err != nil {
			errorResponse := &ErrorResponse{err.Error()}
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		if _, err := url.ParseRequestURI(urlBinding.Url); err != nil {
			errorResponse := &ErrorResponse{"A valid url is required."}
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		uuid := datastore.AddUrlEntry(urlBinding.Url)
		uuidResponse := &UUIDBinding{uuid,}

		return c.JSON(http.StatusCreated, uuidResponse)
	}
}

func GetURLFromUUID() echo.HandlerFunc {
	return func(c echo.Context) error {
		urlUUID := c.Param("uuid")
		urlEntry, err := datastore.FindUrlEntry(urlUUID)

		if err != nil {
			errorResponse := &ErrorResponse{err.Error()}
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		return c.JSON(http.StatusOK, urlEntry)
	}
}
