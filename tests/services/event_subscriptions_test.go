package services

import (
	"context"
	"increase"
	"testing"

	client "increase"
	"increase/types"
)

func TestEventSubscriptionsCreateWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.EventSubscriptions.Create(context.TODO(), &types.CreateAnEventSubscriptionParameters{URL: increase.P("https://website.com/webhooks")})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestEventSubscriptionsRetrieve(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.EventSubscriptions.Retrieve(context.TODO(), "event_subscription_001dzz0r20rcdxgb013zqb8m04g")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestEventSubscriptionsUpdateWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.EventSubscriptions.Update(
		context.TODO(),
		"event_subscription_001dzz0r20rcdxgb013zqb8m04g",
		&types.UpdateAnEventSubscriptionParameters{},
	)
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestEventSubscriptionsListWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.EventSubscriptions.List(context.TODO(), &types.ListEventSubscriptionsQuery{})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}