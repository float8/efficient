package efficient

type L map[string]string

type config struct {
	Addr       string
	Middleware []Middleware
	Debug      bool
	AppName    string
	Lang       L
}

var Config  = config{
	AppName: "efficient",
	Debug:   true,
}