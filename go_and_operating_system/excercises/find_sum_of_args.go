package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one of more floats.")
		os.Exit(1)
	}

	arguments := os.Args

	k := 1
	var err = errors.New("")
	for err != nil {
		if k >= len(arguments) {
			fmt.Println("None of the arguments is a float!")
			return
		}
		_, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}

	fmt.Println("Sum of all Integers:", FindSumOfIntegers(arguments))
}

func FindSumOfIntegers(arguments []string) float64 {
	var err = errors.New("")
	var n float64
	var sum float64
	for _, argument := range arguments {
		n, err = strconv.ParseFloat(argument, 64)
		if err == nil {
			sum += n
		}
	}
	return sum
}
