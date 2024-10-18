package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/milfan/neoten-lib/lib_response_gin"
	"github.com/milfan/rtdb-supabase/src/app/modules/usecase_module"
)

type (
	IOrderController interface {
		Persist(ctx *gin.Context)
	}
	orderController struct {
		response      lib_response_gin.IResponseClient
		usecaseModule usecase_module.UsecaseModules
	}
)

// Persist implements IOrderController.
func (c *orderController) Persist(ctx *gin.Context) {
	panic("unimplemented")
}

func NewOrderController(
	response lib_response_gin.IResponseClient,
	usecaseModule usecase_module.UsecaseModules,
) IOrderController {
	return &orderController{
		response:      response,
		usecaseModule: usecaseModule,
	}
}
