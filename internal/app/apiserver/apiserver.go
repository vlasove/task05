package apiserver

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq" // postgres driver
	"github.com/vlasove/test05/internal/app/store/sqlstore"
)

// Start ...
func Start(config *Config) error {
	log.Println("starting connection to database...")
	config.DatabaseURL = config.DatabaseConnector.buildConnStr()
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	log.Println("connection successfully builded")

	log.Println("starting configurating storage...")
	store := sqlstore.New(db)
	log.Println("starting configurating API server ...")
	srv := newServer(store)
	log.Println("API server is ready. Working on port", config.BindAddr)

	return http.ListenAndServe(config.BindAddr, srv)
}

// newDB ...
func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
