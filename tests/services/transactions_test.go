package services

import (
	"context"
	"testing"

	client "increase"
	"increase/types"
)

func TestTransactionsRetrieve(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Transactions.Retrieve(context.TODO(), "transaction_uyrp7fld2ium70oa7oi")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestTransactionsListWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.Transactions.List(context.TODO(), &types.ListTransactionsQuery{})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}