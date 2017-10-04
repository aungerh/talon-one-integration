package TalonOneClient

// Payload represents the body of the request
type Payload struct {
	Attributes attributes `json:"attributes"`
}

type attributes struct {
	Gender string `json:"Gender,omitempty"`
}
