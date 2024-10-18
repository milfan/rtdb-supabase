package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/milfan/neoten-lib/lib_response_gin"
	"github.com/milfan/rtdb-supabase/src/app/modules/usecase_module"
	"github.com/milfan/rtdb-supabase/src/infra/constants"
	interface_requests "github.com/milfan/rtdb-supabase/src/interface/requests"
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
	var req interface_requests.OrderCreateRequest
	if err := req.Validate(ctx); err != nil {
		c.response.HttpError(ctx, err)
		return
	}

	if err := c.usecaseModule.OrderUsecase.Persist(ctx, req); err != nil {
		c.response.HttpError(ctx, err)
		return
	}
	c.response.JSON(
		ctx,
		constants.RESPONSE_CREATED_SUCCESS,
		nil,
		nil,
	)
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
