package config

type ServiceConfiguration struct {
	Metadata      ServiceDetails
	Dataingestion ServiceDetails
	Gamelogic     ServiceDetails
	Tokenxchange  ServiceDetails
	TerminalApp   ServiceDetails
}

type ServiceDetails struct {
	Endpoint string
	Protocol string
	Headers  map[string]string
	Async    bool
}
