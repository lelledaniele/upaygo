package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	intentCreateURL = "/payment_intents"

	intentCreateMethod      = http.MethodPost
	intentCreateMethodError = "'%v' is the only method supported"
)

// @Summary Create an intent
// @Description Create an unconfirmed and manual intent
// @Tags Intent
// @Accept mpfd
// @Produce json
// @Param currency formData string true "Intent's currency"
// @Param amount formData int true "Intent's amount"
// @Param payment_source formData string true "Intent's payment source"
// @Param customer_reference formData string false "Intent's customer reference"
// @Success 200 {object} main.PaymentIntent
// @Failure 405 {object} main.RESTError
// @Failure 500 {object} main.RESTError
// @Router /payment_intents [post]
func intentCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != intentCreateMethod {
		e := RESTError{M: fmt.Sprintf(intentCreateMethodError, intentCreateMethod)}
		_ = json.NewEncoder(w).Encode(e)
		return
	}

	_, _ = w.Write([]byte("Ciao"))
}
