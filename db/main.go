package db

import (
	"fmt"

	env "github.com/g2xpf/sqli-demo-server/env"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var schemata = [...]string{`
CREATE TABLE IF NOT EXISTS users (
	id int NOT NULL AUTO_INCREMENT,
	name TEXT NOT NULL DEFAULT '',
	password TEXT NOT NULL DEFAULT '',
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(id)
) CHARSET = utf8mb4;
`, `
CREATE TABLE IF NOT EXISTS posts (
	id int NOT NULL AUTO_INCREMENT,
	user_id VARCHAR(63) NOT NULL DEFAULT '',
	message TEXT NOT NULL DEFAULT '',
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(id)
) CHARSET = utf8mb4;
`}

var DB *sqlx.DB

func init() {
	var err error
	DB, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(127.0.0.1:3306)/sqli_demo?parseTime=true&charset=utf8mb4&loc=Asia%%2FTokyo", env.MYSQL_USER, env.MYSQL_PASSWORD))
	if err != nil {
		panic(err)
	}

	for _, schema := range schemata {
		DB.MustExec(schema)
	}
}
