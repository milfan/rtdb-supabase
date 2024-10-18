package controller_modules

import (
	"github.com/milfan/neoten-lib/lib_response_gin"
	"github.com/milfan/rtdb-supabase/src/app/controllers"
	"github.com/milfan/rtdb-supabase/src/app/modules/usecase_module"
)

type (
	ControllerDicts struct {
		OrderController controllers.IOrderController
	}
)

func LoadControllers(
	response lib_response_gin.IResponseClient,
	usecasesModules usecase_module.UsecaseModules,
) ControllerDicts {
	return ControllerDicts{
		OrderController: controllers.NewOrderController(response, usecasesModules),
	}
}
