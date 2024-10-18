package app_usecase

import (
	"encoding/json"
	"time"

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
func (u *orderUsecase) Persist(ctx *gin.Context) error {

	type payload struct {
		ID        uint64    `json:"id"`
		OrderId   uint64    `json:"order_id"`
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"created_at"`
	}
	data, _ := json.Marshal(payload{
		ID:        1,
		OrderId:   1,
		Status:    "done",
		CreatedAt: time.Now(),
	})

	if err := u.rtdbRepository.Persist(ctx.Request.Context(), data); err != nil {
		return err
	}
	return nil
}

func NewOrderUsecase(
	rtdbRepository infra_repositories.IRTDBRepository,
) IOrderUsecase {
	return &orderUsecase{
		rtdbRepository: rtdbRepository,
	}
}
