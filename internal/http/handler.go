package http

import (
	"digit-caster/internal/config"
	"digit-caster/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
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

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		cors.New(corsConfig),
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
