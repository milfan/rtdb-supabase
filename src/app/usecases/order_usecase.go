package app_usecase

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	infra_repositories "github.com/milfan/rtdb-supabase/src/infra/repositories"
	interface_requests "github.com/milfan/rtdb-supabase/src/interface/requests"
)

type (
	IOrderUsecase interface {
		Persist(ctx *gin.Context, request interface_requests.OrderCreateRequest) error
	}
	orderUsecase struct {
		rtdbRepository infra_repositories.IRTDBRepository
	}
)

// Persist implements IOrderUsecase.
func (u *orderUsecase) Persist(ctx *gin.Context, request interface_requests.OrderCreateRequest) error {

	type payload struct {
		OrderId   uint64    `json:"order_id"`
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"created_at"`
	}
	data, _ := json.Marshal(payload{
		OrderId:   request.OrderID,
		Status:    "done",
		CreatedAt: time.Now(),
	})

	if err := u.rtdbRepository.Persist(ctx.Request.Context(), string(data)); err != nil {
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
