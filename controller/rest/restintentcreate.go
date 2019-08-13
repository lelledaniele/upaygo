package apprest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/lelledaniele/upaygo/customer"
	appintent "github.com/lelledaniele/upaygo/payment/intent"

	appsource "github.com/lelledaniele/upaygo/payment/source"

	"github.com/lelledaniele/upaygo/amount"
)

const (
	IntentCreateURL = "/payment_intents"

	intentCreateMethod      = http.MethodPost
	intentCreateMethodError = "'%v' is the only method supported"

	intentCreateParamMissingError = "missing payload mandatory parameters to create a payment intent"
)

// @Summary Create an intent
// @Description Create an unconfirmed and manual intent
// @Tags Intent
// @Accept x-www-form-urlencoded
// @Produce json
// @Param currency formData string true "Intent's currency"
// @Param amount formData int true "Intent's amount"
// @Param payment_source formData string true "Intent's payment source"
// @Param customer_reference formData string false "Intent's customer reference"
// @Success 200 {object} paymentintent.Intent
// @Failure 405 {object} apprest.RESTError
// @Failure 500 {object} apprest.RESTError
// @Router /payment_intents [post]
func IntentCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != intentCreateMethod {
		handleError(&w, fmt.Sprintf(intentCreateMethodError, intentCreateMethod))
		return
	}

	e := r.ParseForm()
	if e != nil {
		handleError(&w, e.Error())
		return
	}

	p := r.Form

	if p.Get("currency") == "" || p.Get("amount") == "" || p.Get("payment_source") == "" {
		e := RESTError{M: intentCreateParamMissingError}
		_ = json.NewEncoder(w).Encode(e)
		return
	}

	ai, e := strconv.Atoi(p.Get("amount"))
	if e != nil {
		handleError(&w, e.Error())
		return
	}

	a, e := amount.New(ai, p.Get("currency"))
	if e != nil {
		handleError(&w, e.Error())
		return
	}

	ps := appsource.New(p.Get("payment_source"))

	var cus customer.Customer
	if p.Get("customer_reference") != "" {
		cus = customer.New(p.Get("customer_reference"))
	}

	pi, e := appintent.New(a, ps, cus)
	if e != nil {
		handleError(&w, e.Error())
		return
	}

	e = json.NewEncoder(w).Encode(pi)
	if e != nil {
		handleError(&w, e.Error())
		return
	}
}
