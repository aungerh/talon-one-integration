package main

import (
	"math/rand"
	"time"

	"github.com/zyxan/talon-one-integration/TalonOneClient"
)

var r *rand.Rand

func init() {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomString pretty much honors its name
func RandomString(strlen int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := ""
	for i := 0; i < strlen; i++ {
		index := r.Intn(len(chars))
		result += chars[index : index+1]
	}
	return result
}

func main() {
	client := TalonOneClient.NewClient(60, "4aece57f964f4784")
	var str = RandomString(10)

	client.Payload = TalonOneClient.Payload{
		URLParams: "tronald_dump",
		Attributes: TalonOneClient.Attributes{
			// triggers a rule in campaign #1
			Name: "Trump",
		},
	}
	client.UpdateCustomerProfile()

	client.Payload = TalonOneClient.Payload{
		ProfileID: "306",
		SessionID: "tronald_dump",
		Type:      "dangerous",
		Attributes: TalonOneClient.Attributes{
			// triggers a rule in campaign #3
			VeryDangerous: "i_am_the_one_who_knocks",
		},
	}
	client.SendEvents()

	client.Payload = TalonOneClient.Payload{
		// triggers a rule in campaign #2, onSessionCreate
		URLParams:  "session_" + str,
		ProfileID:  "306",
		State:      "open",
		Total:      200,
		Attributes: TalonOneClient.Attributes{},
	}
	client.UpdateCustomerSession()
}
