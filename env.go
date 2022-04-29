package efficient

import "os"

var env = func() string {
	env := os.Getenv("EFFICIENT_ENV")
	if env == "" {
		env = "production"
	}
	return env
}()

func GetEnv() string {
	return env
}

func SetEnv(e string) {
	env = e
}
