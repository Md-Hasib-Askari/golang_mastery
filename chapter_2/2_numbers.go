package main

import "math"

func main2() {
	// complex numbers
	c1 := complex(2, 3) // 2 + 3i
	c2 := 4 + 5i

	println("c1:", c1)
	println("c2:", c2)

	/*
		Default complex type → complex128
		Explicit cast 		 → changes it to complex64
	*/
	var c3 complex64 = complex64(c1 + c2) // converting to complex64
	println("c3:", c3)

	cZero := c3 - c3 // still complex64
	println("cZero:", cZero)

	// integer types
	x := 12
	k := 5
	div := x / k         // integer division
	println("div:", div) // prints 2, not 2.4

	divFloat := float64(x) / float64(k) // converting to float64 for accurate division
	println("divFloat:", divFloat)      // prints 2.4

	// Zero Values
	var m, n float64 // zero value is 0.0
	m = 1.223
	println("m:", m, "n:", n)

	/*
		| Type    | Zero Value |
		| ------- | ---------- |
		| int     | 0          |
		| float   | 0          |
		| bool    | false      |
		| string  | ""         |
		| pointer | nil        |
	*/

	// math.MaxInt, math.MinInt
	println("MaxInt:", math.MaxInt) // 2147483647 for 32-bit, 9223372036854775807 for 64-bit
	println("MinInt:", math.MinInt) // -2147483648 for 32-bit, -9223372036854775808 for 64-bit

}
