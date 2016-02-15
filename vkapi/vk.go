package vkapi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// APIRequest makes request to api and returns response as string
func APIRequest(method string, params map[string]string) (string, error) {
	url := getRequestURL(method, params)
	resp, err := http.Get(url.String())
	if err != nil {
		log.Println("Can't get request from api")
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Can't get body from request")
		return "", err
	}

	fmt.Println("data:", string(data))
	return string(data), nil
}

func getRequestURL(method string, params map[string]string) (data url.URL) {
	resultURL := url.URL{
		Host:   "api.vk.com",
		Scheme: "https",
		Path:   "method/" + method,
	}
	query := resultURL.Query()
	query.Set("v", "5.37")
	for key, value := range params {
		query.Set(key, value)
	}
	resultURL.RawQuery = query.Encode()

	return resultURL
}
