package main

import (
	
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/AliSinaDevelo/Test-Signer/internal/database/postgres"
	"github.com/AliSinaDevelo/Test-Signer/internal/handlers"
	"github.com/AliSinaDevelo/Test-Signer/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	appConfig := config.LoadConfig()


	// initialize database connection
	db, err := postgres.NewDB(appConfig.DBHost, appConfig.DBPort, appConfig.DBUser, appConfig.DBPassword, appConfig.DBName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// initialize HTTP server
	router := mux.NewRouter()

	// Handlers for signing and verifying
	signHandler := handlers.NewSignHandler(db)
	verifyHandler := handlers.NewVerifyHandler(db)

	router.HandleFunc("/sign", signHandler.Sign).Methods("POST")
	router.HandleFunc("/verify", verifyHandler.Verify).Methods("GET")

	port := ":8080"
	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))

	
}	