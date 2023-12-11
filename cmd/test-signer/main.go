package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/AliSinaDevelo/Test-Signer/internal/database/postgres"
	"github.com/AliSinaDevelo/Test-Signer/internal/handlers"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// initialize database connection
	db, err := postgres.NewDB(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// initialize HTTP server
	router := mux.NewRouter()

	// Handlers for signing and verifying
	signHandler := handler.NewSignHandler(db)
	verifyHandler := handler.NewVerifyHandler(db)

	router.HandleFunc("/sign", signHandler.Sign).Methods("POST")
	router.HandleFunc("/verify", verifyHandler.Verify).Methods("POST")

	port := ":8080"
	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))

	
}	