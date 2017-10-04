package TalonOneClient

const (
	customerProfileEndpoint  = "https://demo.talon.one/v1/customer_profiles/"
	customerSessionsEndpoint = "https://demo.talon.one/v1/customer_sessions/"
	eventsEndpoint           = "https://demo.talon.one/v1/events"
)

// NewClient credentials initialization
func NewClient(appID int, appKey string) *Payload {
	p := &Payload{
		AppID:  appID,
		AppKey: appKey,
	}
	return p
}
