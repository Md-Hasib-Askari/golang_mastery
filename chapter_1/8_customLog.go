package main

import (
	"io"
	"log"
	"os"
	"path"
)

func customLog() {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")

	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	LstdFlags := log.Ldate | log.Ltime | log.Lshortfile

	iLog := log.New(f, "iLog ", LstdFlags)

	iLog.Println("This is an info log")

	// multiple log destinations
	w := io.MultiWriter(f, os.Stderr)
	logger := log.New(w, "multiLog ", LstdFlags)
	logger.Println("This is a multi log")
}
