package usecase_module

import (
	"github.com/milfan/rtdb-supabase/src/app/modules/repository_module"
	app_usecase "github.com/milfan/rtdb-supabase/src/app/usecases"
)

// register your usecase here
type UsecaseModules struct {
	OrderUsecase app_usecase.IOrderUsecase
}

func LoadUsecaseModules(
	repositoryModule repository_module.RepositoryModules,
) UsecaseModules {

	return UsecaseModules{
		OrderUsecase: app_usecase.NewOrderUsecase(
			repositoryModule.InfraRepo.RTDBRepository,
		),
	}
}
