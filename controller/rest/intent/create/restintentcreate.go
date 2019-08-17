package apprestintentcreate

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	appamount "github.com/lelledaniele/upaygo/amount"
	apprest "github.com/lelledaniele/upaygo/controller/rest"
	appcustomer "github.com/lelledaniele/upaygo/customer"
	appintent "github.com/lelledaniele/upaygo/payment/intent"
	appsource "github.com/lelledaniele/upaygo/payment/source"
)

const (
	URL    = "/payment_intents"
	method = http.MethodPost

	responseTye = "application/json"

	errorMethod          = "'%v' is the only method supported"
	errorParsingParam    = "error during the payload parsing: '%v'"
	errorParamMissing    = "missing payload mandatory parameters to create a payment intent"
	errorParamAmountType = "error during the amount conversion: '%v'"
	errorAmountCreation  = "error during the intent amount creation: '%v'"
	errorIntentCreation  = "error during the intent creation: '%v'"
	errorIntentEncoding  = "error during the intent encoding: '%v'"
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
// @Success 200 {interface} paymentintent.Intent
// @Failure 405 {object} apprest.RESTError
// @Failure 500 {object} apprest.RESTError
// @Router /payment_intents [post]
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", responseTye)

	if r.Method != method {
		e := apprest.RESTError{
			M: fmt.Sprintf(errorMethod, method),
		}
		_ = json.NewEncoder(w).Encode(e)

		return
	}

	a, ps, cus, e := getParams(r)
	if e != nil {
		e := apprest.RESTError{
			M: e.Error(),
		}
		_ = json.NewEncoder(w).Encode(e)

		return
	}

	pi, e := appintent.New(a, ps, cus)
	if e != nil {
		e := apprest.RESTError{
			M: fmt.Sprintf(errorIntentCreation, e),
		}
		_ = json.NewEncoder(w).Encode(e)

		return
	}

	e = json.NewEncoder(w).Encode(pi)
	if e != nil {
		e := apprest.RESTError{
			M: fmt.Sprintf(errorIntentEncoding, e),
		}
		_ = json.NewEncoder(w).Encode(e)

		return
	}
}

// Get and transform the payload params into domain structs
func getParams(r *http.Request) (appamount.Amount, appsource.Source, appcustomer.Customer, error) {
	e := r.ParseForm()
	if e != nil {
		return nil, nil, nil, fmt.Errorf(errorParsingParam, e.Error())
	}

	p := r.Form
	if p.Get("currency") == "" || p.Get("amount") == "" || p.Get("payment_source") == "" {
		return nil, nil, nil, errors.New(errorParamMissing)
	}

	ai, e := strconv.Atoi(p.Get("amount"))
	if e != nil {
		return nil, nil, nil, fmt.Errorf(errorParamAmountType, e.Error())
	}

	a, e := appamount.New(ai, p.Get("currency"))
	if e != nil {
		return nil, nil, nil, fmt.Errorf(errorAmountCreation, e.Error())
	}

	var cus appcustomer.Customer
	if p.Get("customer_reference") != "" {
		cus = appcustomer.New(p.Get("customer_reference"), "")
	}

	ps := appsource.New(p.Get("payment_source"))

	return a, ps, cus, nil
}
