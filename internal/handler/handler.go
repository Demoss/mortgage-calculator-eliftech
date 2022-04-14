package handler

import (
	"github.com/gin-gonic/gin"
	"mortgage-calulator-eliftech/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		bank := api.Group("/bank")
		{
			bank.POST("/", h.CreateBank)
			bank.GET("/", h.GetAll)
			bank.GET("/:name", h.GetOne)
			bank.PUT("/:name", h.Update)
			bank.DELETE("/:name", h.DeleteBank)
		}
		calc := api.Group("/calc")
		{
			calc.GET("/", h.GetMortgage)
		}
	}

	return router
}
