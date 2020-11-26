package env

import "os"

var MYSQL_USER string
var MYSQL_PASSWORD string
var HMAC_SECRET string

func init() {
	MYSQL_USER = readEnv("MYSQL_USER")
	MYSQL_PASSWORD = readEnv("MYSQL_PASSWORD")
	HMAC_SECRET = readEnv("HMAC_SECRET")
}

func readEnv(str string) string {
	if env := os.Getenv(str); env == "" {
		panic("Cannot read the environment variable: " + str)
	} else {
		return env
	}
}
