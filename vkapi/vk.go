package vkapi

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	apiVersion = "5.37"
)

type (
	// APIClient is struct for holding context
	APIClient struct {
		http *http.Client
	}
)

// NewAPIClient constructor
func NewAPIClient() *APIClient {
	apiCient := new(APIClient)
	apiCient.http = http.DefaultClient
	return apiCient
}

// APIRequest makes request to api and returns response as string
func (apiClient *APIClient) APIRequest(method string, params map[string]string) (string, error) {
	url := apiClient.getRequestURL(method, params)
	resp, err := apiClient.http.Get(url.String())
	if err != nil {
		return "", errors.New("Can't get request from api")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("Can't get body from request")
	}

	return string(data), nil
}

func (apiClient *APIClient) getRequestURL(method string, params map[string]string) (data url.URL) {
	resultURL := url.URL{
		Host:   "api.vk.com",
		Scheme: "https",
		Path:   "method/" + method,
	}
	query := resultURL.Query()
	query.Set("v", apiVersion)
	for key, value := range params {
		query.Set(key, value)
	}
	resultURL.RawQuery = query.Encode()

	return resultURL
}
