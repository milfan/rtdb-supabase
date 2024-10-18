package app_usecase

import (
	"github.com/gin-gonic/gin"
	infra_repositories "github.com/milfan/rtdb-supabase/src/infra/repositories"
)

type (
	IOrderUsecase interface {
		Persist(ctx *gin.Context) error
	}
	orderUsecase struct {
		rtdbRepository infra_repositories.IRTDBRepository
	}
)

// Persist implements IOrderUsecase.
func (o *orderUsecase) Persist(ctx *gin.Context) error {
	panic("unimplemented")
}

func NewOrderUsecase(
	rtdbRepository infra_repositories.IRTDBRepository,
) IOrderUsecase {
	return &orderUsecase{
		rtdbRepository: rtdbRepository,
	}
}
