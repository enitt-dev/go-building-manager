package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type event struct {
	EventType   string `json:"event_type"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}

type events struct {
	IsEvent   bool    `json:"is_event"`
	EventList []event `json:"event_list"`
}

func GetEvents(c echo.Context) error {

	// dummy data
	j := events{
		IsEvent: true,
		EventList: []event{
			{
				EventType:   "Reduce",
				Description: "감축명령 1",
				CreatedAt:   time.Now().Add(-time.Second * 30).Unix(),
			},
			{
				EventType:   "ESS OFF",
				Description: "ESS 꺼버려",
				CreatedAt:   time.Now().Unix(),
			},
		},
	}

	return c.JSON(http.StatusOK, j)
}
