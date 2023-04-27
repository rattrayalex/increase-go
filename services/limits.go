package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
	"github.com/increase/increase-go/responses"
)

type LimitService struct {
	Options []option.RequestOption
}

func NewLimitService(opts ...option.RequestOption) (r *LimitService) {
	r = &LimitService{}
	r.Options = opts
	return
}

// Create a Limit
func (r *LimitService) New(ctx context.Context, body requests.LimitNewParams, opts ...option.RequestOption) (res *responses.Limit, err error) {
	opts = append(r.Options[:], opts...)
	path := "limits"
	err = option.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Limit
func (r *LimitService) Get(ctx context.Context, limit_id string, opts ...option.RequestOption) (res *responses.Limit, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("limits/%s", limit_id)
	err = option.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Limit
func (r *LimitService) Update(ctx context.Context, limit_id string, body requests.LimitUpdateParams, opts ...option.RequestOption) (res *responses.Limit, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("limits/%s", limit_id)
	err = option.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Limits
func (r *LimitService) List(ctx context.Context, query requests.LimitListParams, opts ...option.RequestOption) (res *responses.Page[responses.Limit], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "limits"
	cfg, err := option.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Limits
func (r *LimitService) ListAutoPaging(ctx context.Context, query requests.LimitListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.Limit] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
