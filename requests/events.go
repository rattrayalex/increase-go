package requests

import (
	"fmt"
	"net/url"
	"time"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type EventListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter Events to those belonging to the object with the provided identifier.
	AssociatedObjectID fields.Field[string]                   `query:"associated_object_id"`
	CreatedAt          fields.Field[EventListParamsCreatedAt] `query:"created_at"`
	Category           fields.Field[EventListParamsCategory]  `query:"category"`
}

// URLQuery serializes EventListParams into a url.Values of the query parameters
// associated with this value
func (r *EventListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r EventListParams) String() (result string) {
	return fmt.Sprintf("&EventListParams{Cursor:%s Limit:%s AssociatedObjectID:%s CreatedAt:%s Category:%s}", r.Cursor, r.Limit, r.AssociatedObjectID, r.CreatedAt, r.Category)
}

type EventListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After fields.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before fields.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter fields.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore fields.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes EventListParamsCreatedAt into a url.Values of the query
// parameters associated with this value
func (r *EventListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r EventListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&EventListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}

type EventListParamsCategory struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In fields.Field[[]EventListParamsCategoryIn] `query:"in"`
}

// URLQuery serializes EventListParamsCategory into a url.Values of the query
// parameters associated with this value
func (r *EventListParamsCategory) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r EventListParamsCategory) String() (result string) {
	return fmt.Sprintf("&EventListParamsCategory{In:%s}", core.Fmt(r.In))
}

type EventListParamsCategoryIn string

const (
	EventListParamsCategoryInAccountCreated                                       EventListParamsCategoryIn = "account.created"
	EventListParamsCategoryInAccountUpdated                                       EventListParamsCategoryIn = "account.updated"
	EventListParamsCategoryInAccountNumberCreated                                 EventListParamsCategoryIn = "account_number.created"
	EventListParamsCategoryInAccountNumberUpdated                                 EventListParamsCategoryIn = "account_number.updated"
	EventListParamsCategoryInAccountStatementCreated                              EventListParamsCategoryIn = "account_statement.created"
	EventListParamsCategoryInAccountTransferCreated                               EventListParamsCategoryIn = "account_transfer.created"
	EventListParamsCategoryInAccountTransferUpdated                               EventListParamsCategoryIn = "account_transfer.updated"
	EventListParamsCategoryInACHPrenotificationCreated                            EventListParamsCategoryIn = "ach_prenotification.created"
	EventListParamsCategoryInACHPrenotificationUpdated                            EventListParamsCategoryIn = "ach_prenotification.updated"
	EventListParamsCategoryInACHTransferCreated                                   EventListParamsCategoryIn = "ach_transfer.created"
	EventListParamsCategoryInACHTransferUpdated                                   EventListParamsCategoryIn = "ach_transfer.updated"
	EventListParamsCategoryInCardCreated                                          EventListParamsCategoryIn = "card.created"
	EventListParamsCategoryInCardUpdated                                          EventListParamsCategoryIn = "card.updated"
	EventListParamsCategoryInCardDisputeCreated                                   EventListParamsCategoryIn = "card_dispute.created"
	EventListParamsCategoryInCardDisputeUpdated                                   EventListParamsCategoryIn = "card_dispute.updated"
	EventListParamsCategoryInCheckDepositCreated                                  EventListParamsCategoryIn = "check_deposit.created"
	EventListParamsCategoryInCheckDepositUpdated                                  EventListParamsCategoryIn = "check_deposit.updated"
	EventListParamsCategoryInCheckTransferCreated                                 EventListParamsCategoryIn = "check_transfer.created"
	EventListParamsCategoryInCheckTransferUpdated                                 EventListParamsCategoryIn = "check_transfer.updated"
	EventListParamsCategoryInDeclinedTransactionCreated                           EventListParamsCategoryIn = "declined_transaction.created"
	EventListParamsCategoryInDigitalWalletTokenCreated                            EventListParamsCategoryIn = "digital_wallet_token.created"
	EventListParamsCategoryInDigitalWalletTokenUpdated                            EventListParamsCategoryIn = "digital_wallet_token.updated"
	EventListParamsCategoryInDocumentCreated                                      EventListParamsCategoryIn = "document.created"
	EventListParamsCategoryInEntityCreated                                        EventListParamsCategoryIn = "entity.created"
	EventListParamsCategoryInEntityUpdated                                        EventListParamsCategoryIn = "entity.updated"
	EventListParamsCategoryInExternalAccountCreated                               EventListParamsCategoryIn = "external_account.created"
	EventListParamsCategoryInFileCreated                                          EventListParamsCategoryIn = "file.created"
	EventListParamsCategoryInGroupUpdated                                         EventListParamsCategoryIn = "group.updated"
	EventListParamsCategoryInGroupHeartbeat                                       EventListParamsCategoryIn = "group.heartbeat"
	EventListParamsCategoryInInboundACHTransferReturnCreated                      EventListParamsCategoryIn = "inbound_ach_transfer_return.created"
	EventListParamsCategoryInInboundACHTransferReturnUpdated                      EventListParamsCategoryIn = "inbound_ach_transfer_return.updated"
	EventListParamsCategoryInInboundWireDrawdownRequestCreated                    EventListParamsCategoryIn = "inbound_wire_drawdown_request.created"
	EventListParamsCategoryInOauthConnectionCreated                               EventListParamsCategoryIn = "oauth_connection.created"
	EventListParamsCategoryInOauthConnectionDeactivated                           EventListParamsCategoryIn = "oauth_connection.deactivated"
	EventListParamsCategoryInPendingTransactionCreated                            EventListParamsCategoryIn = "pending_transaction.created"
	EventListParamsCategoryInPendingTransactionUpdated                            EventListParamsCategoryIn = "pending_transaction.updated"
	EventListParamsCategoryInRealTimeDecisionCardAuthorizationRequested           EventListParamsCategoryIn = "real_time_decision.card_authorization_requested"
	EventListParamsCategoryInRealTimeDecisionDigitalWalletTokenRequested          EventListParamsCategoryIn = "real_time_decision.digital_wallet_token_requested"
	EventListParamsCategoryInRealTimeDecisionDigitalWalletAuthenticationRequested EventListParamsCategoryIn = "real_time_decision.digital_wallet_authentication_requested"
	EventListParamsCategoryInRealTimePaymentsTransferCreated                      EventListParamsCategoryIn = "real_time_payments_transfer.created"
	EventListParamsCategoryInRealTimePaymentsTransferUpdated                      EventListParamsCategoryIn = "real_time_payments_transfer.updated"
	EventListParamsCategoryInRealTimePaymentsRequestForPaymentCreated             EventListParamsCategoryIn = "real_time_payments_request_for_payment.created"
	EventListParamsCategoryInRealTimePaymentsRequestForPaymentUpdated             EventListParamsCategoryIn = "real_time_payments_request_for_payment.updated"
	EventListParamsCategoryInTransactionCreated                                   EventListParamsCategoryIn = "transaction.created"
	EventListParamsCategoryInWireDrawdownRequestCreated                           EventListParamsCategoryIn = "wire_drawdown_request.created"
	EventListParamsCategoryInWireDrawdownRequestUpdated                           EventListParamsCategoryIn = "wire_drawdown_request.updated"
	EventListParamsCategoryInWireTransferCreated                                  EventListParamsCategoryIn = "wire_transfer.created"
	EventListParamsCategoryInWireTransferUpdated                                  EventListParamsCategoryIn = "wire_transfer.updated"
)
