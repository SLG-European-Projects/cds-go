package models

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

const (
	customLayout = "2006-01-02T15:04:05.000000"
)

// ClientProperties structure
type ClientProperties struct {
	BaseURL    string
	ApiKey     string
	HTTPClient *http.Client
}

// Client structure
type Client struct {
	C ClientProperties
}

// Link structure
type Link struct {
	Rel      *string `json:"rel,omitempty"`
	Type     *string `json:"type,omitempty"`
	Href     string  `json:"href"`
	Title    *string `json:"title,omitempty"`
	Hreflang *string `json:"hreflang,omitempty"`
}

type CADSmessage struct {
	Id       string    `json:"id"`
	Date     time.Time `json:"date,omitempty"`
	Summary  string    `stac_version:"summary,omitempty"`
	URL      string    `json:"url,omitempty"`
	Content  string    `json:"content,omitempty"`
	Severity string    `json:"severity,omitempty"`
	Live     bool      `json:"live,omitempty"`
}

// Collection structure
type Collection struct {
	Type               string      `json:"type"`
	Id                 string      `json:"id"`
	StacVersion        string      `stac_version:"type"`
	Title              string      `json:"title"`
	Description        string      `json:"description"`
	Keywords           []string    `json:"keywords"`
	License            string      `json:"license"`
	Extent             interface{} `json:"extent"`
	Temporal           interface{} `json:"temporal"`
	Links              []Link      `json:"links"`
	Assets             interface{} `json:"assets"`
	Published          time.Time   `json:"published"`
	Updated            time.Time   `json:"updated"`
	CADSdisabledReason string      `json:"cads:disabled_reason"`
	CADSmessage        CADSmessage `json:"cads:message,omitempty"`
	SCIDOI             string      `json:"sci:doi,omitempty"`
}

// CollectionResponse structure
type CollectionList struct {
	Collections    []Collection `json:"collections"`
	Links          []Link       `json:"links"`
	NumberMatched  int          `stac_version:"numberMatched"`
	NumberReturned int          `json:"numberReturned"`
}

type Form struct {
	Css      string  `json:"css,omitempty"`
	Details  Details `json:"details,omitempty"`
	Help     string  `json:"help,omitempty"`
	Label    string  `json:"label,omitempty"`
	Name     string  `json:"name,omitempty"`
	Required bool    `json:"required,omitempty"`
	Type     string  `json:"type,omitempty"`
}

type Details struct {
	Columns           int               `json:"columns,omitempty"`
	ID                int               `json:"id,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	Values            []string          `json:"values,omitempty"`
	Accordion         bool              `json:"accordion,omitempty"`
	AccordionGroups   bool              `json:"accordionGroups,omitempty"`
	Displayaslist     bool              `json:"displayaslist,omitempty"`
	Fullheight        bool              `json:"fullheight,omitempty"`
	Withmap           bool              `json:"withmap,omitempty"`
	Wrapping          bool              `json:"wrapping,omitempty"`
	Precision         int               `json:"precision,omitempty"`
	MaximumSelections int               `json:"maximumSelections,omitempty"`
	TextFile          string            `json:"text:file,omitempty"`
	Information       string            `json:"information,omitempty"`
	AccordionOptions  *AccordionOpts    `json:"accordionOptions,omitempty"`
	Default           any               `json:"default,omitempty"`
	Extentlabels      any               `json:"extentlabels,omitempty"`
	Groups            []interface{}     `json:"groups,omitempty"`
	Range             *RangeLocal       `json:"range,omitempty"`
	ChangeVisible     bool              `json:"changevisible,omitempty"`
	ConCat            string            `json:"concat,omitempty"`
	Latidude          Coords            `json:"latitude,omitempty"`
	Longitude         Coords            `json:"longitude,omitempty"`
	Projection        Projection        `json:"projection,omitempty"`
	Text              string            `json:"text,omitempty"`
	Fields            []Fields          `json:"fields,omitempty"`
}

type AccordionOpts struct {
	OpenGroups interface{} `json:"openGroups,omitempty"`
	Searchable bool        `json:"searchable,omitempty"`
}

type RangeLocal struct {
	E float32 `json:"e,omitempty"`
	N float32 `json:"n,omitempty"`
	W float32 `json:"w,omitempty"`
	S float32 `json:"s,omitempty"`
}

type FormRespLocal struct {
	Name     string `json:"name,omitempty"`
	Required bool   `json:"required,omitempty"`
	Type     string `json:"type,omitempty"`
}

type Coords struct {
	Default   int        `json:"default,omitempty"`
	Precision int        `json:"precision,omitempty"`
	Range     CoordRange `json:"range,omitempty" `
}

type CoordRange struct {
	Min int `json:"min,omitempty"`
	Max int `json:"max,omitempty"`
}

type Projection struct {
	ID      int  `json:"id,omitempty"`
	Overlay bool `json:"overlay,omitempty"`
	Use     bool `json:"use,omitempty"`
}

type Fields struct {
	Comments    string `json:"comments,omitempty"`
	MaxLength   int    `json:"maxlength,omitempty"`
	Placeholder string `json:"placeholder,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Type        string `json:"type,omitempty"`
}

type QueryParams struct {

	// For GetCollections
	Q            *string    `json:"q,omitempty"`
	Keywords     *[]string  `json:"kw,omitempty"`
	IDx          *[]string  `json:"idx,omitempty"`
	SortBy       *string    `json:"sortby,omitempty"`
	Page         *int       `json:"page,omitempty"`
	Limit        *int       `json:"limit,omitempty"`
	SearchStats  *bool      `json:"search_stats,omitempty"`
	ProcessID    *string    `json:"processID,omitempty"`
	Status       *string    `json:"status,omitempty"`
	QoS          *bool      `json:"qos,omitempty"`
	Request      *bool      `json:"request,omitempty"`
	Log          *bool      `json:"log,omitempty"`
	LogStartTime *time.Time `json:"logStartTime,omitempty"`
}

type Process struct {
	Id                 string   `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Version            string   `json:"version"`
	JobControlOptions  []string `json:"jobControlOptions"`
	OutputTransmission []string `json:"outputTransmission"`
	Links              []Link   `json:"links"`
}

type ProcessesList struct {
	Processes []Process `json:"processes"`
	Links     []Link    `json:"links"`
}

// Headers struct represents the HTTP headers
type Headers struct {
	UserAgent      string `json:"User-Agent,omitempty"`
	AcceptEncoding string `json:"Accept-Encoding,omitempty"`
	Accept         string `json:"Accept,omitempty"`
	Connection     string `json:"Connection,omitempty"`
	PrivateToken   string `json:"PRIVATE-TOKEN,omitempty"`
	ContentLength  int    `json:"Content-Length,omitempty"`
	ContentType    string `json:"Content-Type,omitempty"`
	Authorization  string `json:"Authorization,omitempty"`
	XCardsPortal   string `json:"X-CARDS-PORTAL,omitempty"`
}

// Convert struct to `http.Header` and exclude empty fields
func (h *Headers) AddToHTTP(req *http.Request) *http.Request {

	// Use reflection to iterate over struct fields
	val := reflect.ValueOf(*h)
	typ := reflect.TypeOf(*h)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tag := typ.Field(i).Tag.Get("json")

		if tag != "" {
			tagName := tag
			if commaIdx := len(tagName) - len(",omitempty"); commaIdx > 0 {
				tagName = tag[:commaIdx]
			}

			switch field.Kind() {
			case reflect.String:
				if field.String() != "" {
					req.Header.Add(tagName, field.String())
				}
			case reflect.Int:
				if field.Int() > 0 {
					req.Header.Add(tagName, strconv.Itoa(int(field.Int())))
				}
			}
		}
	}

	return req
}

// Execute represents the structure of the request for the "Execute" operation
type Execute struct {
	Inputs     interface{} `json:"inputs,omitempty"`     // Can be "object" or "null"
	Outputs    *Outputs    `json:"outputs,omitempty"`    // Can be "object" or "null"
	Response   *string     `json:"response,omitempty"`   // "raw" or "document" (defaults to "raw")
	Subscriber *Subscriber `json:"subscriber,omitempty"` // Can be "object" or "null"
}

// Outputs represents the "outputs" section
type Outputs struct {
	Format           *Format `json:"format,omitempty"`           // Format object or null
	TransmissionMode *string `json:"transmissionMode,omitempty"` // "value" or "reference" (or null)
}

// Format represents the "format" object inside "outputs"
type Format struct {
	MediaType string      `json:"mediaType,omitempty"` // Media type (e.g., "application/json"), string or null
	Encoding  *string     `json:"encoding,omitempty"`  // Encoding (e.g., "gzip"), string or null
	Schema    interface{} `json:"schema,omitempty"`    // Schema can be string, object, or null
}

// Subscriber represents the "subscriber" section containing the callback URIs
type Subscriber struct {
	SuccessURI    *string `json:"successUri,omitempty"`    // URI string (>= 1 character) or null
	InProgressURI *string `json:"inProgressUri,omitempty"` // URI string (>= 1 character) or null
	FailedURI     *string `json:"failedUri,omitempty"`     // URI string (>= 1 character) or null
}

// Custom time struct
type CustomTime struct {
	time.Time
}

// UnmarshalJSON to handle time format
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	if str == "null" {
		ct.Time = time.Time{} // Zero value for nil time
		return nil
	}

	parsedTime, err := time.Parse(`"`+customLayout+`"`, str)
	if err != nil {
		return fmt.Errorf("time parsing failed: %w", err)
	}
	ct.Time = parsedTime
	return nil
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, ct.Time.Format(customLayout))), nil
}

// PostProcessExecution represents the response or data structure for a process execution
type PostProcessExecution struct {
	ProcessID *string             `json:"processID,omitempty"`
	Type      string              `json:"type"`
	JobID     string              `json:"jobID"`
	Status    string              `json:"status"`
	Message   *string             `json:"message,omitempty"`
	Created   *CustomTime         `json:"created,omitempty"`
	Started   *CustomTime         `json:"started,omitempty"`
	Finished  *CustomTime         `json:"finished,omitempty"`
	Updated   *CustomTime         `json:"updated,omitempty"`
	Progress  *int                `json:"progress,omitempty"`
	Links     *[]Link             `json:"links,omitempty"`
	Metadata  *StatusInfoMetadata `json:"metadata,omitempty"`
}

// StatusInfoMetadata represents the "metadata" field
type StatusInfoMetadata struct {
	Request         map[string]interface{} `json:"request,omitempty"`         // Can be a map representing the request details
	Results         map[string]interface{} `json:"results,omitempty"`         // Can be a map representing the results
	DatasetMetadata map[string]interface{} `json:"datasetMetadata,omitempty"` // Can be a map representing dataset metadata
	QoS             map[string]interface{} `json:"qos,omitempty"`             // Quality of Service (QoS) metadata
	Log             [][][]interface{}      `json:"log,omitempty"`             // Array of array of arrays [string, string, any]
	Origin          *string                `json:"origin,omitempty"`          // Origin string or null
}

type GetJobs struct {
	Jobs     []PostProcessExecution `json:"jobs"`
	Links    *[]Link                `json:"links,omitempty"`
	Metadata *StatusInfoMetadata    `json:"metadata,omitempty"`
}

// FileDetails represents the nested fields prefixed by "file:"
type FileDetails struct {
	Checksum  string `json:"file:checksum"`
	Size      int64  `json:"file:size"`
	LocalPath string `json:"file:local_path"`
}

// Value represents the "value" object
type Value struct {
	Type string      `json:"type"`
	Href string      `json:"href"`
	File FileDetails `json:"-"` // Combined "file" fields into this struct
}

// Asset represents the "asset" object
type Asset struct {
	Value Value `json:"value"`
}

// AssetWrapper represents the full JSON structure
type AssetWrapper struct {
	Asset Asset `json:"asset"`
}
