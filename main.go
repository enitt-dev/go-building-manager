package main

import (
	"go-building-manager/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Post data
	// ZC로부터 받은 데이터를 저장합니다.
	// 우선 출력
	e.POST("/data/:location/:id", handlers.PostHandler)

	// Get event
	// 저장된 이벤트 정보를 아래로 내려줍니다.
	// e.GET("/users/:id", getUser)
	e.GET("/event", handlers.GetEvents)

	e.Logger.Fatal(e.Start(":1323"))
}
