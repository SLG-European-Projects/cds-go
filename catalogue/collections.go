package CDScatalogue

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/SLG-European-Projects/cds-go/models"
)

const (
	collecionPath = "collections/"
)

// GetCollections is a method to get the list of available collections
func GetCollections(BaseURL string, HTTPClient *http.Client) (colls models.CollectionList, err error) {
	var endpoint = BaseURL + path + collecionPath
	req, err := http.NewRequest("GET", endpoint, nil)
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

// GetOneCollection is a method to get a feature by ID.
func GetOneCollection(collectionId string, BaseURL string, HTTPClient *http.Client) (colls models.Collection, err error) {
	var endpoint = BaseURL + path + collecionPath + collectionId
	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := HTTPClient.Do(req)
	var collection models.Collection
	if err != nil {
		fmt.Println(err.Error())
		return collection, err
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return collection, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&collection)
	return collection, err
}

// GetCollectionForm is a methof to get a collections form
func GetCollectionForm(collectionId string, BaseURL string, HTTPClient *http.Client) (form []models.Form, err error) {
	var endpoint = BaseURL + path + collecionPath + collectionId + "/form.json"
	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := HTTPClient.Do(req)
	var forms []models.Form
	if err != nil {
		fmt.Println(err.Error())
		return forms, err
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return forms, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&forms)
	return forms, err
}

// messages
