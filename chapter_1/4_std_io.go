package main

import (
	"fmt"
	"os"
	"strconv"
)

func std_io() {
	// 1️⃣ Reading from Standard Input
	fmt.Printf("Please give me your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is", name)

	// 2️⃣ Command-Line Arguments (os.Args)

	/*
		| Index | Value              |
		| ----- | ------------------ |
		| 0     | Path of executable |
		| 1     | First argument     |
		| 2     | Second argument    |
		| ...   | ...                |
	*/

	// Step 1 — Get Arguments
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}

	// Step 2 — Prepare Variables
	var min, max float64
	var initialized = 0

	// Step 3 — Process Arguments
	for i := 1; i < len(arguments); i++ {
		var number float64
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			fmt.Printf("Invalid number: %s\n", arguments[i])
			continue
		}

		if initialized == 0 {
			min, max = number, n
			initialized = 1
		}

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Printf("Min: %f, Max: %f\n", min, max)
}
