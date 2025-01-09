package test

import (
	"fmt"
	"testing"

	goCDS "github.com/SLG-European-Projects/cds-go"
)

// TestNewClientProperties_DefaultValues tests that default values are applied
func TestGetAllCollections(t *testing.T) {

	client := goCDS.InitClient(apiUrl, apiKey)

	_, err := client.GetCollections()

	if err != nil {
		t.Fatalf("Failed to get the collection list")
	}
}

// TestGetOneCollection tests that default values are applied
func TestGetOneCollection(t *testing.T) {

	client := goCDS.InitClient(apiUrl, apiKey)

	_, err := client.GetOneCollection("reanalysis-era5-pressure-levels")

	if err != nil {
		t.Fatalf("Failed to get the collection")
	}
}

// TestGetOneCollection tests that default values are applied
func TestGetForm(t *testing.T) {

	client := goCDS.InitClient(apiUrl, apiKey)

	_, err := client.GetCollectionForm("reanalysis-era5-pressure-levels")

	if err != nil {
		fmt.Println(err)
		t.Fatalf("Failed to get the collection form")
	}
}

// TestSearchDataset tests that default values are applied
func TestSearchDataset(t *testing.T) {

	client := goCDS.InitClient(apiUrl, apiKey)

	params := map[string]interface{}{
		"q":            "reanalysis-era5-pressure-levels",
		"search_stats": true,
	}

	_, err := client.SearchDatasets(params)

	if err != nil {
		fmt.Println(err)
		t.Fatalf("Failed to get the collection form")
	}
}
