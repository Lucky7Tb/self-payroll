package company

import (
	"net/http"
	"self-payroll/common/structs"
	"self-payroll/models"
	"self-payroll/routes/company/dto"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitCompanyRoute(router *echo.Echo, db *gorm.DB) {
	response := &structs.Response{
		Code:    http.StatusOK,
		Message: "Success get company",
	}

	router.GET("/company", func(ctx echo.Context) error {
		var company models.Company
		result := db.Select("id", "name", "address", "balance").Find(&company)
		if result.Error != nil {
			response.Code = http.StatusInternalServerError
			response.Message = "Internal server error"
			return ctx.JSON(http.StatusInternalServerError, response)
		}
		if result.RowsAffected != 0 {
			response.Data = company
		}

		return ctx.JSON(http.StatusOK, response)
	})

	router.POST("/company", func(ctx echo.Context) error {
		response := &structs.Response{
			Code:    http.StatusCreated,
			Message: "Success save company",
		}

		dto := new(dto.CreateUpdateCompanyDto)
		ctx.Bind(dto)
		if err := ctx.Validate(dto); err != nil {
			return err
		}

		var company models.Company
		err := db.Select("id", "balance").Limit(1).Find(&company).Error
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = "Internal server error"
			return ctx.JSON(http.StatusInternalServerError, response)
		}

		company.Name = dto.Name
		company.Address = dto.Address
		company.Balance = dto.Balance
		company.UpdatedAt = time.Now()
		err = db.Save(&company).Error

		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = "Internal server error"
			return ctx.JSON(http.StatusInternalServerError, response)
		}
		return ctx.JSON(http.StatusCreated, response)
	})

	router.POST("/company/topup", func(ctx echo.Context) error {
		response := &structs.Response{
			Code:    http.StatusOK,
			Message: "Success topup company balance",
		}

		dto := new(dto.TopupCompanyBalanceDto)
		ctx.Bind(dto)
		if err := ctx.Validate(dto); err != nil {
			return err
		}

		var company models.Company
		err := db.Select("id", "balance").First(&company).Error
		if err != nil {
			if err.Error() == "record not found" {
				response.Code = http.StatusNotFound
				response.Message = "Position not found!"
				return ctx.JSON(http.StatusNotFound, response)
			}
			response.Code = http.StatusInternalServerError
			response.Message = "Internal server error"
			return ctx.JSON(http.StatusInternalServerError, response)
		}
		company.Balance = company.Balance + dto.Balance
		err = db.Save(&company).Error
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = "Internal server error"
			return ctx.JSON(http.StatusInternalServerError, response)
		}
		return ctx.JSON(http.StatusOK, response)
	})
}
