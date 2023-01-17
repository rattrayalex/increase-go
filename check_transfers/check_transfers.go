package check_transfers

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type CheckTransferService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewCheckTransferService(requester core.Requester) (r *CheckTransferService) {
	r = &CheckTransferService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedCheckTransferService struct {
	CheckTransfers *CheckTransferService
}

func (r *PreloadedCheckTransferService) Init(service *CheckTransferService) {
	r.CheckTransfers = service
}

func NewPreloadedCheckTransferService(service *CheckTransferService) (r *PreloadedCheckTransferService) {
	r = &PreloadedCheckTransferService{}
	r.Init(service)
	return
}

//
type CheckTransfer struct {
	// The identifier of the Account from which funds will be transferred.
	AccountID *string `json:"account_id"`
	// The street address of the check's destination.
	AddressLine1 *string `json:"address_line1"`
	// The second line of the address of the check's destination.
	AddressLine2 *string `json:"address_line2"`
	// The city of the check's destination.
	AddressCity *string `json:"address_city"`
	// The state of the check's destination.
	AddressState *string `json:"address_state"`
	// The postal code of the check's destination.
	AddressZip *string `json:"address_zip"`
	// The transfer amount in USD cents.
	Amount *int `json:"amount"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt *string `json:"created_at"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
	// currency.
	Currency *CheckTransferCurrency `json:"currency"`
	// The Check transfer's identifier.
	ID *string `json:"id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was mailed.
	MailedAt *string `json:"mailed_at"`
	// The descriptor that is printed on the check.
	Message *string `json:"message"`
	// The name that will be printed on the check.
	RecipientName *string `json:"recipient_name"`
	// The lifecycle status of the transfer.
	Status *CheckTransferStatus `json:"status"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the check was submitted.
	SubmittedAt *string `json:"submitted_at"`
	// After the transfer is submitted, this will contain supplemental details.
	Submission *CheckTransferSubmission `json:"submission"`
	// If the transfer was created from a template, this will be the template's ID.
	TemplateID *string `json:"template_id"`
	// The ID for the transaction caused by the transfer.
	TransactionID *string `json:"transaction_id"`
	// After a stop-payment is requested on the check, this will contain supplemental
	// details.
	StopPaymentRequest *CheckTransferStopPaymentRequest `json:"stop_payment_request"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer`.
	Type *CheckTransferType `json:"type"`
}

// The identifier of the Account from which funds will be transferred.
func (r *CheckTransfer) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The street address of the check's destination.
func (r *CheckTransfer) GetAddressLine1() (AddressLine1 string) {
	if r != nil && r.AddressLine1 != nil {
		AddressLine1 = *r.AddressLine1
	}
	return
}

// The second line of the address of the check's destination.
func (r *CheckTransfer) GetAddressLine2() (AddressLine2 string) {
	if r != nil && r.AddressLine2 != nil {
		AddressLine2 = *r.AddressLine2
	}
	return
}

// The city of the check's destination.
func (r *CheckTransfer) GetAddressCity() (AddressCity string) {
	if r != nil && r.AddressCity != nil {
		AddressCity = *r.AddressCity
	}
	return
}

// The state of the check's destination.
func (r *CheckTransfer) GetAddressState() (AddressState string) {
	if r != nil && r.AddressState != nil {
		AddressState = *r.AddressState
	}
	return
}

// The postal code of the check's destination.
func (r *CheckTransfer) GetAddressZip() (AddressZip string) {
	if r != nil && r.AddressZip != nil {
		AddressZip = *r.AddressZip
	}
	return
}

// The transfer amount in USD cents.
func (r *CheckTransfer) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the transfer was created.
func (r *CheckTransfer) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the check's
// currency.
func (r *CheckTransfer) GetCurrency() (Currency CheckTransferCurrency) {
	if r != nil && r.Currency != nil {
		Currency = *r.Currency
	}
	return
}

// The Check transfer's identifier.
func (r *CheckTransfer) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the check was mailed.
func (r *CheckTransfer) GetMailedAt() (MailedAt string) {
	if r != nil && r.MailedAt != nil {
		MailedAt = *r.MailedAt
	}
	return
}

// The descriptor that is printed on the check.
func (r *CheckTransfer) GetMessage() (Message string) {
	if r != nil && r.Message != nil {
		Message = *r.Message
	}
	return
}

// The name that will be printed on the check.
func (r *CheckTransfer) GetRecipientName() (RecipientName string) {
	if r != nil && r.RecipientName != nil {
		RecipientName = *r.RecipientName
	}
	return
}

// The lifecycle status of the transfer.
func (r *CheckTransfer) GetStatus() (Status CheckTransferStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the check was submitted.
func (r *CheckTransfer) GetSubmittedAt() (SubmittedAt string) {
	if r != nil && r.SubmittedAt != nil {
		SubmittedAt = *r.SubmittedAt
	}
	return
}

// After the transfer is submitted, this will contain supplemental details.
func (r *CheckTransfer) GetSubmission() (Submission CheckTransferSubmission) {
	if r != nil && r.Submission != nil {
		Submission = *r.Submission
	}
	return
}

// If the transfer was created from a template, this will be the template's ID.
func (r *CheckTransfer) GetTemplateID() (TemplateID string) {
	if r != nil && r.TemplateID != nil {
		TemplateID = *r.TemplateID
	}
	return
}

// The ID for the transaction caused by the transfer.
func (r *CheckTransfer) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

// After a stop-payment is requested on the check, this will contain supplemental
// details.
func (r *CheckTransfer) GetStopPaymentRequest() (StopPaymentRequest CheckTransferStopPaymentRequest) {
	if r != nil && r.StopPaymentRequest != nil {
		StopPaymentRequest = *r.StopPaymentRequest
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer`.
func (r *CheckTransfer) GetType() (Type CheckTransferType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type CheckTransferCurrency string

const (
	CheckTransferCurrencyCad CheckTransferCurrency = "CAD"
	CheckTransferCurrencyChf CheckTransferCurrency = "CHF"
	CheckTransferCurrencyEur CheckTransferCurrency = "EUR"
	CheckTransferCurrencyGbp CheckTransferCurrency = "GBP"
	CheckTransferCurrencyJpy CheckTransferCurrency = "JPY"
	CheckTransferCurrencyUsd CheckTransferCurrency = "USD"
)

type CheckTransferStatus string

const (
	CheckTransferStatusPendingApproval   CheckTransferStatus = "pending_approval"
	CheckTransferStatusPendingSubmission CheckTransferStatus = "pending_submission"
	CheckTransferStatusSubmitting        CheckTransferStatus = "submitting"
	CheckTransferStatusSubmitted         CheckTransferStatus = "submitted"
	CheckTransferStatusPendingMailing    CheckTransferStatus = "pending_mailing"
	CheckTransferStatusMailed            CheckTransferStatus = "mailed"
	CheckTransferStatusCanceled          CheckTransferStatus = "canceled"
	CheckTransferStatusDeposited         CheckTransferStatus = "deposited"
	CheckTransferStatusStopped           CheckTransferStatus = "stopped"
	CheckTransferStatusRejected          CheckTransferStatus = "rejected"
	CheckTransferStatusRequiresAttention CheckTransferStatus = "requires_attention"
)

//
type CheckTransferSubmission struct {
	// The identitying number of the check.
	CheckNumber *string `json:"check_number"`
}

// The identitying number of the check.
func (r *CheckTransferSubmission) GetCheckNumber() (CheckNumber string) {
	if r != nil && r.CheckNumber != nil {
		CheckNumber = *r.CheckNumber
	}
	return
}

//
type CheckTransferStopPaymentRequest struct {
	// The ID of the check transfer that was stopped.
	TransferID *string `json:"transfer_id"`
	// The transaction ID of the corresponding credit transaction.
	TransactionID *string `json:"transaction_id"`
	// The time the stop-payment was requested.
	RequestedAt *string `json:"requested_at"`
	// A constant representing the object's type. For this resource it will always be
	// `check_transfer_stop_payment_request`.
	Type *CheckTransferStopPaymentRequestType `json:"type"`
}

// The ID of the check transfer that was stopped.
func (r *CheckTransferStopPaymentRequest) GetTransferID() (TransferID string) {
	if r != nil && r.TransferID != nil {
		TransferID = *r.TransferID
	}
	return
}

// The transaction ID of the corresponding credit transaction.
func (r *CheckTransferStopPaymentRequest) GetTransactionID() (TransactionID string) {
	if r != nil && r.TransactionID != nil {
		TransactionID = *r.TransactionID
	}
	return
}

// The time the stop-payment was requested.
func (r *CheckTransferStopPaymentRequest) GetRequestedAt() (RequestedAt string) {
	if r != nil && r.RequestedAt != nil {
		RequestedAt = *r.RequestedAt
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `check_transfer_stop_payment_request`.
func (r *CheckTransferStopPaymentRequest) GetType() (Type CheckTransferStopPaymentRequestType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type CheckTransferStopPaymentRequestType string

const (
	CheckTransferStopPaymentRequestTypeCheckTransferStopPaymentRequest CheckTransferStopPaymentRequestType = "check_transfer_stop_payment_request"
)

type CheckTransferType string

const (
	CheckTransferTypeCheckTransfer CheckTransferType = "check_transfer"
)

type CreateACheckTransferParameters struct {
	// The identifier for the account that will send the transfer.
	AccountID *string `json:"account_id"`
	// The street address of the check's destination.
	AddressLine1 *string `json:"address_line1"`
	// The second line of the address of the check's destination.
	AddressLine2 *string `json:"address_line2,omitempty"`
	// The city of the check's destination.
	AddressCity *string `json:"address_city"`
	// The state of the check's destination.
	AddressState *string `json:"address_state"`
	// The postal code of the check's destination.
	AddressZip *string `json:"address_zip"`
	// The transfer amount in cents.
	Amount *int `json:"amount"`
	// The descriptor that will be printed on the check.
	Message *string `json:"message"`
	// The name that will be printed on the check.
	RecipientName *string `json:"recipient_name"`
}

// The identifier for the account that will send the transfer.
func (r *CreateACheckTransferParameters) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The street address of the check's destination.
func (r *CreateACheckTransferParameters) GetAddressLine1() (AddressLine1 string) {
	if r != nil && r.AddressLine1 != nil {
		AddressLine1 = *r.AddressLine1
	}
	return
}

// The second line of the address of the check's destination.
func (r *CreateACheckTransferParameters) GetAddressLine2() (AddressLine2 string) {
	if r != nil && r.AddressLine2 != nil {
		AddressLine2 = *r.AddressLine2
	}
	return
}

// The city of the check's destination.
func (r *CreateACheckTransferParameters) GetAddressCity() (AddressCity string) {
	if r != nil && r.AddressCity != nil {
		AddressCity = *r.AddressCity
	}
	return
}

// The state of the check's destination.
func (r *CreateACheckTransferParameters) GetAddressState() (AddressState string) {
	if r != nil && r.AddressState != nil {
		AddressState = *r.AddressState
	}
	return
}

// The postal code of the check's destination.
func (r *CreateACheckTransferParameters) GetAddressZip() (AddressZip string) {
	if r != nil && r.AddressZip != nil {
		AddressZip = *r.AddressZip
	}
	return
}

// The transfer amount in cents.
func (r *CreateACheckTransferParameters) GetAmount() (Amount int) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

// The descriptor that will be printed on the check.
func (r *CreateACheckTransferParameters) GetMessage() (Message string) {
	if r != nil && r.Message != nil {
		Message = *r.Message
	}
	return
}

// The name that will be printed on the check.
func (r *CreateACheckTransferParameters) GetRecipientName() (RecipientName string) {
	if r != nil && r.RecipientName != nil {
		RecipientName = *r.RecipientName
	}
	return
}

type ListCheckTransfersQuery struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int `query:"limit"`
	// Filter Check Transfers to those that originated from the specified Account.
	AccountID *string                           `query:"account_id"`
	CreatedAt *ListCheckTransfersQueryCreatedAt `query:"created_at"`
}

// Return the page of entries after this one.
func (r *ListCheckTransfersQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListCheckTransfersQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Check Transfers to those that originated from the specified Account.
func (r *ListCheckTransfersQuery) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

func (r *ListCheckTransfersQuery) GetCreatedAt() (CreatedAt ListCheckTransfersQueryCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

type ListCheckTransfersQueryCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *string `json:"after,omitempty"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *string `json:"before,omitempty"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *string `json:"on_or_after,omitempty"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *string `json:"on_or_before,omitempty"`
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListCheckTransfersQueryCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListCheckTransfersQueryCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListCheckTransfersQueryCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListCheckTransfersQueryCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

//
type CheckTransferList struct {
	// The contents of the list.
	Data *[]CheckTransfer `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// The contents of the list.
func (r *CheckTransferList) GetData() (Data []CheckTransfer) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *CheckTransferList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *CheckTransferService) Create(ctx context.Context, body *CreateACheckTransferParameters, opts ...*core.RequestOpts) (res *CheckTransfer, err error) {
	err = r.post(
		ctx,
		"/check_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *PreloadedCheckTransferService) Create(ctx context.Context, body *CreateACheckTransferParameters, opts ...*core.RequestOpts) (res *CheckTransfer, err error) {
	err = r.CheckTransfers.post(
		ctx,
		"/check_transfers",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *CheckTransferService) Retrieve(ctx context.Context, check_transfer_id string, opts ...*core.RequestOpts) (res *CheckTransfer, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/check_transfers/%s", check_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PreloadedCheckTransferService) Retrieve(ctx context.Context, check_transfer_id string, opts ...*core.RequestOpts) (res *CheckTransfer, err error) {
	err = r.CheckTransfers.get(
		ctx,
		fmt.Sprintf("/check_transfers/%s", check_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

type CheckTransfersPage struct {
	*pagination.Page[CheckTransfer]
}

func (r *CheckTransfersPage) CheckTransfer() *CheckTransfer {
	return r.Current()
}

func (r *CheckTransfersPage) GetNextPage() (*CheckTransfersPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &CheckTransfersPage{page}, nil
	}
}

func (r *CheckTransferService) List(ctx context.Context, query *ListCheckTransfersQuery, opts ...*core.RequestOpts) (res *CheckTransfersPage, err error) {
	page := &CheckTransfersPage{
		Page: &pagination.Page[CheckTransfer]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/check_transfers",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *PreloadedCheckTransferService) List(ctx context.Context, query *ListCheckTransfersQuery, opts ...*core.RequestOpts) (res *CheckTransfersPage, err error) {
	page := &CheckTransfersPage{
		Page: &pagination.Page[CheckTransfer]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/check_transfers",
			},
			Requester: r.CheckTransfers.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *CheckTransferService) StopPayment(ctx context.Context, check_transfer_id string, opts ...*core.RequestOpts) (res *CheckTransfer, err error) {
	err = r.post(
		ctx,
		fmt.Sprintf("/check_transfers/%s/stop_payment", check_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PreloadedCheckTransferService) StopPayment(ctx context.Context, check_transfer_id string, opts ...*core.RequestOpts) (res *CheckTransfer, err error) {
	err = r.CheckTransfers.post(
		ctx,
		fmt.Sprintf("/check_transfers/%s/stop_payment", check_transfer_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}
