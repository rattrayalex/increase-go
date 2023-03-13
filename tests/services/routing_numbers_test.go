package services

import (
	"context"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

func TestRoutingNumbersListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.RoutingNumbers.List(context.TODO(), &types.RoutingNumberListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0)), RoutingNumber: increase.P("xxxxxxxxx")})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
