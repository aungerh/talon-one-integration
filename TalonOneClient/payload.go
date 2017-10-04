package TalonOneClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Payload represents the body of the request
type Payload struct {
	Attributes attributes `json:"attributes"`
}

type attributes struct {
	Gender string `json:"Gender,omitempty"`
}

// UpdateCustomerProfile updates customer profile
func (p *Payload) UpdateCustomerProfile() {
	dest := "https://demo.talon.one/v1/customer_profiles/"
	profileName := "tronald_dump"
	url := dest + profileName
	js, _ := json.Marshal(*p)

	p.Attributes.Gender = "m"

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(js))
	signatureVal := fmt.Sprintf("signer=%d; signature=%s", 60, signPayload(js))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Signature", signatureVal)

	doRequest(req)
}
