package wrag

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var tokenResponse *TokenResponse

func fetchToken() {
	tokenResponse = getToken()

	time.AfterFunc(intToDuration(tokenResponse.ExpiresIn-20), func() { fetchToken() })
}

// TokenResponse is a struct containing the response of the OAuth token provided by Reddit
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func getToken() *TokenResponse {
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", config.Auth.Username)
	data.Set("password", config.Auth.Password)

	req, err := http.NewRequest("POST", endpoints.Apis["access_token"], strings.NewReader(data.Encode()))
	req.SetBasicAuth(config.Auth.ClientID, config.Auth.ClientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", config.Auth.UserAgent)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("Unable to retrieve OAUTH2 token.")
	}

	var tokenResponse TokenResponse

	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		log.Fatal("Unable to parse token response from Reddit.")
		panic(err)
	}
	return &tokenResponse
}
