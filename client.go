package goCDS

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	CDScatalogue "github.com/SLG-European-Projects/cds-go/catalogue"
	"github.com/SLG-European-Projects/cds-go/models"
	CDSretrieve "github.com/SLG-European-Projects/cds-go/retrieve"
	"github.com/SLG-European-Projects/cds-go/utils"
)

const (
	// ClientVersion is used in User-Agent request header to provide server with API level.
	ClientVersion = "0.0.1"

	// httpClientTimeout is used to limit http.Client waiting time.
	httpClientTimeout = 15 * time.Second
)

// Client object
type Client models.Client

// InitClient creates a CDS Client
func InitClient(baseURL string, apiKey string) *Client {
	if baseURL == "" {
		baseURL = "https://cds.climate.copernicus.eu/api/"
	} else {
		if baseURL[len(baseURL)-1:] != "/" {
			baseURL = baseURL + "/"
		}
	}

	return &Client{
		C: models.ClientProperties{
			BaseURL: baseURL,
			ApiKey:  apiKey,
			HTTPClient: &http.Client{
				Timeout: httpClientTimeout,
			},
		},
	}
}

// ICDSClient is the CDS Client Interface
type ICDSClient interface {
	// Get a list with all collections
	GetCollections() (colls models.CollectionList, err error)
	GetOneCollection(id string) (coll models.Collection, err error)
	GetCollectionForm(id string) (forms []models.Collection, err error)
	// GetCollectionConstrains(id string) (coll models.Collection, err error)

	SearchDatasets(queries interface{}) (colls models.CollectionList, err error)

	GetProcesses(queries interface{}) (procs models.ProcessesList, err error)
	GetOneProcess(id string) (procs models.Process, err error)
	CreateProcess(body interface{})

	GetAllJobs(queries interface{}) (jobs models.GetJobs, err error)
	GetOneJob(jobId string, queries interface{}) (job models.PostProcessExecution, err error)
}

// GetCollections is a method to get a list with all collections
func (client *Client) GetCollections() (colls models.CollectionList, err error) {
	return CDScatalogue.GetCollections(client.C.BaseURL, client.C.HTTPClient)
}

// GetOneCollection is a method to a collection by ID
func (client *Client) GetOneCollection(id string) (coll models.Collection, err error) {
	return CDScatalogue.GetOneCollection(id, client.C.BaseURL, client.C.HTTPClient)
}

// GetCollectionForm is a method to a collection's Form
func (client *Client) GetCollectionForm(id string) (coll []models.Form, err error) {
	return CDScatalogue.GetCollectionForm(id, client.C.BaseURL, client.C.HTTPClient)
}

// SearchDatasets is a method to a collection with query params
func (client *Client) SearchDatasets(queries interface{}) (colls models.CollectionList, err error) {

	// Marshal the interface{} to JSON bytes
	jsonBytes, err := json.Marshal(queries)
	if err != nil {
		return colls, err
	}

	// Unmarshal the JSON bytes into the QueryParams struct
	var qp models.QueryParams
	err = json.Unmarshal(jsonBytes, &qp)
	if err != nil {
		return colls, err
	}

	return CDScatalogue.SearchDatasets(qp, client.C.BaseURL, client.C.HTTPClient)
}

// GetProcesses is a method to get a list with all collections
func (client *Client) GetProcesses(queries interface{}) (procs models.ProcessesList, err error) {

	// Marshal the interface{} to JSON bytes
	jsonBytes, err := json.Marshal(queries)
	if err != nil {
		return procs, err
	}

	// Unmarshal the JSON bytes into the QueryParams struct
	var qp models.QueryParams
	err = json.Unmarshal(jsonBytes, &qp)
	if err != nil {
		return procs, err
	}

	return CDSretrieve.GetProcesses(qp, client.C.BaseURL, client.C.HTTPClient)
}

// GetOneProcess is a method to a process by ID
func (client *Client) GetOneProcess(id string) (proc models.Process, err error) {
	return CDSretrieve.GetOneProcess(id, client.C.BaseURL, client.C.HTTPClient)
}

func (client *Client) CreateProcess(processId string, body interface{}) (proc models.PostProcessExecution, err error) {
	headers := models.Headers{
		UserAgent:      "datapi/0.1.1",
		AcceptEncoding: "gzip, deflate",
		PrivateToken:   client.C.ApiKey,
		ContentType:    "application/json",
	}

	execute := models.Execute{
		Inputs: body,
	}
	return CDSretrieve.PostProcess(execute, headers, processId, client.C.BaseURL, client.C.HTTPClient)
}

func (client *Client) GetAllJobs(queries interface{}) (jobs models.GetJobs, err error) {
	headers := models.Headers{
		UserAgent:      "datapi/0.1.1",
		AcceptEncoding: "gzip, deflate",
		PrivateToken:   client.C.ApiKey,
		ContentType:    "application/json",
	}

	// Marshal the interface{} to JSON bytes
	jsonBytes, err := json.Marshal(queries)
	if err != nil {
		return jobs, err
	}

	// Unmarshal the JSON bytes into the QueryParams struct
	var qp models.QueryParams
	err = json.Unmarshal(jsonBytes, &qp)
	if err != nil {
		return jobs, err
	}

	return CDSretrieve.GetJobs(headers, qp, client.C.BaseURL, client.C.HTTPClient)
}

func (client *Client) GetOneJob(jobId string, queries interface{}) (job models.PostProcessExecution, err error) {
	headers := models.Headers{
		UserAgent:      "datapi/0.1.1",
		AcceptEncoding: "gzip, deflate",
		PrivateToken:   client.C.ApiKey,
		ContentType:    "application/json",
	}

	// Marshal the interface{} to JSON bytes
	jsonBytes, err := json.Marshal(queries)
	if err != nil {
		return job, err
	}

	// Unmarshal the JSON bytes into the QueryParams struct
	var qp models.QueryParams
	err = json.Unmarshal(jsonBytes, &qp)
	if err != nil {
		return job, err
	}

	return CDSretrieve.GetOneJob(jobId, headers, qp, client.C.BaseURL, client.C.HTTPClient)
}

func (client *Client) GetJobResult(jobId string) (job models.AssetWrapper, err error) {
	headers := models.Headers{
		UserAgent:      "datapi/0.1.1",
		AcceptEncoding: "gzip, deflate",
		PrivateToken:   client.C.ApiKey,
		ContentType:    "application/json",
	}

	return CDSretrieve.GetJobResult(jobId, headers, client.C.BaseURL, client.C.HTTPClient)
}

func (client *Client) RetrieveDataset(datasetId string, body interface{}) (fileBytes []byte, err error) {

	maxRetries := 10
	retryDelay := 2 * time.Second // Delay between retries
	status := ""

	params := map[string]interface{}{
		"log":     true,
		"request": true,
	}

	process, err := client.CreateProcess(datasetId, body)
	if err != nil {
		return nil, err
	}

	for i := 1; i <= maxRetries; i++ {
		job, err := client.GetOneJob(process.JobID, params) // Fetch the job status
		if err != nil {
			fmt.Printf("Error fetching job: %v\n", err)
			break
		}

		if job.Status != status {
			status = job.Status
			fmt.Printf("Attempt %d: Download Status is '%s'\n", i, job.Status)
		}

		// Check if the job is successful
		if job.Status == "successful" {
			fmt.Println("Job completed successfully!")
			break
		}

		// Wait before retrying
		if i < maxRetries {
			// fmt.Printf("Retrying in %v...\n", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	result, err := client.GetJobResult(process.JobID)
	if err != nil {
		return nil, err
	}

	return utils.DownloadFileAsBytes(result.Asset.Value.Href)

	// return CDSretrieve.GetJobResult(jobId, headers, client.C.BaseURL, client.C.HTTPClient)
}
