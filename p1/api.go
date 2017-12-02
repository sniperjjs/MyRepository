package p1

import (
	"net/http"

	"github.com/labstack/echo"
)

func Pay(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"name": "xiao"})
}
