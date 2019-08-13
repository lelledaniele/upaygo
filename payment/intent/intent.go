package paymentintent

import (
	"errors"
	"time"

	paymentsource "github.com/lelledaniele/upaygo/payment/source"

	"github.com/lelledaniele/upaygo/config"

	"github.com/lelledaniele/upaygo/amount"
	"github.com/lelledaniele/upaygo/customer"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

type Intent interface {
	GetGatewayReference() string
	GetConfirmationMethod() string
	GetNextAction() string
	IsOffSession() bool
	GetCreatedAt() time.Time
	GetCustomer() customer.Customer
	GetSource() paymentsource.Source
	GetAmount() amount.Amount

	// From the status
	IsCanceled() bool
	IsSucceeded() bool
	RequiresCapture() bool
}

type i struct {
	R    string    `json:"gateway_reference"`   // Gateway reference
	CM   string    `json:"confirmation_method"` // Confirmation method
	NA   string    `json:"next_action"`         // Next Action
	OFFS bool      `json:"off_session"`         // Off session
	CT   time.Time `json:"created_at"`          // Create At

	C  customer.Customer    `json:"customer"`
	PS paymentsource.Source `json:"payment_source"`
	A  amount.Amount        `json:"amount"`
	S  status               `json:"status"`
}

// GetGatewayReference exposes i.R value
func (i *i) GetGatewayReference() string {
	return i.R
}

// GetGatewayReference exposes i.CM value
func (i *i) GetConfirmationMethod() string {
	return i.CM
}

// GetGatewayReference exposes i.NA value
func (i *i) GetNextAction() string {
	return i.NA
}

// GetGatewayReference exposes i.OFFS value
func (i *i) IsOffSession() bool {
	return i.OFFS
}

// GetGatewayReference exposes i.CT. value
func (i *i) GetCreatedAt() time.Time {
	return i.CT
}

// GetGatewayReference exposes i.C value
func (i *i) GetCustomer() customer.Customer {
	return i.C
}

// GetGatewayReference exposes i.PS value
func (i *i) GetSource() paymentsource.Source {
	return i.PS
}

// GetGatewayReference exposes i.A value
func (i *i) GetAmount() amount.Amount {
	return i.A
}

// IsCanceled if status.s is canceled
func (i *i) IsCanceled() bool {
	return i.S.S == canceled
}

// IsSucceeded if status.s is succeeded
func (i *i) IsSucceeded() bool {
	return i.S.S == succeeded
}

// RequiresCapture if status.s is requirescapture
func (i *i) RequiresCapture() bool {
	return i.S.S == requirescapture
}

// New creates an intent in Stripe and returns it as an instance of i
func New(a amount.Amount, p paymentsource.Source, c customer.Customer) (Intent, error) {
	var e error

	if a == nil || p == nil {
		return nil, errors.New("impossible to create a payment intent without amount and payment source")
	}

	sck, e := config.GetStripeAPIConfigByCurrency(a.GetCurrency().GetISO4217())
	if e != nil {
		return nil, e
	}

	stripe.Key = sck.GetSK()

	ic := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(int64(a.GetAmount())),
		Currency:           stripe.String(a.GetCurrency().GetISO4217()),
		PaymentMethod:      stripe.String(p.GetGatewayReference()),
		SetupFutureUsage:   stripe.String("off_session"),
		ConfirmationMethod: stripe.String("manual"),
		CaptureMethod:      stripe.String("manual"),
	}

	if c != nil {
		ic.Customer = stripe.String(c.GetGatewayReference())
		ic.SavePaymentMethod = stripe.Bool(true)
	}

	spi, e := paymentintent.New(ic)
	if e != nil {
		return nil, e
	}

	var nat string
	na := spi.NextAction
	if na != nil {
		nat = string(na.Type)
	}

	return &i{
		R:    spi.ID,
		CM:   string(spi.ConfirmationMethod),
		NA:   nat,
		OFFS: spi.SetupFutureUsage == "off_session",
		CT:   time.Unix(spi.Created, 0),
		C:    c,
		PS:   p,
		A:    a,
		S:    status{S: string(spi.Status)},
	}, nil
}
