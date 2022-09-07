package http

import (
	. "digit-caster/internal/model"
	"fmt"
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

	sum := h.calculationService.Calculate(calcState)
	result := fmt.Sprintf("%g", sum)

	c.String(http.StatusOK, result)
}
