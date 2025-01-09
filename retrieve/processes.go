package CDSretrieve

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/SLG-European-Projects/cds-go/models"
	"github.com/SLG-European-Projects/cds-go/utils"
)

const (
	processPath = "processes/"
)

// GetProcesses is a method to get the list of available processes
func GetProcesses(qp models.QueryParams, BaseURL string, HTTPClient *http.Client) (colls models.ProcessesList, err error) {
	var endpoint = BaseURL + path + processPath

	finalUrl := utils.AddQueries(endpoint, qp)

	req, err := http.NewRequest("GET", finalUrl, nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := HTTPClient.Do(req)
	var processes models.ProcessesList

	if err != nil {
		fmt.Println(err.Error())
		return processes, err
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return processes, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&processes)
	return processes, err
}

// GetOneProcess is a method to get a feature by ID.
func GetOneProcess(processId string, BaseURL string, HTTPClient *http.Client) (proc models.Process, err error) {
	var endpoint = BaseURL + path + processPath + processId
	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := HTTPClient.Do(req)
	var process models.Process
	if err != nil {
		fmt.Println(err.Error())
		return process, err
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return process, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&process)
	return process, err
}

// PostProcess is a method to get a process by ID.
func PostProcess(input models.Execute, headers models.Headers, processId string, BaseURL string, HTTPClient *http.Client) (pproc models.PostProcessExecution, err error) {
	var endpoint = BaseURL + path + processPath + processId + "/execution"

	// Encode the struct to JSON
	body, err := json.Marshal(input)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))

	// Add headers
	req.Header.Set("Content-Type", "application/json")
	req = headers.AddToHTTP(req)
	resp, err := HTTPClient.Do(req)

	var pprocess models.PostProcessExecution
	if err != nil {
		fmt.Println(err.Error())
		return pprocess, err
	}

	if resp.StatusCode != 201 {
		err = errors.New(resp.Status)
		return pprocess, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&pprocess)
	return pprocess, err
}
