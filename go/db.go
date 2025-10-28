package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB() {
	var err error

	dsn := "root:@tcp(127.0.0.1:3306)/produtos"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Banco inacessível:", err)
	}

	log.Println("✅ Banco conectado com sucesso!")
}
