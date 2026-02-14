package main

import (
	"fmt"
	"os"
	"strconv"
)

func process() {
	// Setup
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}

	// Counter for initialized variables
	var total, nInts, nFloats int
	invalid := make([]string, 0)

	// Process Arguments
	for _, k := range arguments[1:] {
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInts++
			continue
		}
		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}

		// Both Fail → Invalid
		invalid = append(invalid, k)

		// Reporting Results
		fmt.Println("#read:", total, "#ints:", nInts, "#floats:", nFloats)

		if len(invalid) > total {
			fmt.Println("Too much invalid input:", len(invalid))
			for _, s := range invalid {
				fmt.Println(s)
			}
		}

	}
}
