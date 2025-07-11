package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/MoXcz/sql-server-linux-practice/internal/models"
	"github.com/joho/godotenv"
	_ "github.com/microsoft/go-mssqldb"
)

type application struct {
	account *models.AccountModel
}

func main() {
	godotenv.Load(".env")
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		Logger.Error("DB_URL variable must be defined (use a .env file)", "DB_URL", dbURL)
	}

	db, err := openDB(dbURL)
	if err != nil {
		Logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := application{
		account: &models.AccountModel{DB: db},
	}

	listenAddr := ":8080"
	srv := http.Server{
		Addr:    listenAddr,
		Handler: app.routes(),
	}

	Logger.Info("starting server", "addr", listenAddr)

	err = srv.ListenAndServe()
	Logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mssql", dsn) // open pool of connections
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
