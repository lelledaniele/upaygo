package main

type PaymentIntent struct {
	R    string `json:"gateway_reference"`      // Gateway reference
	UR   string `json:"user_gateway_reference"` // User gateway reference
	CM   string `json:"confirmation_method"`    // Confirmation method
	CT   int    `json:"created_at"`             // Create At
	NA   string `json:"next_action"`            // Next Action
	OFFS bool   `json:"off_session"`            // Off session

	C Customer      `json:"customer"`
	P PaymentSource `json:"payment_source"`
	A Amount        `json:"amount"`
	S IntentStatus  `json:"intent_status"`
}
