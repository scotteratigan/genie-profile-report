package main

import (
	"log"
	"os"
)

type dualLogger struct {
	info  *log.Logger
	error *log.Logger
}

func makeLogger() dualLogger {
	return dualLogger{
		info:  log.New(os.Stdout, "", 0),
		error: log.New(os.Stderr, "", 0),
	}
}
