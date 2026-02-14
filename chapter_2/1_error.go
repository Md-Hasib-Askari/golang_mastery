package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

/*
	errors.New() → when the message is fixed
	fmt.Errorf() → when the message needs formatting

	# default function signature:
	- func something(...) error
*/

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d. UserID: %d", a, b, os.Getuid())
	}
	return nil
}

func check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("this is a custom error message")
	}
	return nil
}

func main1() {
	// error without formatting
	err := check(0, 10)
	if err != nil {
		println("Error:", err.Error())
	}
	println("check() executed normally!")

	// error with formatting
	err = formattedError(0, 0)
	if err != nil {
		println("Error:", err.Error())
	}

	// Using strconv.Atoi()
	num, err := strconv.Atoi("-abc") // -123 -> valid, "abc" -> invalid
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Number:", num)
	}
}
