package main

import (
	"log"

	"github.com/SaravananPitchaimuthu/REST_API_Practice/Chapter7/base62Example/base62"
)

func main() {
	x := 100
	base62String := base62.ToBase62(x)
	log.Println(base62String)
	normalNumber := base62.ToBase10(base62String)
	log.Println(normalNumber)
}
