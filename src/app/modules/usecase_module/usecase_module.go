package usecase_module

import (
	"github.com/milfan/neoten-lib/lib_http_request"
	app_usecase "github.com/milfan/rtdb-supabase/src/app/usecases"
)

// register your usecase here
type UsecaseModules struct {
	OrderUsecase app_usecase.IOrderUsecase
}

func LoadUsecaseModules(
	httpRequestUtils lib_http_request.IHttpRequestUtils,
) UsecaseModules {
	return UsecaseModules{
		OrderUsecase: app_usecase.NewOrderUsecase(
			httpRequestUtils,
		),
	}
}
