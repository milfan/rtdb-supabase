package usecase_module

import (
	app_usecase "github.com/milfan/rtdb-supabase/src/app/usecases"
)

// register your usecase here
type UsecaseModules struct {
	OrderUsecase app_usecase.IOrderUsecase
}

func LoadUsecaseModules() UsecaseModules {
	return UsecaseModules{
		OrderUsecase: app_usecase.NewOrderUsecase(),
	}
}
