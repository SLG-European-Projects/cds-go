package test

import (
	"fmt"
	"testing"

	goCDS "github.com/SLG-European-Projects/cds-go"
)

// TestGetProcesses tests that default values are applied
func TestGetProcesses(t *testing.T) {

	client := goCDS.InitClient(apiUrl, apiKey)

	params := map[string]interface{}{
		"limit": 2,
	}

	_, err := client.GetProcesses(params)

	if err != nil {
		fmt.Println(err)
		t.Fatalf("Failed to get the collection form")
	}
}

// TestGetProcesses tests that default values are applied
func TestGetOneProcesses(t *testing.T) {

	client := goCDS.InitClient(apiUrl, apiKey)

	_, err := client.GetOneProcess(process)

	if err != nil {
		fmt.Println(err)
		t.Fatalf("Failed to get the collection form")
	}
}

// TestGetProcesses tests that default values are applied
func TestCreateProcesses(t *testing.T) {

	client := goCDS.InitClient(apiUrl, apiKey)

	body := map[string]interface{}{
		"product_type":   []string{"reanalysis"},
		"variable":       []string{"geopotential"},
		"year":           []string{"2024"},
		"month":          []string{"03"},
		"day":            []string{"01"},
		"time":           []string{"13:00"},
		"pressure_level": []string{"1000"},
		"data_format":    "grib",
	}

	_, err := client.CreateProcess(process, body)

	if err != nil {
		fmt.Println(err)
		t.Fatalf("Failed to get the collection form")
	}
}

func TestGetAllJobs(t *testing.T) {
	client := goCDS.InitClient(apiUrl, apiKey)

	params := map[string]interface{}{
		"limit": 2,
	}

	_, err := client.GetAllJobs(params)

	if err != nil {
		fmt.Println(err)
		t.Fatalf("Failed to get the jobs")
	}
}

func TestGetOneJob(t *testing.T) {
	client := goCDS.InitClient(apiUrl, apiKey)

	params := map[string]interface{}{
		"log":     true,
		"request": true,
	}

	_, err := client.GetOneJob("e2da6c73-cf31-4f52-b346-cedd2f7a4861", params)

	if err != nil {
		fmt.Println(err)
		t.Fatalf("Failed to get the jobs")
	}
}

func TestGetJobResult(t *testing.T) {
	client := goCDS.InitClient(apiUrl, apiKey)

	res, err := client.GetJobResult("e2da6c73-cf31-4f52-b346-cedd2f7a4861")
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
		t.Fatalf("Failed to get the jobs")
	}
}

func TestRetreiveDataset(t *testing.T) {
	client := goCDS.InitClient(apiUrl, apiKey)

	body := map[string]interface{}{
		"product_type":   []string{"reanalysis"},
		"variable":       []string{"geopotential"},
		"year":           []string{"2023"},
		"month":          []string{"06"},
		"day":            []string{"10"},
		"time":           []string{"13:00"},
		"pressure_level": []string{"1000"},
		"data_format":    "grib",
	}

	data, err := client.RetrieveDataset(process, body)

	fmt.Println(len(data))
	if err != nil {
		fmt.Println(err)
		t.Fatalf("Failed to get the collection form")
	}
}
