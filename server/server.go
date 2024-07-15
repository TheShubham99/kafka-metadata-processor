package server

import (
	"atlan-lily/producer"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMetadata(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	message := vars["message"]
	fmt.Fprintf(w, "Metadata for ID: %s\n", id)

	if err := producer.Produce("metadata_topic", map[string]interface{}{id: message}); err != nil {
		log.Fatalf("Failed to produce message: %v", err)
	}
}

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/metadata/{id}/{message}", GetMetadata).Methods("GET")

	// Handle the HTTP server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Starting server on :8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
