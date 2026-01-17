package db

import (
	"database/sql"
	"fmt"
)

const (
	host   = "go_db"
	port   = 5432
	user   = "postgres"
	pass   = "password"
	dbname = "go_database"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host = %s port %d user %s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Conectado em " + dbname)
	return db, nil
}
