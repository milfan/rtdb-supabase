package infra_repositories

import (
	"bytes"
	"context"
	"fmt"

	"github.com/milfan/neoten-lib/lib_constants"
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
	headers := map[string]string{
		lib_constants.HTTP_CONTENT_TYPE: lib_constants.HTTP_APPLICATION_JSON,
		"apikey":                        r.rtdbApiKey,
		"Authorization":                 fmt.Sprintf("Bearer %s", r.rtdbApiKey),
		"Prefer":                        "return=minimal",
	}

	if err := r.httpRequest(
		ctx,
		lib_http_request.POST,
		headers,
		bodyRequest,
	); err != nil {
		return err
	}
	return nil
}

func (r *rtdbRepository) httpRequest(
	ctx context.Context,
	httpMethod lib_http_request.HttpMethod,
	headers map[string]string,
	bodyRequest interface{},
) error {

	bodyContentStr := bodyRequest.(string)

	_, _, err := r.httpRequestUtil.DoRequest(
		ctx,
		nil,
		r.rtdbHost,
		httpMethod,
		headers,
		nil,
		bytes.NewBuffer([]byte(bodyContentStr)),
	)
	if err != nil {
		return err
	}
	return nil
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
