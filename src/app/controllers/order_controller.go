package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/milfan/neoten-lib/lib_response_gin"
)

type (
	ICardStockController interface {
		Persist(ctx *gin.Context)
	}
	cardStockController struct {
		response lib_response_gin.IResponseClient
	}
)

// Persist implements ICardStockController.
func (c *cardStockController) Persist(ctx *gin.Context) {
	panic("unimplemented")
}

func NewCardStockController(
	response lib_response_gin.IResponseClient,
) ICardStockController {
	return &cardStockController{
		response: response,
	}
}
