package main

import "github.com/zyxan/talon-one-integration/TalonOneClient"

func main() {
	client := TalonOneClient.NewClient(60, "4aece57f964f4784")

	client.Payload = TalonOneClient.Payload{
		URLParams: "tronald_dump",
		Attributes: TalonOneClient.Attributes{
			Name: "Trump",
		},
	}
	client.UpdateCustomerProfile()

	client.Payload = TalonOneClient.Payload{
		ProfileID: "306",
		SessionID: "tronald_dump",
		Type:      "dangerous",
		Attributes: TalonOneClient.Attributes{
			VeryDangerous: "i_am_the_one_who_knocks",
		},
	}
	client.SendEvents()

	client.Payload = TalonOneClient.Payload{
		URLParams:  "GoCLISession1",
		ProfileID:  "306",
		State:      "open",
		Total:      200,
		Attributes: TalonOneClient.Attributes{},
	}
	client.UpdateCustomerSession()
}
