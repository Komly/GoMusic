package vkapi

import (
	"fmt"
	"testing"
)

func TestGetURL(t *testing.T) {

	resultURL := getRequestURL("wall.get", nil)

	if resultURL.String() != "https://api.vk.com/method/wall.get?v=5.37" {
		t.Error("Invalid url", resultURL.String())
	}

	params := map[string]string{"owner_id": "-1"}

	result := getRequestURL("wall.get", params)
	if result.String() != "https://api.vk.com/method/wall.get?owner_id=-1&v=5.37" {
		t.Error("invalir url", result)
	}
}

func TestGetRequest(t *testing.T) {
	params := map[string]string{"owner_id": "-1"}
	data, err := APIRequest("wall.get", params)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}
