package main

import (
	"log"

	"github.com/SaravananPitchaimuthu/REST_API_Practice/Chapter7/basicExample/helper"
)

func main() {
	_, err := helper.InitDB()
	if err != nil {
		log.Println(err)
	}
	log.Println("Database tables are successfully Initialized")
}
