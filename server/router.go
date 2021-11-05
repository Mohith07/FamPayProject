package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"FamPayProject/controller"
)

func NewRouter() *echo.Echo {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.GET("/status", controller.GetStatus)
	router.GET("/list", controller.GetAllVideos)
	router.GET("/search", controller.SearchVideos)

	return router
}
