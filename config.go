package efficient

import (
	"github.com/whf-sky/efficient/widget/lang"
)

type config struct {
	Addr       string
	Middleware []Middleware
	Debug      bool
	AppName    string
	Lang       lang.Lang
}

var Config  = config{
	AppName: "efficient",
	Debug:   true,
	Lang:    lang.ZH,
}
