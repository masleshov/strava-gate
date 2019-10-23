package exception

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func ThrowHTTP(statusCode int, err error) *echo.HTTPError {
	fmt.Errorf(err.Error(), err)
	return echo.NewHTTPError(statusCode, err.Error())
}
