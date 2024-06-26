package db

import (
	"database/sql"
	"log"
)

var queryCreateTableSoundBox = `CREATE TABLE IF NOT EXISTS soundbox (
		id TEXT PRIMARY KEY,
		name TEXT,
		code TEXT UNIQUE,
		capacity INTEGER
	);`

var queryCreateTableSound = `CREATE TABLE IF NOT EXISTS sound (
	id TEXT PRIMARY KEY,
	name TEXT,
	soundbox_id TEXT NOT NULL,
	FOREIGN KEY (soundbox_id) REFERENCES soundbox (id)
);`

var queryCreateTableUsers = `CREATE TABLE IF NOT EXISTS user (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	authid TEXT UNIQUE NOT NULL,
	role TEXT
	);`

var queryCreateTableUsersSoundBox = `CREATE TABLE IF NOT EXISTS user_soundbox (
	user_authid TEXT UNIQUE,
	soundbox_id TEXT NOT NULL,
	FOREIGN KEY (user_authid) REFERENCES user (authid)
	FOREIGN KEY (soundbox_id) REFERENCES soundbox (id)
	);`

var queryCreateTableUserToken = `CREATE TABLE IF NOT EXISTS user_token (
	user_authid TEXT UNIQUE,
	token TEXT NOT NULL,
	token_exp DATETIME NOT NULL,
	FOREIGN KEY (user_authid) REFERENCES user (authid)
	);`

func CreateTables(db *sql.DB) {
	queries := []string{queryCreateTableSoundBox, queryCreateTableSound, queryCreateTableUsers, queryCreateTableUsersSoundBox, queryCreateTableUserToken}

	for i := range queries {
		_, err := db.Exec(queries[i])
		if err != nil {
			log.Printf("Query %v has an issue", queries[i])
			log.Fatal(err)
		}
	}
}
