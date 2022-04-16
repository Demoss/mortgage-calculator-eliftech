package handler

import (
	"github.com/gin-gonic/gin"
	"mortgage-calulator-eliftech/internal/command"
	"mortgage-calulator-eliftech/internal/domain"
	"mortgage-calulator-eliftech/internal/query"
	"net/http"
)

func mapToBank(request command.CreateBankRequest) domain.Bank {
	return domain.Bank{
		Name:           request.Name,
		Rate:           request.Rate,
		MaxLoan:        request.MaxLoan,
		MinDownPayment: request.MinDownPayment,
		LoanTerm:       request.LoanTerm,
	}
}
func (h *Handler) CreateBank(c *gin.Context) {
	var input command.CreateBankRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Bank.CreateBank(c, mapToBank(input))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) DeleteBank(c *gin.Context) {
	param := c.Param("name")

	err := h.services.Bank.DeleteBank(c, param)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) Update(c *gin.Context) {
	param := c.Param("name")

	var input command.UpdateBankRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Bank.Update(c, param, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) GetAll(c *gin.Context) {
	banks, err := h.services.Bank.GetAll(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, banks)
}

func (h *Handler) GetOne(c *gin.Context) {
	name := c.Param("name")
	input := query.ChooseBank{Name: name}

	bank, err := h.services.Bank.GetOne(c, input.Name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, bank)
}
