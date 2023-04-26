package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/core/field"
	apijson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
)

type AccountNewParams struct {
	// The identifier for the Entity that will own the Account.
	EntityID field.Field[string] `json:"entity_id"`
	// The identifier for the Program that this Account falls under.
	ProgramID field.Field[string] `json:"program_id"`
	// The identifier of an Entity that, while not owning the Account, is associated
	// with its activity. Its relationship to your group must be `informational`.
	InformationalEntityID field.Field[string] `json:"informational_entity_id"`
	// The name you choose for the Account.
	Name field.Field[string] `json:"name,required"`
}

// MarshalJSON serializes AccountNewParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r AccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountUpdateParams struct {
	// The new name of the Account.
	Name field.Field[string] `json:"name"`
}

// MarshalJSON serializes AccountUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r AccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// Filter Accounts for those belonging to the specified Entity.
	EntityID field.Field[string] `query:"entity_id"`
	// Filter Accounts for those with the specified status.
	Status    field.Field[AccountListParamsStatus]    `query:"status"`
	CreatedAt field.Field[AccountListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes AccountListParams into a url.Values of the query parameters
// associated with this value
func (r AccountListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type AccountListParamsStatus string

const (
	AccountListParamsStatusOpen   AccountListParamsStatus = "open"
	AccountListParamsStatusClosed AccountListParamsStatus = "closed"
)

type AccountListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After field.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before field.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter field.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore field.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes AccountListParamsCreatedAt into a url.Values of the query
// parameters associated with this value
func (r AccountListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
