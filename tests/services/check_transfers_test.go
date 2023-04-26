package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func TestCheckTransferNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CheckTransfers.New(context.TODO(), &requests.CheckTransferNewParams{AccountID: increase.F("account_in71c4amph0vgo2qllky"), AddressLine1: increase.F("33 Liberty Street"), AddressLine2: increase.F("x"), AddressCity: increase.F("New York"), AddressState: increase.F("NY"), AddressZip: increase.F("10045"), ReturnAddress: increase.F(requests.CheckTransferNewParamsReturnAddress{Name: increase.F("x"), Line1: increase.F("x"), Line2: increase.F("x"), City: increase.F("x"), State: increase.F("x"), Zip: increase.F("x")}), Amount: increase.F(int64(1000)), Message: increase.F("Check payment"), Note: increase.F("x"), RecipientName: increase.F("Ian Crease"), RequireApproval: increase.F(true)})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckTransferGet(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CheckTransfers.Get(
		context.TODO(),
		"check_transfer_30b43acfu9vw8fyc4f5",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckTransferListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CheckTransfers.List(context.TODO(), &requests.CheckTransferListParams{Cursor: increase.F("string"), Limit: increase.F(int64(0)), AccountID: increase.F("string"), CreatedAt: increase.F(requests.CheckTransferListParamsCreatedAt{After: increase.F(time.Now()), Before: increase.F(time.Now()), OnOrAfter: increase.F(time.Now()), OnOrBefore: increase.F(time.Now())})})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckTransferApprove(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CheckTransfers.Approve(
		context.TODO(),
		"check_transfer_30b43acfu9vw8fyc4f5",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckTransferCancel(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CheckTransfers.Cancel(
		context.TODO(),
		"check_transfer_30b43acfu9vw8fyc4f5",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckTransferStopPayment(t *testing.T) {
	t.Skip("Prism doesn't accept no request body being sent but returns 415 if it is sent")
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.CheckTransfers.StopPayment(
		context.TODO(),
		"check_transfer_30b43acfu9vw8fyc4f5",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
