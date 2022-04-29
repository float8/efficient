package efficient

import (
	"github.com/whf-sky/efficient/widget/lang"
)

type eConfig struct {
	Addr       string
	Middleware []Middleware
	Debug      bool
	APPName    string
	Lang       lang.Lang
}

var EConfig *eConfig = &eConfig{
	APPName: "efficient",
	Debug:   true,
	Lang:    lang.ZH,
}
