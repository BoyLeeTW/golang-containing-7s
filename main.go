package main

import (
	"errors"
	"log"
)

const (
	num = 1000
)

func main() {
	count, err := containing7s(num)
	if err != nil {
		log.Fatalln("error:", err.Error())
	}
	log.Println(count)
}

func containing7s(number int) (int, error) {
	// input should greater than 0
	if number <= 0 {
		return 0, errors.New("invalid input")
	}

	// init parameter if input is valid
	count := 0
	target := 7

	// iterate through all numbers
	for i := 1; i <= number; i++ {
		if checkContainsTarget(i, target) {
			count++
		}
	}

	return count, nil
}

func checkContainsTarget(input int, target int) bool {
	for input > 0 {
		// check if minimal digit is equal to target
		minimalDigitValue := input % 10
		if minimalDigitValue == target {
			return true
		}

		// check next digit
		input /= 10
	}

	return false
}
