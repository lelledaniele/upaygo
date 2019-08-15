package appcustomer

type Customer interface {
	GetGatewayReference() string
}

type c struct {
	R     string `json:"gateway_reference"` // Gateway reference
	Email string `json:"email"`
}

// GetGatewayReference exposes c.R value
func (c *c) GetGatewayReference() string {
	return c.R
}

// New returns a new instance of c
func New(r string, email string) Customer {
	return &c{R: r, Email: email}
}
