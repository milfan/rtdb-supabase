package infra_repositories

import (
	"context"

	"github.com/milfan/neoten-lib/lib_http_request"
)

type (
	IRTDBRepository interface {
		Persist(ctx context.Context, bodyRequest interface{}) error
	}
	rtdbRepository struct {
		rtdbHost        string
		rtdbClientID    string
		rtdbApiKey      string
		httpRequestUtil lib_http_request.IHttpRequestUtils
	}
)

// Persist implements IRTDBRepository.
func (r *rtdbRepository) Persist(ctx context.Context, bodyRequest interface{}) error {
	panic("unimplemented")
}

func NewRTDBRepository(
	rtdbHost, rtdbClientID, rtdbApiKey string,
	httpRequestUtil lib_http_request.IHttpRequestUtils,
) IRTDBRepository {
	return &rtdbRepository{
		rtdbHost:        rtdbHost,
		rtdbClientID:    rtdbClientID,
		rtdbApiKey:      rtdbApiKey,
		httpRequestUtil: httpRequestUtil,
	}
}
