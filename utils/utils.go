package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/SLG-European-Projects/cds-go/models"
)

// GrainFolders helps grab the correct folder from a cursor.
func AddQueries(base string, iQuery models.QueryParams) (URLwithQuery string) {

	parsedURL, _ := url.Parse(base)
	query := url.Values{}

	// Add fields to query if they are non-empty
	if iQuery.Q != nil {
		query.Add("q", *iQuery.Q)
	}
	if iQuery.Keywords != nil {
		query.Add("kw", strings.Join(*iQuery.Keywords, ","))

	}
	if iQuery.IDx != nil {
		query.Add("idx", strings.Join(*iQuery.IDx, ","))
	}
	if iQuery.SortBy != nil {
		query.Add("sortby", *iQuery.SortBy)
	}
	if iQuery.Page != nil {
		query.Add("page", strconv.Itoa(*iQuery.Page))
	}
	if iQuery.Limit != nil {
		query.Add("limit", strconv.Itoa(*iQuery.Limit))
	}
	if iQuery.SearchStats != nil {
		if *iQuery.SearchStats {
			query.Add("search_stats", "true")
		} else {
			query.Add("search_stats", "false")
		}
	}

	// Append the query string to the base URL
	parsedURL.RawQuery = query.Encode()

	return parsedURL.String()
}

// downloadFileAsBytes performs an HTTP GET request and returns the file contents as a byte slice
func DownloadFileAsBytes(url string) ([]byte, error) {
	// Create an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP GET request: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download file: status code %d", resp.StatusCode)
	}

	// Read the response body into a byte slice
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return data, nil
}
