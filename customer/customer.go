package appcustomer

type Customer interface {
	GetGatewayReference() string
	GetEmail() string
}

type c struct {
	R     string `json:"gateway_reference"` // Gateway reference
	Email string `json:"email"`
}

// GetGatewayReference exposes c.R value
func (c *c) GetGatewayReference() string {
	return c.R
}

// GetEmail exposes c.Email value
func (c *c) GetEmail() string {
	return c.Email
}

// New returns a new instance of c
func New(r string, email string) Customer {
	return &c{R: r, Email: email}
}
