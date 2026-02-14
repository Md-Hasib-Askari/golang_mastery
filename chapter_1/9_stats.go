package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func stats() {
	arguments := os.Args
	if len(arguments) == 1 {
		println("Usage: go run 9_stats.go <number1> <number2> ...")
		return
	}

	var min, max float64
	var initialized bool = false
	nValues := 0
	var sum float64

	for _, k := range arguments[1:] {
		n, err := strconv.ParseFloat(k, 64)
		if err != nil {
			println("Invalid number:", k)
			continue
		}

		nValues++
		sum += n

		if !initialized {
			min, max = n, n
			initialized = true
			continue
		}

		min = math.Min(min, n)
		max = math.Max(max, n)
	}

	fmt.Println("Number of values:", nValues)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

	if nValues < 0 {
		fmt.Println("No valid numbers provided.")
		return
	}

	// mean
	meanValue := sum / float64(nValues)
	fmt.Println("Mean:", meanValue)

	// standard deviation
	var sumOfSquares float64
	for _, k := range arguments[1:] {
		n, err := strconv.ParseFloat(k, 64)
		if err != nil {
			continue
		}
		sumOfSquares += math.Pow(n-meanValue, 2)
	}
	if nValues > 0 {
		standardDeviation := math.Sqrt(sumOfSquares / float64(nValues))
		fmt.Println("Standard deviation:", standardDeviation)
	}
}
