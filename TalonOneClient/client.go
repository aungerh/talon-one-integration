package TalonOneClient

// Client data type
type Client struct {
	AppID   int
	AppKey  string
	Payload Payload
}

// UpdateCustomerProfile updates customer profile
func (c *Client) UpdateCustomerProfile() {
	BuildAndRequest("PUT", &c.Payload, customerProfileEndpoint+c.Payload.URLParams, c)
}

// UpdateCustomerSession updates a session
func (c *Client) UpdateCustomerSession() {
	BuildAndRequest("PUT", &c.Payload, customerSessionsEndpoint+c.Payload.URLParams, c)
}

// SendEvents report events
func (c *Client) SendEvents() {
	BuildAndRequest("POST", &c.Payload, eventsEndpoint, c)
}
