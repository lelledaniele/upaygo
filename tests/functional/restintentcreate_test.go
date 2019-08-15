package restintentcreate_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	appcurrency "github.com/lelledaniele/upaygo/currency"
	appcustomer "github.com/lelledaniele/upaygo/customer"

	restintentcreate "github.com/lelledaniele/upaygo/controller/rest/intent/create"
)

const (
	errorRestCreateIntent = "create intent controller failed: %v"
)

type responseIntent struct {
	IntentGatewayReference string           `json:"gateway_reference"`
	Customer               responseCustomer `json:"customer"`
}

type responseCustomer struct {
	R string `json:"gateway_reference"`
}

// Test a create intent request
func Test(t *testing.T) {
	var resI responseIntent

	c, _ := appcurrency.New("EUR")
	cus, e := appcustomer.NewStripe("email@email.com", c)
	if e != nil {
		t.Errorf(errorRestCreateIntent, e)
	}

	a, ps, w := 9999, "pm_card_visa", httptest.NewRecorder()
	p := fmt.Sprintf("currency=%v&amount=%v&payment_source=%v&customer_reference=%v", c.GetISO4217(), a, ps, cus.GetGatewayReference())
	req := httptest.NewRequest("POST", "http://example.com", strings.NewReader(p))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	restintentcreate.Handler(w, req)

	res := w.Result()
	resBody, e := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if e != nil {
		t.Errorf(errorRestCreateIntent, e)
	}

	e = json.Unmarshal(resBody, &resI)
	if e != nil {
		t.Errorf(errorRestCreateIntent, e)
	}

	if resI.IntentGatewayReference == "" {
		t.Errorf(errorRestCreateIntent, "the body response does not have the gateway reference")
	}

	if resI.Customer.R == "" {
		t.Errorf(errorRestCreateIntent, "the body response does not have the customer reference")
	}
}

// Test a create intent request without customer
func TestWithoutCustomer(t *testing.T) {
	var resI responseIntent

	c, a, ps, w := "EUR", 9999, "pm_card_visa", httptest.NewRecorder()
	p := fmt.Sprintf("currency=%v&amount=%v&payment_source=%v", c, a, ps)
	req := httptest.NewRequest("POST", "http://example.com", strings.NewReader(p))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	restintentcreate.Handler(w, req)

	res := w.Result()
	resBody, e := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if e != nil {
		t.Errorf(errorRestCreateIntent, e)
	}

	e = json.Unmarshal(resBody, &resI)
	if e != nil {
		t.Errorf(errorRestCreateIntent, e)
	}

	if resI.IntentGatewayReference == "" {
		t.Errorf(errorRestCreateIntent, "the body response does not have the gateway reference")
	}
}
