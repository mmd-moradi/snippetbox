package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mmd-moradi/snippetbox/internal/db"
	"github.com/mmd-moradi/snippetbox/internal/models"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "Http network address")
	seed := flag.Bool("seed", false, "Seed the database with initial data")
	dsn := flag.String("dsn", "dev:devpass@tcp(localhost:3306)/snippetbox?parseTime=true", "MySQL data source name(connection string)")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	dbConn, err := openDB(*dsn)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer dbConn.Close()

	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: dbConn},
	}

	// db.CreateTable(dbConn)
	if *seed {
		logger.Info("Seeding the database...")
		db.Seed(dbConn)
		return
	}

	logger.Info("Starting server", "add", *addr)

	err = http.ListenAndServe(*addr, app.routes())

	logger.Error(err.Error())
	os.Exit(1)

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
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
