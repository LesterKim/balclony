package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LesterKim/balclony/internal/handlers"
)

// Structs to match the JSON response schema
func main() {
	// Serve static files like CSS, JS, and images
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/asset-registry/", handlers.AssetRegistryHandler)
	http.HandleFunc("/listing/", handlers.ListingDetailsHandler)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
