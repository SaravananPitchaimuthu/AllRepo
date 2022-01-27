package datafile

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}

	err = file.Close()
	if err != nil {
		return numbers, nil
	}
	if scanner.Err() != nil {
		return numbers, nil
	}
	fmt.Println(numbers)
	return numbers, nil
}
