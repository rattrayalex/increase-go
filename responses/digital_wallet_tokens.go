package responses

import (
	"time"

	apijson "github.com/increase/increase-go/core/json"
)

type DigitalWalletToken struct {
	// The Digital Wallet Token identifier.
	ID string `json:"id,required"`
	// The identifier for the Card this Digital Wallet Token belongs to.
	CardID string `json:"card_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// This indicates if payments can be made with the Digital Wallet Token.
	Status DigitalWalletTokenStatus `json:"status,required"`
	// The digital wallet app being used.
	TokenRequestor DigitalWalletTokenTokenRequestor `json:"token_requestor,required"`
	// A constant representing the object's type. For this resource it will always be
	// `digital_wallet_token`.
	Type DigitalWalletTokenType `json:"type,required"`
	JSON DigitalWalletTokenJSON
}

type DigitalWalletTokenJSON struct {
	ID             apijson.Metadata
	CardID         apijson.Metadata
	CreatedAt      apijson.Metadata
	Status         apijson.Metadata
	TokenRequestor apijson.Metadata
	Type           apijson.Metadata
	Raw            []byte
	Extras         map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into DigitalWalletToken using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *DigitalWalletToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DigitalWalletTokenStatus string

const (
	DigitalWalletTokenStatusActive      DigitalWalletTokenStatus = "active"
	DigitalWalletTokenStatusInactive    DigitalWalletTokenStatus = "inactive"
	DigitalWalletTokenStatusSuspended   DigitalWalletTokenStatus = "suspended"
	DigitalWalletTokenStatusDeactivated DigitalWalletTokenStatus = "deactivated"
)

type DigitalWalletTokenTokenRequestor string

const (
	DigitalWalletTokenTokenRequestorApplePay  DigitalWalletTokenTokenRequestor = "apple_pay"
	DigitalWalletTokenTokenRequestorGooglePay DigitalWalletTokenTokenRequestor = "google_pay"
)

type DigitalWalletTokenType string

const (
	DigitalWalletTokenTypeDigitalWalletToken DigitalWalletTokenType = "digital_wallet_token"
)

type DigitalWalletTokenListResponse struct {
	// The contents of the list.
	Data []DigitalWalletToken `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       DigitalWalletTokenListResponseJSON
}

type DigitalWalletTokenListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	Raw        []byte
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// DigitalWalletTokenListResponse using the internal json library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *DigitalWalletTokenListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
