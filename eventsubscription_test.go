package increase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
)

func TestEventSubscriptionNewWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.EventSubscriptions.New(context.TODO(), increase.EventSubscriptionNewParams{URL: increase.F("https://website.com/webhooks"), SharedSecret: increase.F("x"), SelectedEventCategory: increase.F(increase.EventSubscriptionNewParamsSelectedEventCategoryAccountCreated)})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventSubscriptionGet(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.EventSubscriptions.Get(
		context.TODO(),
		"event_subscription_001dzz0r20rcdxgb013zqb8m04g",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventSubscriptionUpdateWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.EventSubscriptions.Update(
		context.TODO(),
		"event_subscription_001dzz0r20rcdxgb013zqb8m04g",
		increase.EventSubscriptionUpdateParams{Status: increase.F(increase.EventSubscriptionUpdateParamsStatusActive)},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventSubscriptionListWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.EventSubscriptions.List(context.TODO(), increase.EventSubscriptionListParams{Cursor: increase.F("string"), Limit: increase.F(int64(0))})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}