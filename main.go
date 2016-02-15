package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Komly/GoMusic/vkapi"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"net/url"
	"os"
)

type (
	Config struct {
		ClientId string `json:"client_id"`
	}
)

func downloadMusic(access_token string) {
	vkApi := vkapi.NewAPIClient()
	params := map[string]string{
		"access_token": access_token,
		"count":        "500",
	}
	response, err := vkApi.APIRequest("audio.get", params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error, while getting audio list: %v", err.Error())
		os.Exit(1)
	}
	fmt.Printf("%v", response)
}

func parseOauthResponse(oauthResponse string) (string, error) {
	u, err := url.Parse(oauthResponse)
	if err != nil {
		return "", errors.New("Invalid url, try again")
	}
	fragment := u.Fragment
	query, err := url.ParseQuery(fragment)
	if err != nil {
		return "", errors.New("Invalid url, try again")
	}
	if _, ok := query["access_token"]; ok && len(query["access_token"]) > 0 {
		return query["access_token"][0], nil
	}
	return "", errors.New("Invalid access token, try again")
}

func main() {
	input := bufio.NewScanner(os.Stdin)

	config := Config{}
	config_source, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read config.json: %v\n", err.Error())
		os.Exit(1)
	}
	if err := json.Unmarshal([]byte(config_source), &config); err != nil {
		fmt.Fprintf(os.Stderr, "Can't decode config.json: %v\n", err.Error())
		os.Exit(1)
	}

	if config.ClientId == "" {
		fmt.Fprintf(os.Stderr, "Invalid client_id in config.json\n")
		os.Exit(1)
	}

	open.Run(fmt.Sprintf("https://oauth.vk.com/authorize?client_id=%s&scope=audio"+
		"&redirect_uri=https://oauth.vk.com/blank.html&display=page&response_type=token", config.ClientId))
	fmt.Printf("Enter url from browser window:\n")
	input.Scan()
	oauthResponse := input.Text()
	if access_token, err := parseOauthResponse(oauthResponse); err == nil {
		downloadMusic(access_token)
	} else {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err.Error())
		os.Exit(1)
	}

}
