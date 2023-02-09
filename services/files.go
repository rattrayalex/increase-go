package services

import (
	"bytes"
	"context"
	"fmt"
	"increase/options"
	"increase/pagination"
	"increase/types"
	"io"
	"net/url"
)

type FileService struct {
	Options []options.RequestOption
}

func NewFileService(opts ...options.RequestOption) (r *FileService) {
	r = &FileService{}
	r.Options = opts
	return
}

// To upload a file to Increase, you'll need to send a request of Content-Type
// `multipart/form-data`. The request should contain the file you would like to
// upload, as well as the parameters for creating a file.
func (r *FileService) New(ctx context.Context, body *types.CreateAFileParameters, opts ...options.RequestOption) (res *types.File, err error) {
	opts = append(r.Options[:], opts...)
	b, err := body.MarshalJSON()
	content := io.NopCloser(bytes.NewBuffer(b))
	u, err := url.Parse(fmt.Sprintf("files"))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "POST", u, content, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// Retrieve a File
func (r *FileService) Get(ctx context.Context, file_id string, opts ...options.RequestOption) (res *types.File, err error) {
	opts = append(r.Options[:], opts...)
	u, err := url.Parse(fmt.Sprintf("files/%s", file_id))
	if err != nil {
		return
	}
	cfg := options.NewRequestConfig(ctx, "GET", u, nil, opts...)
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return
	}

	return
}

// List Files
func (r *FileService) List(ctx context.Context, query *types.FileListParams, opts ...options.RequestOption) (res *types.FilesPage, err error) {
	u, err := url.Parse(fmt.Sprintf("files"))
	if err != nil {
		return
	}
	opts = append(r.Options, opts...)
	cfg := options.NewRequestConfig(ctx, "GET", u, nil, opts...)
	res = &types.FilesPage{
		Page: &pagination.Page[types.File]{
			Config:  *cfg,
			Options: opts,
		},
	}
	err = res.Fire()
	return
}
