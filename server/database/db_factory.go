package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "quotedb.db")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados:", err)
	}
	sqlStmt := `CREATE TABLE IF NOT EXISTS quote (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    code TEXT NOT NULL,
    codein TEXT NOT NULL,
    name TEXT NOT NULL,
    high TEXT NOT NULL,
    low TEXT NOT NULL,
    varBid TEXT NOT NULL,
    pctChange TEXT NOT NULL,
    bid TEXT NOT NULL,
    ask TEXT NOT NULL,
    timestamp TEXT NOT NULL,
    create_date TEXT NOT NULL
	);`
	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, sqlStmt)
	}
}
