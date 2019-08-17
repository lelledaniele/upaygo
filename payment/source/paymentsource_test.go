// +build unit

package apppaymentsource_test

import (
	"testing"

	apppaymentsource "github.com/lelledaniele/upaygo/payment/source"
)

func TestNew(t *testing.T) {
	r := "card_XXX"
	got := apppaymentsource.New(r)

	if got.GetGatewayReference() != r {
		t.Errorf("New payment source reference is incorrect, got %v want %v", got.GetGatewayReference(), r)
	}
}
