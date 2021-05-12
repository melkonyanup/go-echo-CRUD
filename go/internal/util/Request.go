package util

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetIDInt64Param(ctx echo.Context) int64 {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err.Error())
	}

	return int64(id)
}
