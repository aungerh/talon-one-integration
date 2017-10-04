package TalonOneClient

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

func signPayload(body []byte) string {
	const s = "4aece57f964f4784"
	decoded, _ := hex.DecodeString(s)
	sig := hmac.New(md5.New, decoded)
	sig.Write(body)
	return hex.EncodeToString(sig.Sum(nil))
}

func doRequest(req *http.Request) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("well, oops...")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Status)
	fmt.Println("response Body:", string(body))
}
