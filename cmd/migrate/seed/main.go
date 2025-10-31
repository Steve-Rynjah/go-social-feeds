package main

import (
	"log"

	"github.com/Steve-Rynjah/go-social-feeds/internal/db"
	"github.com/Steve-Rynjah/go-social-feeds/internal/env"
	"github.com/Steve-Rynjah/go-social-feeds/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialfeeds?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store)
}
