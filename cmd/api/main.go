package main

import (
	"fmt"
	"log"

	"github.com/Steve-Rynjah/go-social-feeds/internal/db"
	"github.com/Steve-Rynjah/go-social-feeds/internal/env"
	"github.com/Steve-Rynjah/go-social-feeds/internal/store"
)

const version = "0.0.1"

func main() {
	_config := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialfeeds?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "dev"),
	}

	db, err := db.New(
		_config.db.addr,
		_config.db.maxOpenConns,
		_config.db.maxIdleConns,
		_config.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	fmt.Println("Database pool connection establish")

	store := store.NewStorage(db)

	app := &application{
		config: _config,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
