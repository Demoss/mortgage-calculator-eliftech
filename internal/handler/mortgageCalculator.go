package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mortgage-calulator-eliftech/internal/command"
	"net/http"
	"strconv"
)

func (h *Handler) GetMortgage(c *gin.Context) {
	initialLoan, err := strconv.Atoi(c.Query("initialLoan"))
	downPayment, err := strconv.Atoi(c.Query("downPayment"))
	bankName := c.Query("bankName")
	input := command.CreateMortgage{
		InitialLoan: initialLoan,
		DownPayment: downPayment,
		BankName:    bankName,
	}
	bank, err := h.services.Bank.GetOne(c, input.BankName)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	monthPayment, err := h.services.Mortgage.GetSuitableBanks(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"payment by month":   monthPayment,
		"mortgage in":        bank.Name,
		"full loan":          input.InitialLoan,
		"mortgage loan term": fmt.Sprintf("%d month", bank.LoanTerm),
	})
}
