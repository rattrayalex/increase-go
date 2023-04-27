package services

import (
	"context"
	"errors"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func TestRealTimeDecisionGet(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.RealTimeDecisions.Get(
		context.TODO(),
		"real_time_decision_j76n2e810ezcg3zh5qtn",
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRealTimeDecisionActionWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.RealTimeDecisions.Action(
		context.TODO(),
		"real_time_decision_j76n2e810ezcg3zh5qtn",
		requests.RealTimeDecisionActionParams{CardAuthorization: increase.F(requests.RealTimeDecisionActionParamsCardAuthorization{Decision: increase.F(requests.RealTimeDecisionActionParamsCardAuthorizationDecisionApprove)}), DigitalWalletToken: increase.F(requests.RealTimeDecisionActionParamsDigitalWalletToken{Approval: increase.F(requests.RealTimeDecisionActionParamsDigitalWalletTokenApproval{CardProfileID: increase.F("string"), Phone: increase.F("x"), Email: increase.F("x")}), Decline: increase.F(requests.RealTimeDecisionActionParamsDigitalWalletTokenDecline{Reason: increase.F("x")})}), DigitalWalletAuthentication: increase.F(requests.RealTimeDecisionActionParamsDigitalWalletAuthentication{Result: increase.F(requests.RealTimeDecisionActionParamsDigitalWalletAuthenticationResultSuccess)})},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
