package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// InboundWireDrawdownRequestService contains methods and other services that help
// with interacting with the increase API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewInboundWireDrawdownRequestService] method instead.
type InboundWireDrawdownRequestService struct {
	Options []option.RequestOption
}

// NewInboundWireDrawdownRequestService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewInboundWireDrawdownRequestService(opts ...option.RequestOption) (r *InboundWireDrawdownRequestService) {
	r = &InboundWireDrawdownRequestService{}
	r.Options = opts
	return
}

// Retrieve an Inbound Wire Drawdown Request
func (r *InboundWireDrawdownRequestService) Get(ctx context.Context, inbound_wire_drawdown_request_id string, opts ...option.RequestOption) (res *InboundWireDrawdownRequest, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("inbound_wire_drawdown_requests/%s", inbound_wire_drawdown_request_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Inbound Wire Drawdown Requests
func (r *InboundWireDrawdownRequestService) List(ctx context.Context, query InboundWireDrawdownRequestListParams, opts ...option.RequestOption) (res *shared.Page[InboundWireDrawdownRequest], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "inbound_wire_drawdown_requests"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// List Inbound Wire Drawdown Requests
func (r *InboundWireDrawdownRequestService) ListAutoPaging(ctx context.Context, query InboundWireDrawdownRequestListParams, opts ...option.RequestOption) *shared.PageAutoPager[InboundWireDrawdownRequest] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Inbound wire drawdown requests are requests from someone else to send them a
// wire. This feature is in beta; reach out to
// [support@increase.com](mailto:support@increase.com) to learn more.
type InboundWireDrawdownRequest struct {
	// A constant representing the object's type. For this resource it will always be
	// `inbound_wire_drawdown_request`.
	Type InboundWireDrawdownRequestType `json:"type,required"`
	// The Wire drawdown request identifier.
	ID string `json:"id,required"`
	// The Account Number from which the recipient of this request is being requested
	// to send funds.
	RecipientAccountNumberID string `json:"recipient_account_number_id,required"`
	// The drawdown request's originator's account number.
	OriginatorAccountNumber string `json:"originator_account_number,required"`
	// The drawdown request's originator's routing number.
	OriginatorRoutingNumber string `json:"originator_routing_number,required"`
	// The drawdown request's beneficiary's account number.
	BeneficiaryAccountNumber string `json:"beneficiary_account_number,required"`
	// The drawdown request's beneficiary's routing number.
	BeneficiaryRoutingNumber string `json:"beneficiary_routing_number,required"`
	// The amount being requested in cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
	// requested. Will always be "USD".
	Currency string `json:"currency,required"`
	// A message from the drawdown request's originator.
	MessageToRecipient string `json:"message_to_recipient,required,nullable"`
	// Line 1 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine1 string `json:"originator_to_beneficiary_information_line1,required,nullable"`
	// Line 2 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine2 string `json:"originator_to_beneficiary_information_line2,required,nullable"`
	// Line 3 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine3 string `json:"originator_to_beneficiary_information_line3,required,nullable"`
	// Line 4 of the information conveyed from the originator of the message to the
	// beneficiary.
	OriginatorToBeneficiaryInformationLine4 string `json:"originator_to_beneficiary_information_line4,required,nullable"`
	// The drawdown request's originator's name.
	OriginatorName string `json:"originator_name,required,nullable"`
	// Line 1 of the drawdown request's originator's address.
	OriginatorAddressLine1 string `json:"originator_address_line1,required,nullable"`
	// Line 2 of the drawdown request's originator's address.
	OriginatorAddressLine2 string `json:"originator_address_line2,required,nullable"`
	// Line 3 of the drawdown request's originator's address.
	OriginatorAddressLine3 string `json:"originator_address_line3,required,nullable"`
	// The drawdown request's beneficiary's name.
	BeneficiaryName string `json:"beneficiary_name,required,nullable"`
	// Line 1 of the drawdown request's beneficiary's address.
	BeneficiaryAddressLine1 string `json:"beneficiary_address_line1,required,nullable"`
	// Line 2 of the drawdown request's beneficiary's address.
	BeneficiaryAddressLine2 string `json:"beneficiary_address_line2,required,nullable"`
	// Line 3 of the drawdown request's beneficiary's address.
	BeneficiaryAddressLine3 string `json:"beneficiary_address_line3,required,nullable"`
	JSON                    inboundWireDrawdownRequestJSON
}

// inboundWireDrawdownRequestJSON contains the JSON metadata for the struct
// [InboundWireDrawdownRequest]
type inboundWireDrawdownRequestJSON struct {
	Type                                    apijson.Field
	ID                                      apijson.Field
	RecipientAccountNumberID                apijson.Field
	OriginatorAccountNumber                 apijson.Field
	OriginatorRoutingNumber                 apijson.Field
	BeneficiaryAccountNumber                apijson.Field
	BeneficiaryRoutingNumber                apijson.Field
	Amount                                  apijson.Field
	Currency                                apijson.Field
	MessageToRecipient                      apijson.Field
	OriginatorToBeneficiaryInformationLine1 apijson.Field
	OriginatorToBeneficiaryInformationLine2 apijson.Field
	OriginatorToBeneficiaryInformationLine3 apijson.Field
	OriginatorToBeneficiaryInformationLine4 apijson.Field
	OriginatorName                          apijson.Field
	OriginatorAddressLine1                  apijson.Field
	OriginatorAddressLine2                  apijson.Field
	OriginatorAddressLine3                  apijson.Field
	BeneficiaryName                         apijson.Field
	BeneficiaryAddressLine1                 apijson.Field
	BeneficiaryAddressLine2                 apijson.Field
	BeneficiaryAddressLine3                 apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *InboundWireDrawdownRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InboundWireDrawdownRequestType string

const (
	InboundWireDrawdownRequestTypeInboundWireDrawdownRequest InboundWireDrawdownRequestType = "inbound_wire_drawdown_request"
)

type InboundWireDrawdownRequestListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [InboundWireDrawdownRequestListParams]'s query parameters as
// `url.Values`.
func (r InboundWireDrawdownRequestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// A list of Inbound Wire Drawdown Request objects
type InboundWireDrawdownRequestListResponse struct {
	// The contents of the list.
	Data []InboundWireDrawdownRequest `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       inboundWireDrawdownRequestListResponseJSON
}

// inboundWireDrawdownRequestListResponseJSON contains the JSON metadata for the
// struct [InboundWireDrawdownRequestListResponse]
type inboundWireDrawdownRequestListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundWireDrawdownRequestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
