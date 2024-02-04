package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/LesterKim/balclony/internal/models"
)

type ApiResponse struct {
	Count   int             `json:"count"`
	HasMore bool            `json:"has_more"`
	Data    []models.Estate `json:"data"`
}

type ResponseContainer struct {
	Status   string      `json:"status"`
	Response ApiResponse `json:"response"`
}

func AssetRegistryHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the slug from the URL path.
	slug := strings.TrimPrefix(r.URL.Path, "/asset-registry/")

	// The API endpoint URL
	apiEndpoint := fmt.Sprintf("https://admin.balcony.technology/api/search?municipality_slug=%s&sort=featured&page=1", slug)

	// Call the FetchData function to get the listings
	estates, err := FetchData(apiEndpoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse the HTML file with your template
	tmpl, err := template.ParseFiles("templates/listings.html")
	if err != nil {
		// If there's an error, return a 500 internal server error and log the error
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	// Execute the template with the estates data
	err = tmpl.Execute(w, map[string]interface{}{
		"Estates": estates,
	})
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
		return
	}
}

// FetchData calls the provided API endpoint and extracts the "data" field.
func FetchData(apiEndpoint string) ([]models.Estate, error) {
	// Make the GET request to the API endpoint
	resp, err := http.Get(apiEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the body of the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response into the ResponseContainer struct
	var responseContainer ResponseContainer
	err = json.Unmarshal(body, &responseContainer)
	if err != nil {
		return nil, err
	}

	// Check if the status is "success"
	if responseContainer.Status != "success" {
		return nil, fmt.Errorf("API returned non-success status: %s", responseContainer.Status)
	}

	// Return the "data" slice from the response
	return responseContainer.Response.Data, nil
}
