package routes

import (
	"net/http"
	"self-payroll/common/structs"
	"self-payroll/routes/position"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AppRoute(router *echo.Echo, db *gorm.DB) {
	router.GET("/", func(ctx echo.Context) error {
		response := &structs.Response{
			Code:    http.StatusOK,
			Message: "Welcome to self payroll api",
		}
		return ctx.JSON(http.StatusOK, response)
	})

	position.InitPositionRoute(router, db)
}
