package main

import (
	"fmt"

	"github.com/SaravananPitchaimuthu/Practice/Chapter11/prose"
)

func main() {
	phrases := []string{"Myself", "My TCS Friends"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	phrases = []string{"Myself", "My TCS Friends", "My Gym Friends"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	phrases = []string{"My Parents"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
}
