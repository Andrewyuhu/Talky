package main

import (
	"log"
	"net/http"
	"os"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
var infoLogger = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)

func readJSON(r *http.Request, dst any) error {
	return nil
}
