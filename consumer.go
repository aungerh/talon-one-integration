package main

import "github.com/zyxan/talon-one-integration/TalonOneClient"

func main() {
	client := TalonOneClient.NewClient(60, "4aece57f964f4784")

	client.Payload = TalonOneClient.Payload{
		URLParams: "tronald_dump",
		Attributes: TalonOneClient.Attributes{
			Name: "Alex",
		},
	}
	client.Payload.UpdateCustomerProfile(client)

	client.Payload = TalonOneClient.Payload{
		ProfileID:  "306",
		SessionID:  "tronald_dump",
		Type:       "dangerous",
		Attributes: TalonOneClient.Attributes{},
	}
	client.Payload.SendEvents(client)

	client.Payload = TalonOneClient.Payload{
		URLParams:  "NewCustomerSession60",
		ProfileID:  "307",
		State:      "closed",
		Total:      200,
		Attributes: TalonOneClient.Attributes{},
	}
	client.Payload.UpdateCustomerSession(client)
}
