package paymentintent

const (
	requirescapture = "requires_capture"
	canceled        = "canceled"
	succeeded       = "succeeded"
)

type status struct {
	S string `json:"gateway_reference"`
}
