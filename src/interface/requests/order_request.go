package interface_requests

import (
	"github.com/gin-gonic/gin"
)

type OrderCreateRequest struct {
	OrderID uint64 `json:"orderID"`
}

func (r *OrderCreateRequest) Validate(ctx *gin.Context) error {

	if err := ctx.ShouldBind(&r); err != nil {
		return err
	}

	return nil
}
