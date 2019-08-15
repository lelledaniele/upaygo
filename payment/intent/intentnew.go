package paymentintent

import (
	"errors"
	"time"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"

	appamount "github.com/lelledaniele/upaygo/amount"
	appconfig "github.com/lelledaniele/upaygo/config"
	appcustomer "github.com/lelledaniele/upaygo/customer"
	apppaymentsource "github.com/lelledaniele/upaygo/payment/source"
)

// New creates an intent in Stripe and returns it as an instance of i
func New(a appamount.Amount, p apppaymentsource.Source, c appcustomer.Customer) (Intent, error) {
	if a == nil || p == nil {
		return nil, errors.New("impossible to create a payment intent without required parameters")
	}

	sck, e := appconfig.GetStripeAPIConfigByCurrency(a.GetCurrency().GetISO4217())
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
