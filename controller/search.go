package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"FamPayProject/model"
	"FamPayProject/service"
)


// GetStatus provides rest endpoint to check the status of the rest server
func GetStatus(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, model.Response{Message: "OK"})
}

func GetAllVideos(c echo.Context) error {
	data, pageNumber, err := service.GetAllVideos(c.QueryParams())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	response := model.Response{
		Message:      "OK",
		ErrorMessage: nil,
		Body:         data,
		Page: pageNumber,
	}
	return c.JSON(http.StatusOK, response)
}

func SearchVideos(c echo.Context) error {
	data, pageNumber, err := service.SearchVideos(c.QueryParams())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	response := model.Response{
		Message:      "OK",
		ErrorMessage: nil,
		Body:         data,
		Page: pageNumber,
	}
	return c.JSON(http.StatusOK, response)
}
