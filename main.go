package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

// Structs to match the JSON response schema
type Estate struct {
	MunicipalityID       int      `json:"municipality_id"`
	Owner                string   `json:"owner"`
	Address              string   `json:"address"`
	Block                string   `json:"block"`
	Lot                  string   `json:"lot"`
	EstateType           string   `json:"estate_type"`
	Hash                 string   `json:"hash"`
	RecentActivity       string   `json:"recent_activity"`
	Photos               []string `json:"photos"`
	Slug                 string   `json:"slug"`
	MakePaymentURL       string   `json:"make_payment_url"`
	SmartContractAddress string   `json:"smart_contract_address"`
	Qualifier            string   `json:"qualifier"`
	IsVerified           bool     `json:"is_verified"`
	IsFeature            bool     `json:"is_feature"`
	Vacant               bool     `json:"vacant"`
	SaleDate             string   `json:"sale_date"`
	Lat                  string   `json:"lat"`
	Lng                  string   `json:"lng"`
}

type ApiResponse struct {
	Count   int      `json:"count"`
	HasMore bool     `json:"has_more"`
	Data    []Estate `json:"data"`
}

type ResponseContainer struct {
	Status   string      `json:"status"`
	Response ApiResponse `json:"response"`
}

type Listing struct {
	UID                  string   `json:"uid"`
	IsVerified           bool     `json:"is_verified"`
	Photos               []string `json:"photos"`
	EstateType           string   `json:"estate_type"`
	Hash                 string   `json:"hash"`
	Owner                string   `json:"owner"`
	Address              string   `json:"address"`
	Block                string   `json:"block"`
	Lot                  string   `json:"lot"`
	LastTaxPaymentAmount string   `json:"last_tax_payment_amount"`
	LastTaxPaymentDate   string   `json:"last_tax_payment_date"`
	Lat                  string   `json:"lat"`
	Lng                  string   `json:"lng"`
	HouseClass           string   `json:"house_class"`
	YearBuilt            string   `json:"year_built"`
	// ... include other fields as necessary
}

type ListingResponse struct {
	Status   string  `json:"status"`
	Response Listing `json:"response"`
}

func main() {
	// Serve static files like CSS, JS, and images
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/asset-registry/", AssetRegistryHandler)
	http.HandleFunc("/listing/", ListingDetailsHandler)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
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
	tmpl, err := template.ParseFiles("listings.html")
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
func FetchData(apiEndpoint string) ([]Estate, error) {
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

func fetchEstateDetails(slug string) (*Listing, error) {
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
	tmpl, err := template.ParseFiles("listing.html")
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
