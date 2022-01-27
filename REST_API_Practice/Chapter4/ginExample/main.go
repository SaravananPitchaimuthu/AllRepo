package main

import "database/sql"

var DB *sql.DB

type StationResource struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

func main() {

}
