package main

import (
	hub "backend/cmd/api/websocket"
	"backend/internals/data"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type application struct {
	cfg         config
	errorLogger *log.Logger
	infoLogger  *log.Logger
	usermodel   *data.UserModel
	hubManager  *hub.HubManager
}

type config struct {
	port int
	db   struct {
		dsn string
	}
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://myuser:mypassword@localhost:5433/mydatabase?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	app := &application{
		cfg:         cfg,
		errorLogger: errorLogger,
		infoLogger:  infoLogger,
		hubManager:  hub.NewHubManager(),
	}

	db, err := openDB(app.cfg.db.dsn)

	if err != nil {
		app.errorLogger.Fatal(err)
	}

	app.usermodel = data.New(db)
	app.infoLogger.Println("Connected to database")

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", app.cfg.port),
		Handler: app.router(),
	}

	go app.hubManager.Run()

	fmt.Printf("Starting server on: %d\n", app.cfg.port)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
