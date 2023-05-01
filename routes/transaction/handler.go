package transaction

import (
	"net/http"
	"self-payroll/common/structs"
	"self-payroll/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitTransactionRoute(router *echo.Echo, db *gorm.DB) {
	response := &structs.Response{
		Code:    http.StatusOK,
		Message: "Success get transaction",
	}

	router.GET("/transactions", func(ctx echo.Context) error {
		var transactions []models.Transaction
		err := db.Select("note", "amount", "type", "created_at").Order("created_at desc").Find(&transactions).Error
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = "Internal server error"
			return ctx.JSON(http.StatusInternalServerError, response)
		}
		response.Data = transactions
		return ctx.JSON(http.StatusOK, response)
	})
}
