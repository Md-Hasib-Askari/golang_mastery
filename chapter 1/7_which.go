package main

import (
	"os"
	"path/filepath"
)

func which() {
	arguments := os.Args
	if len(arguments) == 1 {
		println("Usage: go run 7_which.go <command>")
		os.Exit(1)
	}

	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, p := range pathSplit {
		fullPath := filepath.Join(p, file)
		fileInfo, err := os.Stat(fullPath)
		if err != nil {
			println(fullPath)
			continue
		}

		mode := fileInfo.Mode()
		if !mode.IsRegular() {
			continue
		}

		if mode&0111 != 0 {
			println(fullPath)
			os.Exit(1)
		}

	}
}
