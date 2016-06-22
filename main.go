package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	contentTypeHdr = "Content-Type"
	strJSONMIME    = "application/json"
	errParseJSON   = `{"error": "Could not decode request: JSON parsing failed"}`
)

// Simple setup for passing different log.Logger configurations.
var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger

	appPort = "8080"
)

// TVShow is the model used to capture information about... TV shows.
type TVShow struct {
	Title        string `json:"title"`
	Slug         string `json:"slug"`
	DRM          bool   `json:"drm"`
	EpisodeCount int
	Image        struct {
		ShowImage string `jsom:"showImage"`
	} `json:"image"`
}

// Validate checks if the show is DRMed and has episodes.
func (s *TVShow) Validate() bool {
	// TODO: Make this generic by acceptiong an input predicate instead.
	return s.DRM && s.EpisodeCount > 0
}

// ReqPayload is the place holder where we marshal or requests to.
type ReqPayload struct {
	Shows []TVShow `json:"payload"`
	Skip  int      `json:"skip"`
	Take  int      `json:"take"`
	Total int      `json:"totalRecords"`
}

// Response is the staging struct used to encapsulate the expected resoponse schema.
type Response struct {
	Image string `json:"image"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

// filterShows run Validate() against all shows and returns only those that pass.
func filterShows(p *ReqPayload) []Response {
	var r []Response
	for _, s := range p.Shows {
		if s.Validate() {
			t := Response{
				Title: s.Title,
				Slug:  s.Slug,
				Image: s.Image.ShowImage,
			}
			r = append(r, t)
		}
	}
	return r
}

// indexHandler is the main http.Handler that handles all the app logic.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Use a simple routing scheme here
	//
	// TODO: Upgrade to a proper or more robust route handling mechanism
	if r.Method != "PUT" {
		http.Error(w, "Unsupported operation", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "Unsupported endpoint", http.StatusForbidden)
		return
	}

	payload := ReqPayload{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		w.Header().Set(contentTypeHdr, strJSONMIME)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, errParseJSON)
		return
	}
	resp := map[string]interface{}{
		"response": filterShows(&payload),
	}
	data, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func init() {
	flags := log.Ldate | log.Ltime | log.Lshortfile

	Trace = log.New(ioutil.Discard, "TRACE: ", flags)
	Info = log.New(os.Stdout, "INFO: ", flags)
	Warning = log.New(os.Stdout, "WARNING: ", flags)
	Error = log.New(os.Stderr, "ERROR: ", flags)

	if os.Getenv("PORT") != "" {
		appPort = os.Getenv("PORT")
	}
}

func main() {
	http.Handle("/", Adapt(http.HandlerFunc(indexHandler), AddHeader(XHeader, XValue), Recover(Error)))
	log.Printf("Starting server on port %v...\n", appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, nil))
}
