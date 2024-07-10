package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	l "github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	l.Infoln("Artifactory download service is listening at port 8081...")

	e.GET("/api/artifactory/download", HandleGetUpload)

	e.Logger.Fatal(e.Start(":8081"))
}

func HandleGetUpload(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"msg": "Test endpoint"})
}
