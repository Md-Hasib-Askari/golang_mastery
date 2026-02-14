package main

import "fmt"

func main() {
	/*
		Non-numeric types:
		- string -> collection of bytes (UTF-8 encoded)
			- string literals are immutable
			- can be created using double quotes " " or backticks ` `
			- backticks allow multi-line strings and ignore escape sequences
		- bool
		- pointer
		- character (byte or rune, no char in Go)
		- date and time (using time package)

		ASCII and Unicode:
		- ASCII: byte
		- Unicode: rune (alias for int32)
		- len(string) ≠ number of characters, Unicode characters can take multiple bytes
	*/

	rune := '€'
	println("Rune:", rune)                 // prints the Unicode code point value of '€'
	println("Rune as char:", string(rune)) // converts the rune back to a string representation
	fmt.Printf("%c", rune)                 // prints '€' using formatted output

	// NOTE: event though rune == int32, Go treats them as different types.

	// Converting Between String and []byte
	str := "Hello, Go!"
	byteSlice := []byte(str)  // string -> byte slice
	str2 := string(byteSlice) // byte slice -> string
	println("Original String:", str)
	println("Byte Slice:", byteSlice)
	println("Converted String:", str2)

	/*
		- File I/O uses byte slices
		- Networking uses byte slices
		- Low-level work uses byte slices
	*/
}
