package main

import (
	"fmt"
	"log"

	"github.com/SaravananPitchaimuthu/Practice/Chapter3/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("chicken_data.txt")
	if err != nil {
		log.Panic(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("The average chicken required:%0.2f kg", sum/float64(len(numbers)))
}
