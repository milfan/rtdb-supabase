package app_usecase

import "github.com/gin-gonic/gin"

type (
	IOrderUsecase interface {
		Persist(ctx *gin.Context) error
	}
	orderUsecase struct{}
)

// Persist implements IOrderUsecase.
func (o *orderUsecase) Persist(ctx *gin.Context) error {
	panic("unimplemented")
}

func NewOrderUsecase() IOrderUsecase {
	return &orderUsecase{}
}
