package CDSretrieve

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/SLG-European-Projects/cds-go/models"
	"github.com/SLG-European-Projects/cds-go/utils"
)

const (
	jobsPath = "jobs/"
)

// GetJobs is a method to get the list of submitted jobs
func GetJobs(headers models.Headers, qp models.QueryParams, BaseURL string, HTTPClient *http.Client) (jbs models.GetJobs, err error) {
	var endpoint = BaseURL + path + jobsPath

	finalUrl := utils.AddQueries(endpoint, qp)

	req, err := http.NewRequest("GET", finalUrl, nil)
	req.Header.Set("Content-Type", "application/json")
	req = headers.AddToHTTP(req)

	resp, err := HTTPClient.Do(req)
	var jobs models.GetJobs
	if err != nil {
		fmt.Println(err.Error())
		return jobs, err
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return jobs, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&jobs)
	return jobs, err
}

// GetOneJob is a method to get a job by ID.
func GetOneJob(jobId string, headers models.Headers, qp models.QueryParams, BaseURL string, HTTPClient *http.Client) (jb models.PostProcessExecution, err error) {
	var endpoint = BaseURL + path + jobsPath + jobId

	finalUrl := utils.AddQueries(endpoint, qp)

	req, err := http.NewRequest("GET", finalUrl, nil)
	req.Header.Set("Content-Type", "application/json")
	req = headers.AddToHTTP(req)

	resp, err := HTTPClient.Do(req)
	var job models.PostProcessExecution
	if err != nil {
		fmt.Println(err.Error())
		return job, err
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return job, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&job)
	return job, err
}

// GetJobResult is a method to get a jobs result by ID.
func GetJobResult(jobId string, headers models.Headers, BaseURL string, HTTPClient *http.Client) (jobResult models.AssetWrapper, err error) {
	var endpoint = BaseURL + path + jobsPath + jobId + "/results"

	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req = headers.AddToHTTP(req)

	resp, err := HTTPClient.Do(req)
	var result models.AssetWrapper
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return result, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}
