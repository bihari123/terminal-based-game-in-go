package config

type AppConfiguration struct {
	Host          string
	Port          int
	LogLevel      string
	AccessLogFile string
	LogFolder     string
	LogFile       string
}
