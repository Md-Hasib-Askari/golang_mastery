package main

import (
	"fmt"
	"strconv"
)

func someFunction() error {
	// Simulate some work and return an error
	return nil // or return an actual error
}

func if_else() {
	// 1️⃣ The if Statement
	condition := true
	if condition {
		// code
	}

	// The Most Important Go Pattern
	err := someFunction()
	if err != nil {
		// handle error
	}

	// 2️⃣ Comments
	// This is a single-line comment

	/*
		This is a multi-line comment
	*/

	// 3️⃣ The switch Statement

	// 🔹 Form 1 — With Expression
	argument := "1"
	switch argument {
	case "0":
		fmt.Println("Zero!")
	case "1":
		fmt.Println("One!")
	default:
		fmt.Println("Value:", argument)
	}

	// 🔹 Form 2 — Without Expression
	value := 10
	switch {
	case value == 0:
		fmt.Println("Zero!")
	case value > 0:
		fmt.Println("Positive")
	case value < 0:
		fmt.Println("Negative")
	}

	// 4️⃣ Converting Input
	argument = "123"
	number, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Println("Error converting input:", err)
	}

	fmt.Println("Converted number:", number)
}
