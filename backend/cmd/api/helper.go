package main

import (
	"encoding/json"
	"log"
	"maps"
	"net/http"
	"os"
)

type envelope map[string]any

var errorLogger = log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
var infoLogger = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)

func (app *application) readJSON(r *http.Request, dst any) error {
	js := json.NewDecoder(r.Body)
	js.DisallowUnknownFields()

	err := js.Decode(dst)

	if err != nil {
		return err
	}

	return nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {

	js, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		return err
	}

	js = append(js, '\n')

	maps.Copy(w.Header(), headers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
