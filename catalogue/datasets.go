package CDScatalogue

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/SLG-European-Projects/cds-go/models"
	"github.com/SLG-European-Projects/cds-go/utils"
)

const (
	datasetPath = "datasets/"
)

// SearchDatasets is a method to get the list of available collections
func SearchDatasets(qp models.QueryParams, BaseURL string, HTTPClient *http.Client) (colls models.CollectionList, err error) {
	var endpoint = BaseURL + path + datasetPath

	finalUrl := utils.AddQueries(endpoint, qp)

	req, err := http.NewRequest("GET", finalUrl, nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := HTTPClient.Do(req)
	var collections models.CollectionList

	if err != nil {
		fmt.Println(err.Error())
		return collections, err
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return collections, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&collections)
	return collections, err
}
