package main

import (
	"fmt"
	"time"
)

func myPrint(start, end int) {
	for i := start; i < end; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	time.Sleep(100 * time.Microsecond)
}

func goRoutines() {
	for i := 0; i < 10; i++ {
		go myPrint(i*10, (i+1)*10)
	}

	// Wait for all go routines to finish
	time.Sleep(1 * time.Second)
}
