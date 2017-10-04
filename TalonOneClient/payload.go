package TalonOneClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Payload represents the body of the request
type Payload struct {
	ProfileID  string     `json:"profileId,omitempty"`
	Coupon     string     `json:"coupon,omitempty"`
	Referral   string     `json:"referral,omitempty"`
	State      string     `json:"state,omitempty"`
	CartItems  string     `json:"cartItems,omitempty"`
	Total      string     `json:"total,omitempty"`
	Attributes attributes `json:"attributes"`
}

type attributes struct {
	Name               string `json:"Name,omitempty"`
	Gender             string `json:"Gender,omitempty"`
	BirthDate          string `json:"BirthDate,omitempty"`
	Email              string `json:"Email,omitempty"`
	AdditionalEmails   string `json:"Additionalemails,omitempty"`
	Phone              string `json:"Phone,omitempty"`
	Fax                string `json:"Fax,omitempty"`
	URL                string `json:"URL,omitempty"`
	Language           string `json:"Language,omitempty"`
	Locale             string `json:"Locale,omitempty"`
	SignupDate         string `json:"SignupDate,omitempty"`
	BillingSalutation  string `json:"BillingSalutation,omitempty"`
	BillingName        string `json:"BillingName,omitempty"`
	BillingAddress1    string `json:"BillingAddress1,omitempty"`
	BillingAddress2    string `json:"BillingAddress2,omitempty"`
	BillingAddress3    string `json:"BillingAddress3,omitempty"`
	BillingAddress4    string `json:"BillingAddress4,omitempty"`
	BillingCity        string `json:"BillingCity,omitempty"`
	BillingPostalCode  string `json:"BillingPostalCode,omitempty"`
	Billing            string `json:"Billing,omitempty"`
	ShippingSalutation string `json:"ShippingSalutation,omitempty"`
	ShippingName       string `json:"ShippingName,omitempty"`
	ShippingAddress1   string `json:"ShippingAddress1,omitempty"`
	ShippingAddress2   string `json:"ShippingAddress2,omitempty"`
	ShippingAddress3   string `json:"ShippingAddress3,omitempty"`
	ShippingAddress4   string `json:"ShippingAddress4,omitempty"`
	ShippingCity       string `json:"ShippingCity,omitempty"`
	ShippingPostalCode string `json:"ShippingPostalCode,omitempty"`
	ShippingCountry    string `json:"ShippingCountry,omitempty"`
	PaymentMethod      string `json:"PaymentMethod,omitempty"`
	PaymentHash        string `json:"PaymentHash,omitempty"`
}

// UpdateCustomerProfile updates customer profile
// TODO abstract common blocks
func (p *Payload) UpdateCustomerProfile() {
	dest := "https://demo.talon.one/v1/customer_profiles/"
	profileName := "tronald_dump"
	url := dest + profileName

	p.Attributes.Gender = "m"

	js, _ := json.Marshal(*p)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(js))
	signatureVal := fmt.Sprintf("signer=%d; signature=%s", 60, signPayload(js))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Signature", signatureVal)

	doRequest(req)
}

// UpdateCustomerSession updates a session
// TODO abstract common blocks
func (p *Payload) UpdateCustomerSession() {
	dest := "https://demo.talon.one/v1/customer_sessions/"
	sessionName := "SessionTest11"
	url := dest + sessionName

	p.Coupon = "LJI523KI"

	js, _ := json.Marshal(*p)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(js))
	signatureVal := fmt.Sprintf("signer=%d; signature=%s", 60, signPayload(js))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Signature", signatureVal)

	doRequest(req)
}
