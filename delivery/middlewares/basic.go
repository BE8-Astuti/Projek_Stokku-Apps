package middlewares

import "github.com/labstack/echo/v4"

func BasicCheck(username, password string, ctx echo.Context) (bool, error) {
	if username == "astuti" && password == "08113601319" {
		return true, nil
	}

	return false, nil
}
