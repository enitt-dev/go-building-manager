package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func PostHandler(c echo.Context) error {
	location := c.Param("location")
	id := c.Param("id")

	log.Println(location, id)

	params := make(map[string]string)
	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusOK, params["type"])
}
