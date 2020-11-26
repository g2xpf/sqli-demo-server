package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	. "github.com/g2xpf/sqli-demo-server/db"
	env "github.com/g2xpf/sqli-demo-server/env"

	"github.com/labstack/echo"
)

func sha256sum(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

func hmacEncode(s string) string {
	raw := ([]byte)(s)
	mac := hmac.New(sha256.New, []byte(env.HMAC_SECRET))
	mac.Write(raw)
	certBytes := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(certBytes)
}

type Session struct {
	UserID      string `json:"id" db:"id"`
	Certificate string `json:"cert"`
}

type UserInfo struct {
	Name string `json:"name" db:"name"`
}

func HandlePostUsersLogin(c echo.Context) error {
	formParams, err := c.FormParams()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	username := formParams.Get("username")
	password := formParams.Get("password")

	if username == "" || password == "" {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	var id string
	passwordHashed := sha256sum(password)

	// sqli vulnerability
	query := fmt.Sprintf("SELECT id FROM users WHERE name = '%s' AND password = '%s'", username, passwordHashed)
	err = DB.Get(&id, query)

	// correct
	// err := DB.QueryRow("SELECT id FROM users WHERE name = ? AND password = ?", username, password).Scan(&id)

	switch {
	case err == sql.ErrNoRows:
		return echo.NewHTTPError(http.StatusUnauthorized)
	case err == nil:
		session := Session{id, hmacEncode(id)}
		return c.JSON(http.StatusOK, session)
	default:
		log.Print(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
}

func HandlePostUsersRegister(c echo.Context) error {
	formParams, err := c.FormParams()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	username := formParams.Get("username")
	password := formParams.Get("password")

	var userCount int

	if err := DB.QueryRow("SELECT count(*) FROM users WHERE name = ?", username).Scan(&userCount); err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	} else if userCount >= 1 {
		log.Print(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	passwordHashed := sha256sum(password)

	if _, err := DB.Exec("INSERT INTO users (name, password) VALUES (?, ?)", username, passwordHashed); err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return echo.NewHTTPError(http.StatusOK)
}

func HandleDeleteUsersLogout(c echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

func HandleGetUsersInfo(c echo.Context) error {
	id := c.Get("user_id")
	info := UserInfo{}

	if err := DB.Get(&info, "SELECT name FROM users WHERE id = ?", id); err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, info)
}

func HandlePostPosts(c echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}
