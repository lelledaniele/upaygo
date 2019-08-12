package controllerintentcreate

import (
	"fmt"
	"net/http"
)

const (
	URL = "/payment_intents"

	method      = http.MethodPost
	methodError = "'%v' is the only method supported"
)

// @Summary Create an intent
// @Description Create an unconfirmed and manual intent
// @Tags intent
// @Accept json
// @Produce json
// @Param currency formData string true "Intent's currency"
// @Param amount formData int true "Intent's amount"
// @Param payment_source formData string true "Intent's payment source"
// @Param customer_reference formData string false "Intent's customer reference"
// @Success 200 {object} intent.PaymentIntent
// @Failure 400 {object} controllererror.RestApiError
// @Failure 404 {object} controllererror.RestApiError
// @Failure 500 {object} controllererror.RestApiError
// @Router /payment_intents [post]
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != method {
		_, _ = w.Write([]byte(fmt.Sprintf(methodError, method)))

		return
	}
}
