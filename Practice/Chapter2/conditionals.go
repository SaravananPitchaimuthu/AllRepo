package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I have chosen a random number between 1 to 100")
	fmt.Println("Can you guess it?")

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have ", 10-guesses, " guesses left")
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Make a guess")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Your guess is Low")
		} else if guess > target {
			fmt.Println("your guess is High")
		} else {
			success = true
			fmt.Println("Good Job ! you guessed it")
			break
		}
	}
	if !success {
		fmt.Println("Sry you have lost all your chances correct answer is...", target)
	}

}
