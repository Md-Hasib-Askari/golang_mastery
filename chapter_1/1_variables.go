package main

import (
	"fmt"
	"math"
)

// Global Variables
var Global int = 1234
var AnotherGlobal int = -5678

// Local Variable Without Initialization
var j int

func variables() {
	// Local Variable With Initialization
	var k int

	i := Global + AnotherGlobal
	println("i:", i)

	j = Global
	println("j:", j)

	k = AnotherGlobal
	println("k:", k)

	// Type Casting
	m := math.Abs(float64(AnotherGlobal))

	fmt.Printf("Global=%d  i=%d  j=%d  k=%d  m: %f\n", Global, i, j, k, m)
}
