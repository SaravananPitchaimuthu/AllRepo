package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Open(filename string) (*os.File, error) {
	fmt.Println("Opening the file")
	return os.Open(filename)
}

func Close(file *os.File) {
	fmt.Println("Closing the file")
	file.Close()
}

func GetFloats(filename string) ([]float64, error) {
	var numbers []float64
	file, err := Open(filename)
	if err != nil {
		return nil, err
	}
	defer Close(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

func main() {
	numbers, err := GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("Sum :%0.2f", sum)

}
