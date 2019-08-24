package apprestintentconfirm

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	appcurrency "github.com/lelledaniele/upaygo/currency"
	apperror "github.com/lelledaniele/upaygo/error"
	apppaymentintentconfirm "github.com/lelledaniele/upaygo/payment/intent/confirm"

	"github.com/gorilla/mux"
)

const (
	URL    = "/payment_intents/{id}/confirm"
	method = http.MethodPost

	responseTye = "application/json"

	errorMethod              = "'%v' is the only method supported"
	errorParamPathMissing    = "missing URL in-path mandatory parameters to confirm a payment intent"
	errorParsingParam        = "error during the payload parsing: '%v'"
	errorParamPayloadMissing = "missing payload mandatory parameters to confirm a payment intent"
	errorAmountCreation      = "error during the intent amount creation: '%v'"
	errorIntentConfirmation  = "error during the intent confirmation: '%v'"
	errorIntentEncoding      = "error during the intent encoding: '%v'"
)

// @Summary Confirm an intent
// @Description Confirm an unconfirmed intent
// @Tags Intent
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id path string true "Intent's ID"
// @Param currency formData string true "Intent's currency"
// @Success 200 {interface} apppaymentintent.Intent
// @Failure 405 {object} apperror.RESTError
// @Failure 500 {object} apperror.RESTError
// @Router /payment_intents/{id}/confirm [post]
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", responseTye)

	if r.Method != method {
		e := apperror.RESTError{
			M: fmt.Sprintf(errorMethod, method),
		}
		_ = json.NewEncoder(w).Encode(e)

		return
	}

	ID, cur, e := getParams(r)
	if e != nil {
		e := apperror.RESTError{
			M: e.Error(),
		}
		_ = json.NewEncoder(w).Encode(e)

		return
	}

	appintent, e := apppaymentintentconfirm.Confirm(ID, cur)
	if e != nil {
		e := apperror.RESTError{
			M: fmt.Sprintf(errorIntentConfirmation, e),
		}
		_ = json.NewEncoder(w).Encode(e)

		return
	}

	e = json.NewEncoder(w).Encode(appintent)
	if e != nil {
		e := apperror.RESTError{
			M: fmt.Sprintf(errorIntentEncoding, e),
		}
		_ = json.NewEncoder(w).Encode(e)

		return
	}
}

// Get and transform the payload params into domain structs
func getParams(r *http.Request) (string, appcurrency.Currency, error) {
	vars := mux.Vars(r)

	ID, ok := vars["id"]
	if !ok || ID == "" {
		return "", nil, errors.New(errorParamPathMissing)
	}

	e := r.ParseForm()
	if e != nil {
		return "", nil, fmt.Errorf(errorParsingParam, e.Error())
	}

	p := r.Form
	if p.Get("currency") == "" {
		return "", nil, errors.New(errorParamPayloadMissing)
	}

	cur, e := appcurrency.New(p.Get("currency"))
	if e != nil {
		return "", nil, fmt.Errorf(errorAmountCreation, e.Error())
	}

	return ID, cur, nil
}
