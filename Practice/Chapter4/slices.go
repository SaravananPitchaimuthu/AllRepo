package main

import (
	"fmt"
)

func findAverage(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	fmt.Println(len(numbers))
	return sum / len(numbers)
}
func main() {
	areas := make([]int, 5)
	litres := []float64{12.0, 213.0, 45.8}
	areas[0] = 1
	fmt.Println(areas[1:], litres[:2])
	avg := findAverage(1, 2, 3, 54, 34)
	fmt.Println(avg)
}
