package db

import (
	"database/sql"
	"fmt"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgress"
	pass   = 1234
	dbname = "postgress"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host = %s port %d user %s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)

	db, err := sql.Open("postgress", psqlInfo)

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
