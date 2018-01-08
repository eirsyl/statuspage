package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eirsyl/statuspage/pkg"
	"net/http"
	"time"
)

// NewClient returns a http.Client used to access the API
func NewClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

type API struct {
	apiUrl string
	token  string
	client *http.Client
}

// NewAPI initializes a new API client that can be used to access the API
func NewAPI(apiUrl, token string) (*API, error) {
	client := NewClient()

	if apiUrl == "" || token == "" {
		return nil, errors.New("Both apiUrl and token is required")
	}

	return &API{
		apiUrl: apiUrl,
		token:  token,
		client: client,
	}, nil
}

func (a *API) CreateRequest(url, method string, payload interface{}) *http.Request {
	var data []byte

	if payload != nil {
		data, _ = json.Marshal(payload)
	}

	request, _ := http.NewRequest(
		method, fmt.Sprintf("%s%s", pkg.RemoveLastSlash(a.apiUrl), url), bytes.NewReader(data),
	)
	request.Header.Add("AUTHORIZATION", a.token)
	request.Header.Add("CONTENT-TYPE", "application/json")
	return request
}
