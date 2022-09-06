package http

import "github.com/gin-gonic/gin"

func (h *Handler) initCalculationRoutes(api *gin.RouterGroup) {
	calculation := api.Group("/calculation")
	{
		calculation.GET("/sum/", h.summation)
		calculation.GET("/sub/", h.subtraction)
		calculation.GET("/mul/", h.multiple)
		calculation.GET("/div/", h.division)
		calculation.GET("/squ/", h.square)
	}
}

func (h *Handler) summation(c *gin.Context) {

}

func (h *Handler) subtraction(c *gin.Context) {

}

func (h *Handler) multiple(c *gin.Context) {

}

func (h *Handler) division(c *gin.Context) {

}

func (h *Handler) square(c *gin.Context) {

}
