package employee

import (
	"errors"
	"net/http"
	"self-payroll/common/structs"
	"self-payroll/models"
	"self-payroll/routes/employee/dto"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitEmployeeRoute(router *echo.Echo, db *gorm.DB) {
	response := &structs.Response{
		Code:    http.StatusOK,
		Message: "Success get employee",
	}

	router.GET("/employee", func(ctx echo.Context) error {
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
		var employees []models.User
		result := db
		if limit != 0 {
			result = result.Limit(limit)
		}
		if skip != 0 {
			result = result.Offset(skip)
		}
		result = result.Select("id", "employee_id", "position_id", "name", "phone").Preload("Position", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id", "name")
		}).Find(&employees)
		if result.Error != nil {
			response.Code = http.StatusInternalServerError
			response.Message = "Internal server error"
			return ctx.JSON(http.StatusInternalServerError, response)
		}
		response.Data = employees
		return ctx.JSON(http.StatusOK, response)
	})

	router.POST("/employee", func(ctx echo.Context) error {
		response := &structs.Response{
			Code:    http.StatusCreated,
			Message: "Success create employee",
		}
		dto := new(dto.CreateEmployeeDto)
		ctx.Bind(dto)
		if err := ctx.Validate(dto); err != nil {
			return err
		}

		err := db.Select("id").Where("id = ?", dto.PositionId).Take(&models.Position{}).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				response.Code = http.StatusNotFound
				response.Message = "Position not found!"
				return ctx.JSON(http.StatusNotFound, response)
			}
			response.Code = http.StatusInternalServerError
			response.Message = "Internal server error"
			return ctx.JSON(http.StatusInternalServerError, response)
		}

		err = db.Create(&models.User{
			PositionId: dto.PositionId,
			SecretId:   dto.SecretId,
			EmployeeId: dto.EmployeeId,
			Name:       dto.Name,
			Email:      dto.Email,
			Phone:      dto.Phone,
			Address:    dto.Address,
		}).Error
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = "Internal server error"
			return ctx.JSON(http.StatusInternalServerError, response)
		}
		return ctx.JSON(http.StatusCreated, response)
	})

	router.GET("/employee/:id", func(ctx echo.Context) error {
		response := &structs.Response{
			Code:    http.StatusOK,
			Message: "Success get employee",
		}

		id := ctx.Param("id")
		var employee models.User
		err := db.Select("id", "position_id", "employee_id", "secret_id", "name", "phone", "email", "address").Where("id = ?", id).Preload("Position").Take(&employee).Error
		if err != nil {
			if err.Error() == "record not found" {
				response.Code = http.StatusNotFound
				response.Message = "Emplpyee not found!"
				return ctx.JSON(http.StatusNotFound, response)
			}
			response.Code = http.StatusInternalServerError
			response.Message = "Internal server error"
			return ctx.JSON(http.StatusInternalServerError, response)
		}
		response.Data = employee
		return ctx.JSON(http.StatusOK, response)
	})

	router.DELETE("/employee/:id", func(ctx echo.Context) error {
		response := &structs.Response{
			Code:    http.StatusNoContent,
			Message: "Success delete employee",
		}

		id := ctx.Param("id")
		var employee models.User
		err := db.Select("id").Where("id = ?", id).Take(&employee).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				response.Code = http.StatusNotFound
				response.Message = "Employee not found!"
				return ctx.JSON(http.StatusNotFound, response)
			}
			response.Code = http.StatusInternalServerError
			response.Message = "Internal server error"
			return ctx.JSON(http.StatusInternalServerError, response)
		}
		db.Delete(&employee)
		return ctx.JSON(http.StatusNoContent, response)
	})
}
