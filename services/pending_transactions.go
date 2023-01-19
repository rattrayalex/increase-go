package services

import "context"
import "increase/core"
import "increase/types"
import "fmt"
import "increase/pagination"

type PendingTransactionService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewPendingTransactionService(requester core.Requester) (r *PendingTransactionService) {
	r = &PendingTransactionService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *PendingTransactionService) Retrieve(ctx context.Context, pending_transaction_id string, opts ...*core.RequestOpts) (res *types.PendingTransaction, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/pending_transactions/%s", pending_transaction_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

func (r *PendingTransactionService) List(ctx context.Context, query *types.ListPendingTransactionsQuery, opts ...*core.RequestOpts) (res *types.PendingTransactionsPage, err error) {
	page := &types.PendingTransactionsPage{
		Page: &pagination.Page[types.PendingTransaction]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/pending_transactions",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}