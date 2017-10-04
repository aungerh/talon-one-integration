package TalonOneClient

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
