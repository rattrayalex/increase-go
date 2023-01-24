package services

import (
	"context"
	"testing"

	client "increase"
)

func TestCheckTransfersMail(t *testing.T) {
	t.Skip("Prism incorrectly returns an invalid JSON error")
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.CheckTransfers.Simulations.Mail(context.TODO(), "check_transfer_30b43acfu9vw8fyc4f5")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
