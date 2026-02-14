package main

import (
	"fmt"
	"runtime"
	"time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Printf("Sys = %v MiB", mem.Sys)
	fmt.Printf("NumGC = %v\n", mem.NumGC)
}

func main() {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 5e8) // Allocate 50 MB
		if s == nil {
			fmt.Println("Allocation failed")
		}
		time.Sleep(5 * time.Second)
		printStats(mem)

	}
}
