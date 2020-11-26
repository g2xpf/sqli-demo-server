package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"

	env "github.com/g2xpf/sqli-demo-server/env"
	"github.com/labstack/echo"
)

const HEADER_ID = "X-Session-Id"
const HEADER_CERT = "X-Session-Cert"

func Certificate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Request().Header.Get(HEADER_ID)
		cert := c.Request().Header.Get(HEADER_CERT)

		if id == "" || cert == "" || !hmacEquals(id, cert, env.HMAC_SECRET) {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		c.Set("user_id", id)

		return next(c)
	}
}

func hmacEquals(id, cert, key string) bool {
	raw := ([]byte)(id)
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write(raw)

	rawCert, err := base64.StdEncoding.DecodeString(cert)
	if err != nil {
		return false
	}

	return hmac.Equal(rawCert, mac.Sum(nil))
}
