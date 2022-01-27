package main

import (
	"fmt"
	"log"

	"github.com/SaravananPitchaimuthu/Practice/Chapter5/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Panic(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	for name, votes := range counts {
		fmt.Printf("No of votes for %s:%d \n", name, votes)
	}
}
