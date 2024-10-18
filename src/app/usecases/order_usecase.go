package app_usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/milfan/neoten-lib/lib_http_request"
)

type (
	IOrderUsecase interface {
		Persist(ctx *gin.Context) error
	}
	orderUsecase struct {
		httpRequestUtils lib_http_request.IHttpRequestUtils
	}
)

// Persist implements IOrderUsecase.
func (o *orderUsecase) Persist(ctx *gin.Context) error {
	panic("unimplemented")
}

func NewOrderUsecase(
	httpRequestUtils lib_http_request.IHttpRequestUtils,
) IOrderUsecase {
	return &orderUsecase{
		httpRequestUtils: httpRequestUtils,
	}
}
