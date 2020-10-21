package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func PostHandler(c echo.Context) error {
	log.Println(c.FormValue("type"))
	return c.String(http.StatusOK, "")
}
