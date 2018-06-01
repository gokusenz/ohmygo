package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pallat/ohmygo/http"
	"github.com/pallat/ohmygo/postgres"
)

func main() {
	// Connect to database.
	db, err := postgres.Open(os.Getenv("DB"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create services.
	us := &postgres.UserService{DB: db}

	// Attach to HTTP handler.
	h := &http.Handler{}
	h.UserService = us

	mux := http.NewServeMux()
	mux.Handle("/api/", h)

	fmt.Println("listen :8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
