package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const version = "0.0.0"

type AppStatus struct {
	Version     string `json:"version"`
	Environment string `json:"environemnt"`
	Status      string `json:"status"`
}

type config struct {
	port int
	env  string
	db   struct {
		uri string
	}
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	godotenv.Load()

	var cfg config = config{
		db: struct{uri string}{uri: os.Getenv("POSTGRES_URI")},
	}

	flag.IntVar(&cfg.port, "port", 4000, "Server to listen on") // we can specify the port via command line ex: go run cmd/http/main.go -port 8080
	flag.StringVar(&cfg.env, "env", "local", "(local|dev|prod)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDb(cfg)

	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Second,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	app.logger.Println("Starting server on port", app.config.port)

	err = server.ListenAndServe()

	if err != nil {
		log.Println(err)
	}

}

func openDb(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.uri)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		return nil, err
	}

	return db, nil
}
