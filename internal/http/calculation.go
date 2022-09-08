package http

import (
	. "digit-caster/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initCalculationRoutes(api *gin.RouterGroup) {
	api.POST("/calculation", h.addition)
}

func (h *Handler) addition(c *gin.Context) {
	var calcState CalcState
	if err := c.ShouldBindJSON(&calcState); err != nil {
		return
	}

	result := h.calculationService.Calculate(calcState)

	c.String(http.StatusOK, result)
}
