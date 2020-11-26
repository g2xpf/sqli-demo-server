package main

import (
	"net/http"

	_ "github.com/g2xpf/sqli-demo-server/db"
	"github.com/labstack/echo"
	em "github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(em.CORSWithConfig(em.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:8082"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	u := e.Group("/cert", Certificate)

	e.POST("/users/register", HandlePostUsersRegister)
	e.POST("/users/login", HandlePostUsersLogin)

	u.DELETE("/users/logout", HandleDeleteUsersLogout)
	u.GET("/users/info", HandleGetUsersInfo)
	u.POST("/posts", HandlePostPosts)

	e.Logger.Fatal(e.Start(":3001"))
}
