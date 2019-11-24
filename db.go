package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		fmt.Println("WARNING: No database url specified, not connecting to postgres!")
		return
	}
	var err error
	DB, err = sql.Open("postgres", url)
	if err != nil {
		// ok if there IS a url then we're expected to be able to connect
		// so if THAT fails, then that's a real error
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		// apparently this DOUBLE CHECKS that it's up?
		panic(err)
	}
	schema()
}

func schema() {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS tempmutes(

		discord_id TEXT   NOT NULL UNIQUE PRIMARY KEY,
		expiration BIGINT NOT NULL

	);`)
	if err != nil {
		panic(err)
	}
}
