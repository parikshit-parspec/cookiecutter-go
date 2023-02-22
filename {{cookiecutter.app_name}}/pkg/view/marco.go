package view

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Marco(c echo.Context) error {
	return c.String(http.StatusOK, "polo!")
}
