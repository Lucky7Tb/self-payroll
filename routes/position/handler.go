package position

import (
	"net/http"
	"self-payroll/common/structs"
	"self-payroll/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitPositionRoute(router *echo.Echo, db *gorm.DB) {
	response := &structs.Response{
		Code:    http.StatusOK,
		Message: "Success get positions",
	}

	router.GET("/positions", func(ctx echo.Context) error {
		limit, errLimit := strconv.Atoi(ctx.QueryParam("limit"))
		if errLimit != nil && limit != 0 {
			response.Code = http.StatusBadRequest
			response.Message = "Limit must be a number"
			response.Data = nil
			return ctx.JSON(http.StatusBadRequest, response)
		}
		skip, errSkip := strconv.Atoi(ctx.QueryParam("skip"))
		if errSkip != nil && skip != 0 {
			response.Code = http.StatusBadRequest
			response.Message = "Skip must be a number"
			response.Data = nil
			return ctx.JSON(http.StatusBadRequest, response)
		}
		var positions []models.Position
		result := db
		if limit != 0 {
			result = result.Limit(limit)
		}
		if skip != 0 {
			result = result.Offset(skip)
		}
		result = result.Select("id", "name", "salary").Find(&positions)
		if result.Error != nil {
			response.Code = http.StatusInternalServerError
			response.Message = "Internal server error"
			return ctx.JSON(http.StatusInternalServerError, response)
		}
		response.Data = positions
		return ctx.JSON(http.StatusOK, response)
	})

	router.GET("/positions/:id", func(ctx echo.Context) error {
		response := &structs.Response{
			Code:    http.StatusOK,
			Message: "Success get position",
		}

		id := ctx.Param("id")
		var position models.Position
		result := db.Select("id", "name", "salary").Where("id = ?", id).Take(&position)
		if result.Error != nil {
			if result.Error.Error() == "record not found" {
				response.Code = http.StatusNotFound
				response.Message = "Position not found!"
				return ctx.JSON(http.StatusNotFound, response)
			}
			response.Code = http.StatusInternalServerError
			response.Message = "Internal server error"
			return ctx.JSON(http.StatusInternalServerError, response)
		}
		response.Data = position
		return ctx.JSON(http.StatusOK, response)
	})
}
