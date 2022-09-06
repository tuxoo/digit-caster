package http

import (
	"digit-caster/internal/config"
	"digit-caster/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	calculationService service.Calculations
}

func NewHandler(calculationService service.Calculations) *Handler {
	return &Handler{
		calculationService: calculationService,
	}
}

func (h *Handler) InitHandler(cfg config.HTTPConfig) *gin.Engine {
	router := gin.New()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	h.initApi(router)

	return router
}

func (h *Handler) initApi(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		h.initCalculationRoutes(api)
	}
}
