package company

import (
	"errors"
	"net/http"
	"self-payroll/common/structs"
	"self-payroll/models"
	"self-payroll/routes/company/dto"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitCompanyRoute(router *echo.Echo, db *gorm.DB) {
	router.GET("/company", func(ctx echo.Context) error {
		response := &structs.Response{
			Code:    http.StatusOK,
			Message: "Success get company",
		}

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

		if company.Id != "" {
			err = db.Model(&company).Updates(&models.Company{Name: dto.Name, Address: dto.Address, Balance: dto.Balance}).Where("id = ?", company.Id).Error
		} else {
			err = db.Model(&company).Create(&models.Company{Name: dto.Name, Address: dto.Address, Balance: dto.Balance}).Error
		}

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

		err := db.Transaction(func(transaction *gorm.DB) error {
			var company models.Company
			var err error

			err = db.Select("id", "balance").First(&company).Error
			if err != nil {
				return err
			}

			err = transaction.Create(&models.Transaction{Type: "Kredit", Amount: dto.Balance, Note: "Topup balance company"}).Error
			if err != nil {
				return err
			}

			err = transaction.Model(&company).Update("Balance", company.Balance+dto.Balance).Error
			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				response.Code = http.StatusNotFound
				response.Message = "Position not found!"
				return ctx.JSON(http.StatusNotFound, response)
			}

			if errors.Is(err, gorm.ErrInvalidTransaction) {
				response.Code = http.StatusInternalServerError
				response.Message = "Failed to update company balance and insert transaction history"
				return ctx.JSON(http.StatusInternalServerError, response)
			}

			response.Code = http.StatusInternalServerError
			response.Message = "Internal server error"
			return ctx.JSON(http.StatusInternalServerError, response)
		}

		return ctx.JSON(http.StatusOK, response)
	})
}
