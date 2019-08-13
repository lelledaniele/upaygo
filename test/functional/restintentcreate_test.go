package functionalcreateintent_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lelledaniele/upaygo/config"
	apprest "github.com/lelledaniele/upaygo/controller/rest"
)

type responseIntent struct {
	IntentGatewayReference string `json:"gateway_reference"`
}

// Test a create intent request
//func Test(t *testing.T) {
//
//}

// Test a create intent request without customer
func TestWithoutCustomer(t *testing.T) {
	var resI responseIntent
	c, a, ps := "EUR", 5040, "pm_card_visa"

	var e error
	es := "create intent controller failed: %v"

	ts := httptest.NewServer(http.HandlerFunc(apprest.IntentCreateHandler))
	defer ts.Close()

	body := strings.NewReader(fmt.Sprintf("currency=%v&amount=%v&payment_source=%v", c, a, ps))
	res, e := http.Post(config.GetServerConfig().GetURI()+apprest.IntentCreateURL, "application/x-www-form-urlencoded", body)
	if e != nil {
		t.Errorf(es, e)
	}

	resBody, e := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if e != nil {
		t.Errorf(es, e)
	}

	e = json.Unmarshal(resBody, &resI)
	if e != nil {
		t.Errorf(es, e)
	}

	if resI.IntentGatewayReference == "" {
		t.Errorf(es, "the body response does not have the gateway reference")
	}
}
