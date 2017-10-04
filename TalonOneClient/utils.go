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

// BuildAndRequest builds the request and fires it
func BuildAndRequest(p *Payload, dest string) {
	js, _ := json.Marshal(*p)
	fmt.Println(string(js))
	req, _ := http.NewRequest("PUT", dest, bytes.NewBuffer(js))
	signatureVal := fmt.Sprintf("signer=%d; signature=%s", 60, signPayload(js))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Signature", signatureVal)
	doRequest(req)
}

func signPayload(body []byte) string {
	// environment variable (not in demo)
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
