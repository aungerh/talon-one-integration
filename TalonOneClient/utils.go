package TalonOneClient

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// BuildAndRequest builds and triggers the request
func BuildAndRequest(method string, p *Payload, client *Client, dest string) {
	body, _ := json.Marshal(*p)
	req, _ := http.NewRequest(method, dest, bytes.NewBuffer(body))
	signatureVal := fmt.Sprintf("signer=%d; signature=%s", client.AppID, signPayload(client.AppKey, body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Signature", signatureVal)
	doRequest(req)
}

func signPayload(appKey string, body []byte) string {
	decoded, _ := hex.DecodeString(appKey)
	sig := hmac.New(md5.New, decoded)
	sig.Write(body)
	return hex.EncodeToString(sig.Sum(nil))
}

func doRequest(req *http.Request) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Status)
	fmt.Println("response Body:", string(body))
}
