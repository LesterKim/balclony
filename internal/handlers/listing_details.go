package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/LesterKim/balclony/internal/models"
)

type ListingResponse struct {
	Status   string         `json:"status"`
	Response models.Listing `json:"response"`
}

func ListingDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the slug from the URL path.
	slug := strings.TrimPrefix(r.URL.Path, "/listing/")

	// Fetch estate details
	estateDetails, err := fetchEstateDetails(slug)
	if err != nil {
		http.Error(w, "Failed to fetch estate details", http.StatusInternalServerError)
		log.Println("Error fetching estate details:", err)
		return
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/listing.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	// Render the template with the estate details
	if err := tmpl.Execute(w, estateDetails); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Println("Error rendering template:", err)
		return
	}
}

func fetchEstateDetails(slug string) (*models.Listing, error) {
	// Define the API URL
	apiURL := "https://admin.balcony.technology/api/estate/" + slug

	// Make the GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read and parse the response body
	var estateResponse ListingResponse
	if err := json.NewDecoder(resp.Body).Decode(&estateResponse); err != nil {
		return nil, err
	}

	// Check if the status is success
	if estateResponse.Status != "success" {
		return nil, fmt.Errorf("API returned non-success status: %s", estateResponse.Status)
	}

	return &estateResponse.Response, nil
}
