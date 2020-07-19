package main

import (
	"database/sql"
	"fmt"
	"log"
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
	err = schema()
	if err != nil {
		// Failed to create table or something
		panic(err)
	}
}

func schema() (err error) {
	_, err = DB.Exec(`
		CREATE EXTENSION IF NOT EXISTS "pgcrypto";
	`)
	if err != nil {
		log.Println("Unable to load pgcrypto extension")
		panic(err)
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS mutes (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			discord_id TEXT NOT NULL,
			channel_id TEXT,
			expiration TIMESTAMP
		);
	`)
	if err != nil {
		log.Println("Unable to create mutes table")
		panic(err)
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS cringe (
			image TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Println("Unable to create cringe table")
		panic(err)
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS nicks (
			id   TEXT PRIMARY KEY,
			nick SERIAL
		)
	`)
	if err != nil {
		log.Println("Unable to create nicks table")
		panic(err)
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS nicktrade (
			id   TEXT,
			desirednick INTEGER,

			UNIQUE(id, desirednick)
		)
	`)
	if err != nil {
		log.Println("Unable to create nicktrade table")
		panic(err)
	}



	return
}
