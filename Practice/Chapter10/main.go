package main

import (
	"fmt"
)

func abc(Channel chan string) {
	Channel <- "a"
	Channel <- "b"
	Channel <- "c"

}

func def(Channel chan string) {
	Channel <- "d"
	Channel <- "e"
	Channel <- "f"
}

func main() {
	Channel1 := make(chan string)
	Channel2 := make(chan string)
	go abc(Channel1)
	go def(Channel2)
	fmt.Println(<-Channel1)
	fmt.Println(<-Channel2)
	fmt.Println(<-Channel1)
	fmt.Println(<-Channel2)
	fmt.Println(<-Channel1)
	fmt.Println(<-Channel2)
	fmt.Println("End")
}
