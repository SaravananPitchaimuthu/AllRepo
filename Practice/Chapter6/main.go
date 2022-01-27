package main

import (
	"fmt"
	"log"

	"github.com/SaravananPitchaimuthu/Practice/Chapter6/datetime"
)

func main() {
	var err error
	event := datetime.Event{}
	err = event.SetTitle("Hello Saravanan")
	if err != nil {
		log.Panic(err)
	}
	err = event.AddYear(2021)
	if err != nil {
		log.Panic(err)
	}
	err = event.AddMonth(8)
	if err != nil {
		log.Panic(err)
	}
	err = event.AddDay(31)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(event.Title())
	fmt.Println(event.Year())
}
