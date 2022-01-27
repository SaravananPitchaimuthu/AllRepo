package helper

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "0.0.0.0"
	port     = 5432
	user     = "root"
	password = "secret"
	dbname   = "root"
)

func InitDB() (*sql.DB, error) {

	connectionString := fmt.Sprintf("host =%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS web_url(ID SERIAL PRIMARY KEY, URL TEXT NOT NULL);")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec()
	if err != nil {
		return nil, err
	}
	return db, nil

}
